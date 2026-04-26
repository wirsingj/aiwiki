$CanonicalRepoRoot = "C:\Users\wirsi\OneDrive\Desktop\git\aiwiki"

$combinedPath = @(
    ($env:Path -split ";")
    ([Environment]::GetEnvironmentVariable("Path", "Machine") -split ";")
    ([Environment]::GetEnvironmentVariable("Path", "User") -split ";")
) | Where-Object { $_ } | Select-Object -Unique
$env:Path = ($combinedPath -join ";")

function Use-LocalDockerHostFallback {
    if (-not (Get-Command docker -ErrorAction SilentlyContinue)) {
        return
    }

    $oldErrorActionPreference = $ErrorActionPreference
    $oldNativePreference = $PSNativeCommandUseErrorActionPreference
    $ErrorActionPreference = "Continue"
    $PSNativeCommandUseErrorActionPreference = $false

    docker info *> $null
    if ($LASTEXITCODE -eq 0) {
        $ErrorActionPreference = $oldErrorActionPreference
        $PSNativeCommandUseErrorActionPreference = $oldNativePreference
        return
    }

    $existingDockerHost = [Environment]::GetEnvironmentVariable("DOCKER_HOST", "Process")
    if (Get-Command wsl -ErrorAction SilentlyContinue) {
        if (-not (Get-CimInstance Win32_Process -Filter "Name='wsl.exe'" -ErrorAction SilentlyContinue | Where-Object { $_.CommandLine -like "*aiwiki-docker-keeper*" })) {
            Start-Process wsl.exe -WindowStyle Hidden -ArgumentList @(
                "-d", "Ubuntu", "-u", "root", "--", "sh", "-lc",
                "systemctl start docker >/dev/null 2>&1 || service docker start >/dev/null 2>&1 || true; exec -a aiwiki-docker-keeper tail -f /dev/null"
            ) | Out-Null
            Start-Sleep -Seconds 4
        } else {
            wsl -d Ubuntu -u root -- sh -lc "systemctl start docker >/dev/null 2>&1 || service docker start >/dev/null 2>&1 || true" *> $null
        }
    }

    if ($existingDockerHost) {
        docker info *> $null
        $ErrorActionPreference = $oldErrorActionPreference
        $PSNativeCommandUseErrorActionPreference = $oldNativePreference
        return
    }

    $tcpProbe = Test-NetConnection -ComputerName "127.0.0.1" -Port 2375 -InformationLevel Quiet -WarningAction SilentlyContinue
    if ($tcpProbe) {
        $env:DOCKER_HOST = "tcp://127.0.0.1:2375"
    }

    $ErrorActionPreference = $oldErrorActionPreference
    $PSNativeCommandUseErrorActionPreference = $oldNativePreference
}

Use-LocalDockerHostFallback

function Get-AiwikiRoot {
    $current = (Get-Location).Path

    while ($current) {
        $hasBackend = Test-Path (Join-Path $current "backend\go.mod")
        $hasFrontend = Test-Path (Join-Path $current "frontend\package.json")
        $hasCompose = Test-Path (Join-Path $current "docker-compose.yml")

        if ($hasBackend -and $hasFrontend -and $hasCompose) {
            return $current
        }

        $parent = Split-Path -Parent $current
        if ($parent -eq $current) {
            break
        }
        $current = $parent
    }

    if (Test-Path $CanonicalRepoRoot) {
        return $CanonicalRepoRoot
    }

    throw "Could not find AIWIKI repo root. Expected $CanonicalRepoRoot or a parent directory containing backend, frontend, and docker-compose.yml."
}

function Enter-AiwikiRoot {
    $root = Get-AiwikiRoot
    Set-Location $root
    return $root
}

function Require-Command {
    param(
        [Parameter(Mandatory = $true)][string]$Name,
        [string]$InstallHint = ""
    )

    if (-not (Get-Command $Name -ErrorAction SilentlyContinue)) {
        if ($InstallHint) {
            throw "Missing required command '$Name'. $InstallHint"
        }
        throw "Missing required command '$Name'."
    }
}

function Get-DockerComposeCommand {
    if (Get-Command docker -ErrorAction SilentlyContinue) {
        $help = (& docker compose version 2>$null)
        if ($LASTEXITCODE -eq 0 -and $help) {
            return @{
                Command = "docker"
                Prefix = @("compose")
                Label = "docker compose"
            }
        }
    }

    if (Get-Command docker-compose -ErrorAction SilentlyContinue) {
        return @{
            Command = "docker-compose"
            Prefix = @()
            Label = "docker-compose"
        }
    }

    throw "Missing Docker Compose. Install Docker Desktop or Docker CLI plus Docker Compose, then restart your terminal."
}

function Test-PortListening {
    param([Parameter(Mandatory = $true)][int]$Port)

    return [bool](Get-NetTCPConnection -LocalPort $Port -State Listen -ErrorAction SilentlyContinue)
}

function Show-PortOwner {
    param([Parameter(Mandatory = $true)][int]$Port)

    $conn = Get-NetTCPConnection -LocalPort $Port -State Listen -ErrorAction SilentlyContinue | Select-Object -First 1
    if (-not $conn) {
        Write-Host "Port $Port is free."
        return
    }

    $proc = Get-CimInstance Win32_Process -Filter "ProcessId=$($conn.OwningProcess)" -ErrorAction SilentlyContinue
    if ($proc) {
        Write-Host "Port $Port is used by PID $($proc.ProcessId): $($proc.CommandLine)"
    } else {
        Write-Host "Port $Port is used by PID $($conn.OwningProcess)."
    }
}

function Load-DotEnv {
    param([Parameter(Mandatory = $true)][string]$Path)

    if (-not (Test-Path $Path)) {
        return
    }

    Get-Content $Path | ForEach-Object {
        $line = $_.Trim()
        if ($line -eq "" -or $line.StartsWith("#") -or -not $line.Contains("=")) {
            return
        }

        $parts = $line.Split("=", 2)
        $name = $parts[0].Trim()
        $value = $parts[1].Trim()

        if (($value.StartsWith('"') -and $value.EndsWith('"')) -or ($value.StartsWith("'") -and $value.EndsWith("'"))) {
            $value = $value.Substring(1, $value.Length - 2)
        }

        if ($name -and -not [Environment]::GetEnvironmentVariable($name, "Process")) {
            [Environment]::SetEnvironmentVariable($name, $value, "Process")
        }
    }
}

function Stop-ProcessTree {
    param([int]$ProcessId)

    if (-not $ProcessId) {
        return
    }

    $children = Get-CimInstance Win32_Process -Filter "ParentProcessId=$ProcessId" -ErrorAction SilentlyContinue
    foreach ($child in $children) {
        Stop-ProcessTree -ProcessId $child.ProcessId
    }

    Stop-Process -Id $ProcessId -Force -ErrorAction SilentlyContinue
}

function Invoke-Checked {
    param(
        [Parameter(Mandatory = $true)][string]$Command,
        [string[]]$Arguments = @()
    )

    & $Command @Arguments
    if ($LASTEXITCODE -ne 0) {
        throw "Command failed: $Command $($Arguments -join ' ')"
    }
}
