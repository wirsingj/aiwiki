# AIWIKI Status

Last audited: 2026-07-11

## Current State

- Project root is `C:\Users\wirsi\OneDrive\Desktop\git\aiwiki`.
- Canonical repo/project spelling is `aiwiki` / `AIWIKI`.
- The directory is currently initialized as a Git repository on branch `main`; `git status --short` was clean on 2026-07-11.
- No old project-name typo was found in source, docs, package metadata, or Go module paths outside ignored build/dependency output.
- Backend structure matches the requested layout:
  - `backend/cmd/server/main.go`
  - `backend/internal/ollama/client.go`
  - `backend/internal/article/types.go`
  - `backend/internal/article/service.go`
  - `backend/internal/cache/cache.go`
  - `backend/internal/http/handlers.go`
- Frontend structure matches the requested layout:
  - `frontend/src/App.tsx`
  - `frontend/src/pages/Home.tsx`
  - `frontend/src/pages/ArticlePage.tsx`
  - `frontend/src/components/SearchBar.tsx`
  - `frontend/src/components/ArticleView.tsx`
  - `frontend/src/components/Infobox.tsx`
  - `frontend/src/components/Layout.tsx`
  - `frontend/src/api.ts`

## Verification Run

Backend tests passed:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki\backend
go test ./...
```

Frontend production build passed:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki\frontend
npm run build
```

Live backend health check passed:

```json
{
  "ok": true,
  "ollama": "online",
  "ollamaModel": "llama3.2:3b",
  "service": "aiwiki-backend"
}
```

Cache behavior was probed through the API:

- First `GET /api/article/cache-audit-topic?refresh=1`: `cached=false`.
- Next `GET /api/article/cache-audit-topic`: `cached=true` and same `generatedAt`.
- Next `GET /api/article/cache-audit-topic?refresh=1`: `cached=false` and new `generatedAt`.

Malformed JSON fallback tests currently cover:

- JSON extracted from markdown fences.
- JSON extracted from surrounding prose.

## UI Inspection Notes

- The in-app browser plugin failed to initialize in this session, so UI inspection used local HTTP checks plus headless Edge screenshots.
- `http://localhost:5173` is currently occupied by another local repo (`chaitroom`) on this machine.
- AIWIKI frontend was temporarily started on `http://127.0.0.1:5183` for visual inspection.
- Home desktop layout looks clean and recognizable as AIWIKI.
- Mobile layout has horizontal overflow in the top search/header area.
- The earlier temporary-port CORS issue has been mitigated: Vite dev now proxies `/api`, and production frontend uses same-origin Nginx proxying by default.

## Security Posture

- Generated article body is rendered as React text, not raw HTML.
- No `dangerouslySetInnerHTML` usage was found.
- Generated article links are rendered from structured fields.
- Backend response body is limited to 2 MB when reading Ollama responses.
- CORS is configurable through `AIWIKI_ALLOWED_ORIGINS`; production should prefer same-origin `/api` proxying.
- No auth, payments, tracking, external content fetches, or database access are present.
