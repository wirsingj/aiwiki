param()

$ErrorActionPreference = "Stop"
. "$PSScriptRoot\common.ps1"

$root = Enter-AiwikiRoot
Write-Host "AIWIKI doctor"
Write-Host "Repo: $root"
Write-Host ""

function Show-CommandVersion {
    param(
        [string]$Name,
        [string[]]$CommandArgs
    )

    if (Get-Command $Name -ErrorAction SilentlyContinue) {
        $version = (& $Name @CommandArgs 2>$null | Select-Object -First 1)
        Write-Host "$Name`: $version"
    } else {
        Write-Host "$Name`: missing"
    }
}

Show-CommandVersion -Name "go" -CommandArgs @("version")
Show-CommandVersion -Name "node" -CommandArgs @("--version")
Show-CommandVersion -Name "npm" -CommandArgs @("--version")

if (Get-Command docker -ErrorAction SilentlyContinue) {
    Write-Host "docker: $(& docker --version)"
    try {
        $compose = Get-DockerComposeCommand
        Write-Host "docker compose: $(& $compose.Command @($compose.Prefix + @('version')))"
    } catch {
        Write-Host "docker compose: missing"
    }
    try {
        docker info *> $null
        Write-Host "docker daemon: running"
    } catch {
        Write-Host "docker daemon: not running or unavailable"
    }
} else {
    Write-Host "docker: missing"
}

Write-Host ""
if (Get-Command ollama -ErrorAction SilentlyContinue) {
    Write-Host "ollama: installed"
    try {
        $tags = Invoke-RestMethod -Uri "http://localhost:11434/api/tags" -TimeoutSec 3
        Write-Host "ollama status: reachable at http://localhost:11434"
        if ($tags.models.Count -gt 0) {
            Write-Host "ollama models:"
            $tags.models | ForEach-Object { Write-Host "  - $($_.name)" }
        } else {
            Write-Host "ollama models: none detected"
        }
    } catch {
        Write-Host "ollama status: not reachable at http://localhost:11434"
    }
} else {
    Write-Host "ollama: missing"
}

Write-Host ""
if (Test-PortListening 8080) {
    Write-Host "backend port 8080: listening"
    Show-PortOwner 8080
    try {
        $health = Invoke-RestMethod -Uri "http://localhost:8080/api/health" -TimeoutSec 3
        Write-Host "backend health: $($health | ConvertTo-Json -Compress)"
    } catch {
        Write-Host "backend health: no AIWIKI health response"
    }
} else {
    Write-Host "backend port 8080: not listening"
}

if (Test-PortListening 5173) {
    Write-Host "frontend URL: http://localhost:5173"
    Show-PortOwner 5173
} else {
    Write-Host "frontend port 5173: not listening"
}

Write-Host ""
Write-Host "Common fixes:"
Write-Host "- Missing npm packages: cd frontend; npm install"
Write-Host "- Ollama offline: ollama serve"
Write-Host "- Missing model: ollama pull llama3.1"
Write-Host "- Port busy: run scripts\doctor.ps1 to see owner, then stop that app"
Write-Host "- Docker daemon unavailable: start Docker Desktop; if it is not installed, install it with admin approval"
Write-Host "- Docker command missing: install Docker CLI/Desktop, then restart your terminal"
