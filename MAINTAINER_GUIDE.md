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

Provide the YAIML reference from the human prompt, workspace context, or a team-approved source at run time. Do not commit machine-specific reference paths, local drive names, user profile paths, `file://` URIs, localhost URLs, private workspace URLs, secrets, private transcripts, or raw sensitive logs into versioned project memory.

Phrases such as "update YAIML", "updated YAIML", "check new YAIML", "refresh YAIML", or "run a YAIML update" mean convention refresh, not an ordinary project-memory rewrite. For this repository, compare local YAIML scaffolding against the provided reference and update only compatible prompts, templates, discovery hints, agent-instruction pointers, or YAIML-maintenance guidance. Preserve project memory and the existing discovery layout unless a human explicitly authorizes a layout migration.

Phrases such as "clean up YAIML", "compress YAIML", "compact project memory", "prune project memory", or "prune SoT" mean to remove stale or repetitive memory while preserving current truth, evidence, human direction, decisions, unresolved conflicts, and uncertainty. Do not pad documents to meet budgets or delete necessary governed knowledge just to reduce word count.

This repository does not need local YAIML prompt or template copies unless it already keeps them for a concrete workflow. Do not add `prompts/` or `templates/` just because the reference has them. Preserve project-specific SoT, Architecture, Maintainer Guide, risks, commands, naming, and supporting documents unless the reference changes how future agents should maintain YAIML here.
