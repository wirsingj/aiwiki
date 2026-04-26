param(
    [switch]$WithOllama
)

$ErrorActionPreference = "Stop"
. "$PSScriptRoot\common.ps1"

$root = Enter-AiwikiRoot
Write-Host "Starting AIWIKI Docker stack from $root"

$compose = Get-DockerComposeCommand
$composeArgs = @($compose.Prefix)
if ($WithOllama) {
    $composeArgs += @("-f", "docker-compose.yml", "-f", "docker-compose.ollama.yml")
}
$composeArgs += @("up", "--build")

Write-Host "$($compose.Label) $($composeArgs[$compose.Prefix.Count..($composeArgs.Count - 1)] -join ' ')"
& $compose.Command @composeArgs
if ($LASTEXITCODE -ne 0) {
    throw "Docker stack failed."
}
