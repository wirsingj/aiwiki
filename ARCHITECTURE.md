---
yaiml: 0.2
role: architecture
title: AIWIKI Architecture
purpose: Durable system shape, boundaries, and invariants for AIWIKI.
belongs-here: components, provider boundaries, data flow, deployment shape, invariants, danger zones.
not-here: volatile priorities, command reference, full work history.
durability: stable; update when module boundaries, provider contracts, deployment, or runtime surfaces change.
budget: About 700 words; a working target, not a length to fill.
read-with: SOT; Maintainer Guide.
update-when: backend/frontend responsibilities, provider architecture, cache behavior, or deployment topology changes.
last-verified: not established; claims not rechecked in this refresh.
agent-guidance: Preserve provider abstraction and brand-safety boundaries.
---

# AIWIKI Architecture

## Components

- `backend/cmd/server/main.go`: backend entrypoint and server wiring.
- `backend/internal/article/`: article request/response types, validation, and generation service.
- `backend/internal/ollama/`: local Ollama client.
- `backend/internal/hosted/`: hosted OpenAI-compatible provider client.
- `backend/internal/demo/`: deterministic demo provider.
- `backend/internal/cache/`: in-memory cache behavior.
- `backend/internal/http/`: HTTP handlers and rate limiting.
- `frontend/src/`: Vite/React app, article pages, search, layout, infobox, and API client.
- `scripts/`: local PowerShell workflows.

## Data Flow

The frontend requests generated article data from the backend API. The backend validates topics, selects the configured provider mode, generates structured article content, caches recent article responses, and returns structured fields for React rendering.

## Invariants

- Generated article body should be rendered as text/structured fields, not raw HTML.
- Hosted provider secrets belong in environment variables, not source.
- Brand/non-affiliation language must remain visible in public docs.
- Provider-specific behavior should remain behind local, hosted, or demo client boundaries.
- The default local demo path should remain usable without a database.

## Danger Zones

- Weakening generated-content validation or rendering untrusted HTML.
- Blurring parody/non-affiliation posture.
- Hardcoding local machine paths into YAIML project memory.
- Treating a no-auth local demo as production-ready without explicit new evidence.
