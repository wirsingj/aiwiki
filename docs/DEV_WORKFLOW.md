# AIWIKI Developer Workflow

These commands assume the repo lives at:

```powershell
C:\Users\wirsi\OneDrive\Desktop\git\aiwiki
```

The scripts also work when launched from any directory inside the repo.

## Daily PowerShell Commands

This machine may block direct `.ps1` execution. The most reliable copy-paste form is:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\dev.ps1
```

If your PowerShell allows local scripts, the shorter `.\scripts\dev.ps1` form works too.

Start local dev:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\dev.ps1
```

Stop local dev:

```text
Press Ctrl+C in the dev.ps1 terminal.
```

Run tests/build:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\test.ps1
```

Check your machine:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\doctor.ps1
```

Run Docker stack:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\docker-up.ps1
```

Stop Docker stack:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\docker-down.ps1
```

Run Docker stack with an Ollama container:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\docker-up.ps1 -WithOllama
```

Stop Docker stack with the Ollama override:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\docker-down.ps1 -WithOllama
```

Clean generated artifacts:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\clean.ps1
```

Also remove `node_modules`:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\clean.ps1 -NodeModules
```

## What dev.ps1 Does

- Verifies the repo root.
- Checks Go, Node, npm, Docker if present, and Ollama reachability.
- Starts backend at `http://localhost:8080`.
- Starts frontend at `http://localhost:5173`.
- Opens `http://localhost:5173`.
- Prints backend and frontend logs in the same terminal.
- Handles Ctrl+C by stopping the dev processes.

## Non-Windows Commands

Backend:

```bash
cd backend
go test ./...
go run ./cmd/server
```

Frontend:

```bash
cd frontend
npm install
npm run dev
npm run build
```

Docker:

```bash
docker compose up --build
docker compose down
```

## Common Fixes

- `go` missing: install Go 1.26+.
- `node` or `npm` missing: install Node.js 20+.
- `ollama` missing: install Ollama.
- Ollama offline: run `ollama serve`.
- Model missing: run `ollama pull llama3.1`.
- Docker missing: install Docker Desktop and restart your terminal.
- Port busy: run `.\scripts\doctor.ps1` and stop the process it reports.
