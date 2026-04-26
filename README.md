# AIWIKI

AIWIKI is an independent local demo encyclopedia that generates wiki-style articles with a local Ollama model. It is a parody/inspired-by layout, not a Wikimedia project, and it does not use Wikipedia branding, logos, or scraped content. AIWIKI is not affiliated with Wikipedia, the Wikimedia Foundation, MediaWiki, Ollama, or any model provider.

## Stack

- Go backend with standard library HTTP
- Vite + React frontend
- Ollama local API at `http://localhost:11434`
- Optional hosted OpenAI-compatible model provider for public demos
- Instant template demo provider for no-model deployments
- In-memory article cache with a 10 minute TTL
- No database

## Prerequisites

- Go 1.26+
- Node.js 20+
- Ollama installed and running

## Ollama Setup

Install Ollama from `https://ollama.com`, then pull a model:

```powershell
ollama pull llama3.1
```

Start Ollama if it is not already running:

```powershell
ollama serve
```

The backend reads:

- `PORT`, default `8080`
- `AIWIKI_MODEL_PROVIDER`, default `ollama`; options: `ollama`, `hosted`, `demo`
- `OLLAMA_BASE_URL`, default `http://localhost:11434`
- `OLLAMA_MODEL`, default `llama3.1`
- `AIWIKI_GENERATION_MODE`, default `pipeline`
- `AIWIKI_GENERATION_CONCURRENCY`, default `2`
- `OLLAMA_WORKER_BASE_URLS`, optional comma-separated Ollama workers for parallel section generation
- `HOSTED_MODEL_BASE_URL`, default `https://api.openai.com/v1`
- `HOSTED_MODEL_NAME`, default `gpt-4o-mini`
- `HOSTED_MODEL_API_KEY`, required only when `AIWIKI_MODEL_PROVIDER=hosted`

If `OLLAMA_MODEL` is not set and Ollama is reachable at startup, the backend will prefer an installed `llama3.1` variant and otherwise use the first installed model it sees.

## Run the Backend

From the project root:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki\backend
go run .\cmd\server
```

The backend runs at `http://localhost:8080`.

Useful endpoints:

- `GET http://localhost:8080/api/health`
- `GET http://localhost:8080/api/article/alan-turing`
- `POST http://localhost:8080/api/article` with body `{ "topic": "Alan Turing" }`
- `GET http://localhost:8080/api/random`

## Run the Frontend

In another terminal:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki\frontend
npm install
npm run dev
```

Open `http://localhost:5173`.

The Vite dev server proxies `/api` to `http://localhost:8080`. If port `5173` is busy, use another port:

```powershell
npm run dev -- --port 5183
```

For unusual setups where the browser should call a separate API origin directly, set `VITE_API_BASE_URL`.

## Run With Docker

Docker expects Ollama to already be running and reachable. For local Docker Desktop, the default Compose config points the backend at Ollama on the host through `host.docker.internal`.

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki
copy .env.example .env
docker compose up --build
```

Open `http://localhost:8088`.

Docker services:

- Frontend: `http://localhost:8088`
- Backend: `http://localhost:8080`
- Backend health: `http://localhost:8080/api/health`

To run Ollama as part of the Compose stack:

```powershell
docker compose -f docker-compose.yml -f docker-compose.ollama.yml up -d ollama
docker compose -f docker-compose.yml -f docker-compose.ollama.yml exec ollama ollama pull llama3.1
docker compose -f docker-compose.yml -f docker-compose.ollama.yml up --build
```

Build individual images:

```powershell
docker build -t aiwiki-backend:local .\backend
docker build -t aiwiki-frontend:local .\frontend
```

For AWS/container deployment notes, see `DEPLOYMENT.md`.
For a no-cost public AWS demo strategy, see `docs/DEV_WORKFLOW.md` and `docs/AWS_FREE_DEMO.md`.
For the monetized AWS/ad-readiness plan, see `docs/MONETIZED_AWS_LAUNCH.md` and `docs/PUBLIC_DEMO_RUNBOOK.md`.

## Test and Build

PowerShell helpers:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\doctor.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\dev.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\test.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\docker-up.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\docker-down.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\clean.ps1
```

Full workflow docs: `docs/DEV_WORKFLOW.md`.

Backend:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki\backend
go test ./...
```

Frontend:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki\frontend
npm run build
```

Non-Windows equivalents:

```bash
cd backend && go test ./...
cd backend && go run ./cmd/server
cd frontend && npm install && npm run dev
cd frontend && npm run build
docker compose up --build
docker compose down
```

## Notes

- Articles include an on-page disclaimer: "AI-generated article. May contain inaccuracies. Verify important claims."
- Generated references are explicitly AI-generated notes, not real sources or verified citations.
- Users are responsible for verifying important claims before relying on generated content.
- Users are responsible for complying with the license and acceptable-use terms for any model they run through Ollama or another provider.
- Internal links are rendered from structured article JSON and route to `/wiki/<slug>`.
- The "Regenerate article" button bypasses the 10 minute cache.
- The frontend renders generated text safely as React text, without `dangerouslySetInnerHTML`.
