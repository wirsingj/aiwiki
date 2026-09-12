---
yaiml: 0.2
role: maintainer
title: AIWIKI Maintainer Guide
purpose: Current setup, commands, validation, diagnostics, and YAIML maintenance for AIWIKI.
belongs-here: verified commands, setup notes, diagnostics, important files, danger files, troubleshooting.
not-here: product direction, durable architecture, complete history.
durability: current-only; remove dead commands quickly.
budget: About 900 words; a working target, not a length to fill.
read-with: SOT; Architecture.
update-when: setup, commands, scripts, deployment, or provider configuration changes.
last-verified: not established; claims not rechecked in this refresh.
agent-guidance: Record sanitized command outcomes. Do not record secrets or local-only provider state.
---

# AIWIKI Maintainer Guide

## Setup

Backend:

```powershell
cd backend
go run ./cmd/server
```

Frontend:

```powershell
cd frontend
npm install
npm run dev
```

## Verified Commands

Backend tests:

```powershell
cd backend
go test ./...
```

Frontend build:

```powershell
cd frontend
npm run build
```

PowerShell helper:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\test.ps1
```

See [SOT.md](SOT.md) for recorded validation outcomes.

## Important Files

- `README.md`: project identity, provider setup, and local run guide.
- `STATUS.md`: prior audit/status notes.
- `backend/internal/article/`: generation service and validation.
- `backend/internal/ollama/`, `backend/internal/hosted/`, `backend/internal/demo/`: provider implementations.
- `frontend/src/`: user interface.
- `docs/DEV_WORKFLOW.md`: fuller local workflow notes.
- `DEPLOYMENT.md`, `DEPLOYMENT_AWS.md`: deployment notes.

## YAIML Maintenance

Follow the synthesis, privacy, and retention rules in [AGENTS.md](AGENTS.md). Ordinary material updates replace changed facts and remove superseded or resolved state; keep one detailed home per fact. Measure affected memory before and after, compress safely first, and report necessary growth or retained budget overages in the task response.

For convention refresh requests (including "refresh YAIML"), obtain the human-provided, workspace-local, or team-approved reference at run time; request one if unavailable. Compare its update and init guidance with local instructions and maintenance notes. Keep private reference locations out of versioned files. Preserve project knowledge, local names, discovery layout/version, paths, and meaningful custom fields; layout migration requires explicit human authorization and consumer checks. Do not add prompt/template copies without a concrete local workflow.

For cleanup requests, remove stale or repetitive memory while preserving direction, evidence limits, decisions, uncertainty, unresolved conflicts, and governed records. Do not relocate history or delete necessary knowledge to meet budgets. Verify affected links, discovery paths, stable headers, and instruction scope; rerunning the same refresh should leave healthy files unchanged. Report configured persistence separately from observed fresh-session loading.
