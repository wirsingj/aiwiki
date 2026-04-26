# AIWIKI Public Release Checklist

Use this before pushing AIWIKI to GitHub or deploying a public demo.

## Required Before Push

- [ ] Choose a repo license. Do not imply the repo is open source until a license is added.
- [ ] Run `powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\test.ps1`.
- [ ] Run `powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\clean.ps1`.
- [ ] Initialize Git if needed and run `git status --short`.
- [ ] Confirm no `.env`, logs, build output, dependency folders, cache files, or private screenshots are tracked.
- [ ] Run a secrets scan such as:

```powershell
rg -n "sk-[A-Za-z0-9]|OPENAI_API_KEY|OPENROUTER_API_KEY|AWS_ACCESS_KEY|AWS_SECRET|TOKEN=|SECRET=|PASSWORD=|PRIVATE KEY|BEGIN .*KEY" .
```

- [ ] Review `README.md`, `SECURITY.md`, `LEGAL_AUDIT.md`, `LICENSE_RECOMMENDATION.md`, and `THIRD_PARTY_NOTICES.md`.
- [ ] Confirm all public-facing pages say generated content may be inaccurate and that AIWIKI is unaffiliated with Wikipedia/Wikimedia/MediaWiki/Ollama.

## Branding and Content

- [ ] Do not use Wikipedia, Wikimedia, MediaWiki, or Ollama logos.
- [ ] Do not use the Wikipedia puzzle globe, stylized Wikipedia wordmark, "The Free Encyclopedia", or Wikimedia visual identity assets.
- [ ] Avoid "clone" wording tied to Wikipedia in public marketing copy.
- [ ] Do not commit generated article samples unless reviewed for copied text, defamation, and false factual claims.
- [ ] If using screenshots, crop browser bookmarks, accounts, private paths, and personal data.

## Models and AI

- [ ] Verify the current license and acceptable-use policy for the exact model used in production.
- [ ] If using Llama 3.1 or another Meta Llama model publicly, verify whether "Built with Llama" or other notices are required.
- [ ] Do not expose Ollama directly to the public internet.
- [ ] Confirm public deployment has rate limiting outside the app container.
- [ ] Confirm generated reference notes cannot be mistaken for real citations.

## Deployment

- [ ] Set production `AIWIKI_ALLOWED_ORIGINS` to the real HTTPS origin only.
- [ ] Keep frontend and backend same-origin through the nginx `/api` proxy when possible.
- [ ] Use HTTPS through the hosting provider or load balancer.
- [ ] Keep model-provider secrets out of Dockerfiles, Compose files, and public docs.
- [ ] If Docker is available, run `docker compose up --build` and verify `/api/health`.
- [ ] Add AWS budget alerts before exposing the site publicly.
- [ ] Keep Ollama private; never expose the Ollama port directly to the internet.
- [ ] Start with the cheapest provider mode that meets the demo goal, then scale model spend only after traffic/revenue justify it.

## Ads and Monetization

- [ ] Publish About, Contact, Privacy Policy, Terms/Disclaimer, and AI Content Policy pages before applying to an ad provider.
- [ ] Keep real ads disabled until the site is approved and policy pages are live.
- [ ] Do not place ads where users may mistake them for navigation, search controls, citations, downloads, or retry buttons.
- [ ] Do not click your own ads or buy low-quality traffic.
- [ ] Track AWS/model/domain costs against ad revenue at least weekly.
- [ ] Complete ad-provider payment, identity, address, and tax steps before expecting payout.

## Dependency Hygiene

- [ ] Run dependency/license review after every dependency update.
- [ ] Confirm no GPL, AGPL, unlicensed, or unexpected commercial-license packages are introduced.
- [ ] Update `THIRD_PARTY_NOTICES.md` when dependency versions change.
- [ ] Run vulnerability checks appropriate to the release workflow, such as `npm audit` and `govulncheck` if installed.

## Attorney / Owner Signoff

- [ ] Attorney review for AIWIKI name/domain if used commercially.
- [ ] Attorney review for Wikimedia trade-dress risk if public design becomes more exact.
- [ ] Owner approval for repo license.
- [ ] Owner approval for model/provider terms and attribution obligations.
