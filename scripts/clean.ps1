param(
    [switch]$NodeModules,
    [switch]$Audit
)

$ErrorActionPreference = "Stop"
. "$PSScriptRoot\common.ps1"

$root = Enter-AiwikiRoot
Write-Host "Cleaning AIWIKI build artifacts from $root"

$targets = @(
    "$root\frontend\dist",
    "$root\frontend\tsconfig.tsbuildinfo",
    "$root\backend\server.exe",
    "$root\backend\backend.out.log",
    "$root\backend\backend.err.log",
    "$root\backend\backend.log"
)

if ($Audit) {
    $targets += "$root\.audit"
}

if ($NodeModules) {
    $targets += "$root\frontend\node_modules"
}

foreach ($target in $targets) {
    if (Test-Path $target) {
        Write-Host "Removing $target"
        Remove-Item -LiteralPath $target -Recurse -Force
    }
}

Write-Host "Clean complete."
if (-not $NodeModules) {
    Write-Host "node_modules was preserved. Use -NodeModules to remove it."
}
