# AIWIKI Legal and Public-Repo Readiness Audit

Date: 2026-04-26

This is an engineering/compliance preflight, not legal advice. It separates confirmed repo facts from likely risks and items that need owner or attorney review before a public launch.

## Sources Checked

- Wikimedia Foundation Trademark Policy: https://foundation.wikimedia.org/wiki/Trademark_policy
- Ollama license: https://github.com/ollama/ollama/blob/main/LICENSE
- Ollama Terms of Service: https://ollama.com/terms
- Ollama Llama 3.1 library page: https://ollama.com/library/llama3.1
- Meta Llama 3.1 Community License: https://github.com/meta-llama/llama-models/blob/main/models/llama3_1/LICENSE
- Meta Llama 3.1 Acceptable Use Policy: https://github.com/meta-llama/llama-models/blob/main/models/llama3_1/USE_POLICY.md

## Confirmed Repo Facts

- Project name and repo name are consistently `AIWIKI` / `aiwiki`.
- No old project-name typo/path drift was found in source files, docs, package names, or module paths.
- No combined AI-plus-Wikipedia project name was found.
- The backend generates articles by calling an Ollama-compatible API. It does not scrape Wikipedia or call Wikimedia APIs.
- No stored article corpus, sample Wikipedia article text, seed article dataset, or copied encyclopedia content was found.
- The frontend renders generated/user text as React text nodes. No `dangerouslySetInnerHTML` usage was found.
- The app uses an `A` brand mark and text `AIWIKI`; it does not include the Wikipedia puzzle globe, Wikimedia logos, or MediaWiki marks.
- Go backend uses only the standard library. Frontend dependencies are listed in `THIRD_PARTY_NOTICES.md`.
- No obvious secrets, API keys, tokens, passwords, or private keys were found by regex scan.
- `.env` files, logs, build output, TypeScript build info, and `.audit/` are ignored.
- This directory is not currently an initialized Git repository, so tracked/untracked release status could not be verified with `git status`.

## Branding and Trademark

### Confirmed

- The README and app now say AIWIKI is independent and unaffiliated with Wikipedia, Wikimedia Foundation, MediaWiki, Ollama, or model providers.
- The app avoids Wikipedia/Wikimedia logos and exact brand phrases.
- The repo does mention Wikipedia/Wikimedia/MediaWiki in disclaimers and audit documents only.

### Likely Risks

- The Wikimedia trademark policy states that Wikimedia marks include wordmarks and site trade dress. The FAQ also discusses trade dress as the look and feel of a site or article that identifies source. AIWIKI intentionally uses a classic encyclopedia/wiki layout, so the safest posture is to keep the layout inspired and generic, not a pixel-faithful clone.
- The term `AIWIKI` includes `wiki`. That is probably safer than using a Wikimedia project name or a combined AI-plus-Wikipedia name, but an attorney should evaluate public branding and domain choices if this becomes more than a small demo.

### Safer Wording

Use:

- "independent AI-generated encyclopedia demo"
- "inspired by classic wiki/encyclopedia layouts"
- "generated notes are not verified citations"

Avoid:

- "clone" wording tied to Wikipedia
- "Wikipedia AI"
- "AIWikipedia"
- "The Free Encyclopedia"
- Any Wikimedia logo, puzzle globe, stylized wordmark, or close visual imitation.

## Copyright and Generated Content

### Confirmed

- No Wikipedia scraping or copied article text was found.
- Generated content is created from a model response at request time and cached in memory for 10 minutes.
- References are generated notes, not real citations.

### Risks

- Model outputs can be inaccurate, defamatory, infringing, or too similar to training data or third-party content.
- Users may mistake generated notes for verified references if the wording becomes too source-like.

### Mitigations Already Present

- Visible article banner warns that generated content may be inaccurate.
- Prompts instruct the model not to fabricate URLs, real citations, books, papers, article names, publication names, or named source attributions.
- App footer and About page now include independence and verification disclaimers.

