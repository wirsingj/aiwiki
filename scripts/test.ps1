param()

$ErrorActionPreference = "Stop"
. "$PSScriptRoot\common.ps1"

$root = Enter-AiwikiRoot
Write-Host "Running AIWIKI tests from $root"

Require-Command "go" "Install Go 1.26+."
Require-Command "node" "Install Node.js 20+."
Require-Command "npm" "Install npm with Node.js."

Write-Host ""
Write-Host "== Backend: go test ./... =="
Push-Location "$root\backend"
try {
    Invoke-Checked -Command "go" -Arguments @("test", "./...")
} finally {
    Pop-Location
}

Write-Host ""
Write-Host "== Frontend: npm install if needed =="
Push-Location "$root\frontend"
try {
    if (-not (Test-Path "node_modules")) {
        Invoke-Checked -Command "npm" -Arguments @("install")
    }

    Write-Host ""
    Write-Host "== Frontend: npm run build =="
    Invoke-Checked -Command "npm" -Arguments @("run", "build")

    $package = Get-Content "package.json" | ConvertFrom-Json
    $hasTest = $package.scripts.PSObject.Properties.Name -contains "test"
    if ($hasTest) {
        Write-Host ""
        Write-Host "== Frontend: npm test =="
        Invoke-Checked -Command "npm" -Arguments @("test")
    } else {
        Write-Host ""
        Write-Host "== Frontend: no npm test script present; skipping =="
    }
} finally {
    Pop-Location
}

Write-Host ""
Write-Host "All checks passed."
