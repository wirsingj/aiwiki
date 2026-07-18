---
yaiml: 0.2
role: maintainer
title: AIWIKI Maintainer Guide
purpose: Current setup, commands, validation, diagnostics, and YAIML maintenance for AIWIKI.
belongs-here: verified commands, setup notes, diagnostics, important files, danger files, troubleshooting.
not-here: product direction, durable architecture, complete history.
durability: current-only; remove dead commands quickly.
read-with: SOT; Architecture.
update-when: setup, commands, scripts, deployment, or provider configuration changes.
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

Latest sanitized verification on 2026-07-16:

- `cd backend; go test ./...` passed.
- `cd frontend; npm run build` passed.

## Important Files

- `README.md`: project identity, provider setup, and local run guide.
- `STATUS.md`: prior audit/status notes.
- `backend/internal/article/`: generation service and validation.
- `backend/internal/ollama/`, `backend/internal/hosted/`, `backend/internal/demo/`: provider implementations.
- `frontend/src/`: user interface.
- `docs/DEV_WORKFLOW.md`: fuller local workflow notes.
- `DEPLOYMENT.md`, `DEPLOYMENT_AWS.md`: deployment notes.

## YAIML Maintenance

Phrases such as "update YAIML", "updated YAIML", "check new YAIML", or "run a YAIML update" mean: compare this repository's local YAIML setup against a human-provided, workspace-provided, or team-approved YAIML reference, refresh compatible convention scaffolding, and preserve AIWIKI-specific project memory.

Do not commit machine-specific YAIML reference paths, local drive names, user profile paths, `file://` URIs, secrets, private chat transcripts, raw sensitive logs, or invented legal/IP conclusions.