## Ollama and Model Usage

### Confirmed

- AIWIKI calls an Ollama-compatible API through `OLLAMA_BASE_URL`, defaulting to `http://localhost:11434`.
- `OLLAMA_MODEL` defaults to `llama3.1`, but the backend may auto-discover a local installed model if the env var is unset.
- The normal app containers do not bundle Ollama or model weights.
- `docker-compose.ollama.yml` can run `ollama/ollama:latest` as an optional local service and stores model data in a Docker volume.

### Terms and License Notes

- Ollama's repository license is MIT, but Ollama's hosted/services terms separately govern use of Ollama services.
- Ollama's terms say users must use services lawfully, respect third-party rights, and verify AI outputs before use.
- Llama 3.1 is under the Meta Llama 3.1 Community License, not a plain MIT/Apache-style open-source license.
- The Llama 3.1 license includes attribution/notice obligations when distributing or making available products/services that contain Llama Materials, acceptable-use obligations, commercial scale terms, warranty disclaimers, and trademark limits.

### Release Guidance

- Do not imply AIWIKI bundles Ollama, Llama, or rights to any model.
- Keep docs saying users must comply with the license and acceptable-use terms for the model they run.
- If a public deployment runs Llama 3.1 or another Meta Llama model behind the app, verify whether UI/docs should display the required "Built with Llama" notice and include the applicable license notice.
- If switching to OpenAI, OpenRouter, or another hosted model provider, update docs, notices, privacy expectations, and terms.

## Dependency Licenses

See `THIRD_PARTY_NOTICES.md`.

Summary:

- No GPL or AGPL packages were found in the current dependency report.
- Frontend dependencies are primarily MIT/Apache-2.0/ISC/BSD-3-Clause.
- `lightningcss` and platform packages are MPL-2.0. This is usually manageable for app usage, but it is a file-level copyleft license and should remain in notices.
- A few optional/native package licenses were missing from the installed `package-lock.json` metadata; npm registry checks resolved them to MIT, MPL-2.0, or 0BSD as documented in `THIRD_PARTY_NOTICES.md`.

## Public Repo Safety

### Confirmed

- No hardcoded API keys, AWS credentials, tokens, passwords, or private keys were found by regex scan.
- `.env` is ignored.
- `.env.example` files contain non-secret defaults only.
- Dockerfiles and Compose files do not contain secrets.

### Items to Clean Before Push

- Local generated artifacts currently exist but are ignored:
  - `backend/backend.err.log`
  - `backend/backend.out.log`
  - `frontend/tsconfig.tsbuildinfo`
- Run `powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\clean.ps1` before release.
- Initialize Git before the final push and verify `git status --short` does not include logs, `.env`, build output, or dependency folders.

## Naming Risk

`AIWIKI` is probably acceptable for a small independent demo if the app keeps the current disclaimers and avoids Wikimedia marks. The risk is not zero because the product is encyclopedia-like and uses "wiki" in its name. A name that combines AI with the Wikipedia mark would be much riskier and was not found in the repo.

Safer alternatives if this grows commercially:

- `AIGlossary`
- `Synthopedia`
- `LocalLexicon`
- `Generated Atlas`

No rename is recommended for v0 without attorney input.

## Needs Attorney Review

- Whether `AIWIKI` is acceptable as a public/commercial brand and domain.
- Whether the visual design is sufficiently distinct from Wikimedia trade dress for the target jurisdiction and deployment.
- Whether generated article outputs create defamation, copyright, or consumer-protection obligations for a public site.
- Which repo license best matches the owner's business goals.
- Whether the actual deployed model requires visible model attribution or additional notices.

## Needs Owner Verification

- Current Ollama terms and Docker image terms at release time.
- Current license and acceptable-use policy for the exact model served in production.
- Whether public traffic will send user prompts to a third-party hosted model provider.
- Whether screenshots, demos, or README media contain copied browser content, bookmarks, private paths, or private user data.
