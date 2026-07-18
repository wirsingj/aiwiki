---
yaiml: 0.2
role: sot
title: AIWIKI SOT
purpose: Current engineering state, direction, risks, and priorities for AIWIKI.
belongs-here: product identity, verified capabilities, validation state, known risks, next steps, uncertainty.
not-here: durable architecture, command reference, complete audit history.
durability: volatile; synthesize and prune aggressively.
read-with: Architecture; Maintainer Guide.
update-when: product direction, implementation reality, risks, validation, or priorities change.
agent-guidance: Verify implementation claims against code and tests. Preserve legal/brand caution.
---

# AIWIKI SOT

Updated: 2026-07-16

## Identity

AIWIKI is an independent local demo encyclopedia that generates wiki-style articles with local Ollama, optional hosted OpenAI-compatible providers, or a deterministic demo provider.

The project explicitly positions itself as parody/inspired-by and not affiliated with Wikipedia, Wikimedia, MediaWiki, Ollama, or model providers.

## Verified Capabilities

- Go backend with standard library HTTP.
- Vite/React frontend.
- Provider modes for local Ollama, hosted OpenAI-compatible APIs, and deterministic demo output.
- In-memory article cache with a documented TTL.
- Backend article generation, validation, cache, hosted/demo/Ollama client, and HTTP handler packages.
- Backend article normalization keeps optional generated arrays encoded as arrays instead of `null`.
- Frontend layout, search, article page, article view, infobox, and API client components.
- PowerShell helper scripts for doctor, dev, tests, Docker up/down, and cleanup.

## Active Risks

- Brand/legal posture matters because the product resembles wiki-style encyclopedia UX; preserve explicit non-affiliation language.
- Local model behavior depends on available Ollama models and machine resources.
- Hosted model mode requires secrets through environment variables only.
- Mobile overflow CSS guardrails were added for topbar/search/action/error/footer surfaces, but a screenshot-based visual smoke pass is still pending.
- No database, auth, payments, tracking, or external content fetches are currently present; do not imply production SaaS maturity without new evidence.

## Validation State

Documented verification commands:

```powershell
cd backend
go test ./...

cd frontend
npm run build
```

Latest sanitized validation on 2026-07-16:

- `cd backend; go test ./...` passed.
- `cd frontend; npm run build` passed.
