param(
    [switch]$WithOllama,
    [switch]$Volumes
)

$ErrorActionPreference = "Stop"
. "$PSScriptRoot\common.ps1"

$root = Enter-AiwikiRoot
Write-Host "Stopping AIWIKI Docker stack from $root"

$compose = Get-DockerComposeCommand
$composeArgs = @($compose.Prefix)
if ($WithOllama) {
    $composeArgs += @("-f", "docker-compose.yml", "-f", "docker-compose.ollama.yml")
}
$composeArgs += @("down")
if ($Volumes) {
    $composeArgs += @("-v")
}

Write-Host "$($compose.Label) $($composeArgs[$compose.Prefix.Count..($composeArgs.Count - 1)] -join ' ')"
& $compose.Command @composeArgs
if ($LASTEXITCODE -ne 0) {
    throw "Docker stop failed."
}
