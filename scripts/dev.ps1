param()

$ErrorActionPreference = "Stop"
. "$PSScriptRoot\common.ps1"

$root = Enter-AiwikiRoot
Write-Host "Starting AIWIKI dev stack from $root"

Require-Command "go" "Install Go 1.26+."
Require-Command "node" "Install Node.js 20+."
Require-Command "npm" "Install npm with Node.js."

if (Get-Command docker -ErrorAction SilentlyContinue) {
    Write-Host "Docker: $(& docker --version)"
} else {
    Write-Host "Docker: not found; OK for non-Docker dev."
}

if (-not (Get-Command ollama -ErrorAction SilentlyContinue)) {
    Write-Host "Warning: ollama command not found. Backend can start, but generation will fail until Ollama is available."
} else {
    try {
        Invoke-RestMethod -Uri "http://localhost:11434/api/tags" -TimeoutSec 3 | Out-Null
        Write-Host "Ollama: reachable at http://localhost:11434"
    } catch {
        Write-Host "Warning: Ollama is not reachable at http://localhost:11434. Try: ollama serve"
    }
}

if (Test-PortListening 8080) {
    Show-PortOwner 8080
    throw "Port 8080 is busy. Stop that process before running dev.ps1."
}
if (Test-PortListening 5173) {
    Show-PortOwner 5173
    throw "Port 5173 is busy. Stop that process before running dev.ps1."
}

Load-DotEnv "$root\backend\.env"
Load-DotEnv "$root\.env"
if (-not $env:PORT) { $env:PORT = "8080" }
if (-not $env:OLLAMA_BASE_URL) { $env:OLLAMA_BASE_URL = "http://localhost:11434" }
if (-not $env:OLLAMA_MODEL) { $env:OLLAMA_MODEL = "llama3.1" }

Push-Location "$root\frontend"
try {
    if (-not (Test-Path "node_modules")) {
        Write-Host "Installing frontend dependencies..."
        Invoke-Checked -Command "npm" -Arguments @("install")
    }
} finally {
    Pop-Location
}

$backendJob = $null
$frontendJob = $null

try {
    Write-Host ""
    Write-Host "Launching backend on http://localhost:8080"
    $backendJob = Start-Job -Name "aiwiki-backend" -ScriptBlock {
        param($RepoRoot)
        Set-Location "$RepoRoot\backend"
        go run .\cmd\server 2>&1
    } -ArgumentList $root

    Write-Host "Launching frontend on http://localhost:5173"
    $frontendJob = Start-Job -Name "aiwiki-frontend" -ScriptBlock {
        param($RepoRoot)
        Set-Location "$RepoRoot\frontend"
        npm run dev -- --host 127.0.0.1 --port 5173 --strictPort 2>&1
    } -ArgumentList $root

    $opened = $false
    Write-Host ""
    Write-Host "Logs follow. Press Ctrl+C to stop backend and frontend."
    Write-Host ""

    while ($true) {
        foreach ($job in @($backendJob, $frontendJob)) {
            if ($job) {
                Receive-Job $job | ForEach-Object {
                    Write-Host "[$($job.Name)] $_"
                }
            }
        }

        if (-not $opened -and (Test-PortListening 5173)) {
            Start-Process "http://localhost:5173"
            $opened = $true
        }

        $failed = @($backendJob, $frontendJob) | Where-Object { $_ -and $_.State -in @("Failed", "Stopped", "Completed") }
        if ($failed.Count -gt 0) {
            foreach ($job in $failed) {
                Receive-Job $job | ForEach-Object { Write-Host "[$($job.Name)] $_" }
                Write-Host "$($job.Name) exited with state $($job.State)."
            }
            throw "Dev stack stopped unexpectedly."
        }

        Start-Sleep -Milliseconds 500
    }
} finally {
    Write-Host ""
    Write-Host "Stopping AIWIKI dev stack..."
    foreach ($job in @($backendJob, $frontendJob)) {
        if ($job) {
            Stop-Job $job -ErrorAction SilentlyContinue
            Remove-Job $job -Force -ErrorAction SilentlyContinue
        }
    }

    foreach ($port in @(8080, 5173)) {
        $conn = Get-NetTCPConnection -LocalPort $port -State Listen -ErrorAction SilentlyContinue | Select-Object -First 1
        if ($conn) {
            $proc = Get-CimInstance Win32_Process -Filter "ProcessId=$($conn.OwningProcess)" -ErrorAction SilentlyContinue
            if ($proc -and ($proc.CommandLine -match "aiwiki|cmd\\server|vite|go run|npm")) {
                Stop-ProcessTree -ProcessId $conn.OwningProcess
            }
        }
    }
    Write-Host "Stopped."
}
