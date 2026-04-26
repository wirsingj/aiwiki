# AIWIKI Risks

## P1 - Mobile Layout Overflows

At 390px wide, the top search bar and some page text overflow horizontally. The search button is partially clipped, and the article error heading can be cut off. This makes the demo feel broken on phones.

Likely areas: `frontend/src/styles.css`, especially `.topbar`, `.search`, `.article-shell`, and mobile media queries.

## P1 - Frontend Can Crash On Valid But Incomplete Article JSON

The backend only requires `title` and `summary` for parsed article JSON. If Ollama returns valid JSON that omits `sections`, `references`, or `seeAlso`, Go may encode nil slices as `null`, while the frontend expects arrays and calls `.length` / `.map`.

Likely areas:

- `backend/internal/article/service.go`
- `frontend/src/components/ArticleView.tsx`

## P3 - Default Frontend Port Collision Needs Awareness

On this machine, `5173` was serving another repo during audit. The app now handles alternate Vite ports better because dev requests proxy `/api`, and production uses same-origin Nginx proxying. Keep documenting alternate ports so this stays easy to run.

Likely areas:

- `frontend/vite.config.ts`
- `backend/internal/http/handlers.go`
- `README.md`

## P1 - Ollama Failure UX Is Too Generic

When the browser cannot reach the backend or CORS blocks the request, the UI shows `Failed to fetch`. When Ollama itself fails, the backend can return useful text, but the frontend does not distinguish offline backend, offline Ollama, timeout, malformed JSON, or blocked origin.

Likely areas:

- `frontend/src/api.ts`
- `frontend/src/pages/ArticlePage.tsx`
- `backend/internal/ollama/client.go`

## P2 - Malformed JSON Handling Needs Broader Tests

Current tests cover fenced JSON and embedded JSON extraction. They do not cover the full service-level repair path, missing arrays, invalid field types, empty references, or JSON that is syntactically valid but structurally unsafe for the frontend.

Likely areas:

- `backend/internal/article/service_test.go`
- `backend/internal/article/service.go`

## P2 - Unicode Slugs Are Only Partially Safe

Slugify supports Unicode letters and digits, but `TopicFromSlug` uses byte slicing for capitalization. A slug beginning with a multi-byte rune could produce malformed display text or odd prompts.

Likely area: `backend/internal/article/service.go`

## P2 - Factual Trust Still Needs Tightening

The banner and reference note disclaimer are good, but the UI still uses a plain `References` heading and an infobox that can look authoritative. The article prose itself can also contain real-sounding factual claims even when references are clearly labeled as generated notes.

Likely areas:

- `frontend/src/components/ArticleView.tsx`
- `backend/internal/ollama/client.go`

## P2 - Random Article Has No Frontend Error/Loading State

The sidebar `Random article` button awaits a generated article before navigating, but it has no disabled, loading, or error state. If Ollama is slow or offline, nothing obvious happens.

Likely area: `frontend/src/components/Layout.tsx`

## P2 - Regeneration Drops The Current Article Immediately

Clicking `Regenerate article` switches the whole page to the loading state. This is functional, but it removes the current article while waiting and can feel jumpy for slow local models.

Likely area: `frontend/src/pages/ArticlePage.tsx`

## P3 - Reference Footnote Numbering Is Approximate

Each section links to `Math.min(sectionIndex + 1, references.length)`. If there are fewer references than sections, later sections repeat the last reference number. This is acceptable for v0 but visually odd.

Likely area: `frontend/src/components/ArticleView.tsx`

## P3 - Duplicate Section Headings Can Duplicate Anchors

TOC anchors are generated from section headings. Duplicate headings will create duplicate `id` attributes.

Likely area: `frontend/src/components/ArticleView.tsx`

## P3 - Local Demo Has No Request Size Guard On POST

Ollama responses are size-limited when read, but inbound request bodies are not explicitly limited. This is low risk for a local-only demo, but easy to harden.

Likely area: `backend/internal/http/handlers.go`
