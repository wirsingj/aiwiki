# AIWIKI Monetized AWS Launch Plan

This is an engineering/product checklist, not legal, tax, or financial advice. Before depending on ad revenue or Wikimedia-derived content, verify the current provider terms and talk to a CPA or attorney where the risk is material.

## Goal

Launch AIWIKI as a low-cost public demo that can earn ad revenue without creating avoidable AWS bills, ad-policy problems, or Wikimedia licensing risk.

The first public version should optimize for:

- low fixed monthly cost
- conservative ad placement
- clear AI-generated content disclosure
- no secret exposure
- fast cached pages
- source/legal traceability if Wikimedia data is used

## Recommended Production Shape

Start with the cheapest reliable setup:

- AWS Lightsail instance or small EC2 instance for the web app
- Docker Compose running frontend and backend
- `AIWIKI_MODEL_PROVIDER=demo` for the cheapest possible public shell, or `AIWIKI_MODEL_PROVIDER=hosted` for real generated articles
- no public Ollama server
- no database until traffic proves the need
- aggressive in-memory cache for v0, then persistent cache later

Avoid GPU/Ollama on AWS for the first monetized demo. GPU instances will usually cost more than early ad revenue.

## Cost Guardrails

Before creating AWS resources:

1. In AWS Billing, create a zero-ish budget alert:
   - Alert at `$1`
   - Alert at `$5`
   - Alert at `$10`
2. Turn on Free Tier / credit usage alerts.
3. Use one region only.
4. Do not create load balancers, NAT gateways, managed databases, or GPU instances for v0.
5. Stop/delete unused test resources immediately.
6. Keep all model API keys backend-only.

AWS notes:

- Newer AWS Free Tier accounts may receive credits and a limited free-account period, but credits can run out and paid resources can still create charges.
- Lightsail is simpler and more predictable than ECS for the first public demo.
- ECS Fargate is cleaner long-term, but it is not the cheapest first step.

Official starting points:

- AWS Free Tier: https://aws.amazon.com/free/
- Lightsail: https://aws.amazon.com/lightsail/
- Fargate pricing: https://aws.amazon.com/fargate/pricing/

## Ad Provider Plan

Start with an ad abstraction rather than baking one provider into the UI.

Suggested environment variables:

```text
VITE_AD_PROVIDER=none|placeholder|adsense
VITE_ADSENSE_CLIENT=ca-pub-...
VITE_ADSENSE_ARTICLE_TOP_SLOT=...
VITE_ADSENSE_ARTICLE_RAIL_SLOT=...
VITE_ADSENSE_ARTICLE_BOTTOM_SLOT=...
```

Initial modes:

- `none`: no ads, safest for development and review
- `placeholder`: shows reserved ad boxes without loading third-party scripts
- `adsense`: loads AdSense only when approved and configured

Conservative placements:

- one slot below the AI disclaimer
- one desktop right-rail slot below the infobox
- one slot after references / see-also

Avoid:

- ads near search buttons, retry buttons, random article links, or navigation
- sticky/interstitial ads in v0
- layouts where ads can be mistaken for content, citations, downloads, or nav

Before applying to AdSense:

- publish an About page
- publish a Contact page
- publish a Privacy Policy
- publish a Terms/Disclaimer page
- publish an AI-generated content / editorial policy
- keep ad density low
- make sure article pages have useful content beyond thin generated filler

Official policy references:

- AdSense required privacy/cookie content: https://support.google.com/adsense/answer/1348695
- Google Publisher Policies: https://support.google.com/publisherpolicies/answer/10502938
- AdSense Program Policies: https://support.google.com/adsense/answer/48182
- Google guidance on generative AI content: https://developers.google.com/search/docs/fundamentals/using-gen-ai-content

## Getting Paid

For AdSense-style monetization:

1. Create or sign in to Google AdSense.
2. Add and verify the site domain.
3. Complete identity, address, payment profile, and tax steps when requested.
4. Add a bank/payment method once the account reaches the payment method threshold.
5. Payments are monthly if finalized earnings exceed the payment threshold and there are no holds.

Important AdSense payment facts to verify in your own account:

- In USD, the usual payment threshold is `$100`.
- Estimated earnings are finalized early the next month.
- Payments are generally issued between the 21st and 26th if the balance meets threshold and holds are cleared.
- Tax information may be required before payments are released.

Official payment references:

- AdSense payment thresholds: https://support.google.com/adsense/answer/1709871
- AdSense payment timelines: https://support.google.com/adsense/answer/7164703
- Submit US tax info to Google: https://support.google.com/adsense/answer/2490070

## Tax-Legal Operating Notes

This is not tax advice. For a US-based solo project, the simple practical workflow is:

1. Keep a separate spreadsheet or bookkeeping file for:
   - ad revenue
   - AWS charges
   - domain costs
   - model/API costs
   - software/tools used for the site
2. Save monthly invoices:
   - AWS
   - domain registrar
   - hosted model provider
   - ad provider payment statements
3. Complete the ad provider tax interview accurately.
   - US individuals commonly submit W-9 information.
   - Non-US publishers may have withholding/treaty questions.
4. Expect ad income to be taxable.
5. Ask a CPA whether to report as hobby income, sole proprietorship income, or through an LLC once revenue becomes meaningful.

Do not wait until year end to reconstruct costs. Make the bookkeeping boring from day one.

## Wikimedia / Wikipedia Source Use

Safest launch path:

- do not copy Wikipedia article text
- do not scrape Wikipedia HTML pages
- do not imply affiliation with Wikipedia, Wikimedia Foundation, or MediaWiki
- if Wikimedia data is used, use official APIs or dumps with proper attribution and rate limiting

Recommended source modes:

```text
AIWIKI_SOURCE_MODE=none
AIWIKI_SOURCE_MODE=wikimedia_metadata
AIWIKI_SOURCE_MODE=wikimedia_grounded
```

Mode guidance:

- `none`: model-only generation, lowest licensing complexity
- `wikimedia_metadata`: fetch title, description, canonical URL, and lightweight facts; cite consulted source links
- `wikimedia_grounded`: use article text as grounding context; requires careful attribution, revision tracking, license notices, and attorney review before monetization

If using Wikimedia APIs:

- send a descriptive User-Agent with contact info
- respect rate limits and backoff / 429 responses
- cache responsibly
- attribute reused content
- track source URL and revision when article text influences output
- avoid copying or closely paraphrasing large passages

Official Wikimedia references:

- API Usage Guidelines: https://foundation.wikimedia.org/wiki/Policy:Wikimedia_Foundation_API_Usage_Guidelines
- User-Agent Policy: https://foundation.wikimedia.org/wiki/Policy:Wikimedia_Foundation_User-Agent_Policy
- Robot policy: https://wikitech.wikimedia.org/wiki/Robot_policy

## Self-Refinement Loop

AIWIKI should improve generated pages without silently pretending they are verified facts.

Recommended loop for each generated article:

1. Generate article.
2. Run a second model pass as a critic:
   - identify unsupported confident claims
   - flag repeated paragraphs
   - flag vague headings
   - flag medical/legal/financial claims
   - suggest safer wording
3. Apply only safe edits:
   - hedging where needed
   - remove filler
   - improve headings
   - add missing disclaimer/source notes
4. Store quality metadata:
   - generated timestamp
   - model/provider
   - source mode
   - quality score
   - whether human reviewed
5. Never claim the page is independently verified unless a human or trusted source workflow actually verified it.

For v0, keep this as a backend generation pipeline. Later, add a persistent cache and review queue.

Suggested quality gates:

- no raw HTML from the model
- references are AI-generated notes unless real sources are tracked
- no real URLs invented by the model
- YMYL topics receive stronger warnings
- pages with low quality score show fewer/no ads until improved

## AWS Launch Steps

### Phase 1: Local Release Candidate

Run:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\test.ps1
docker compose up --build
```

Verify:

- `http://localhost:8088`
- `http://localhost:8080/api/health`
- search works
- regenerate works
- disclaimers are visible
- no `.env` or secrets are tracked

### Phase 2: Domain and Public Pages

1. Confirm you own the final domain.
2. Add DNS only after the app is deployed.
3. Publish:
   - About
   - Contact
   - Privacy Policy
   - Terms / Disclaimer
   - AI-generated content policy
4. Keep `VITE_AD_PROVIDER=none` until review-ready.

### Phase 3: Cheap AWS Host

Recommended first path:

1. Create a Lightsail instance or small EC2 instance.
2. Install Docker and Docker Compose.
3. Clone the repo.
4. Create production `.env`.
5. Start with:

```bash
docker compose up -d --build
```

6. Put HTTPS in front:
   - Lightsail static IP + DNS + certificate, or
   - Caddy/Nginx reverse proxy with Let's Encrypt, or
   - AWS Load Balancer only when traffic justifies it

Production `.env` starting point:

```text
AIWIKI_MODEL_PROVIDER=demo
AIWIKI_ALLOWED_ORIGINS=https://your-domain.example
AIWIKI_RATE_LIMIT_PER_MINUTE=12
AIWIKI_RATE_LIMIT_BURST=4
AIWIKI_GENERATION_CONCURRENCY=1
AIWIKI_SOURCE_MODE=none
```

For real hosted generation:

```text
AIWIKI_MODEL_PROVIDER=hosted
HOSTED_MODEL_BASE_URL=https://api.openai.com/v1
HOSTED_MODEL_NAME=gpt-4o-mini
HOSTED_MODEL_API_KEY=server-side-secret
HOSTED_MODEL_TIMEOUT_SECONDS=60
```

### Phase 4: Ad Review

1. Apply to AdSense only after the domain is live and stable.
2. Use placeholder ad slots during layout testing.
3. Turn on real ads only after approval:

```text
VITE_AD_PROVIDER=adsense
VITE_ADSENSE_CLIENT=ca-pub-...
```

4. Watch policy center and invalid traffic warnings.
5. Do not buy low-quality traffic.
6. Do not click your own ads.

### Phase 5: Revenue Control

Track:

- AWS daily cost
- model API daily cost
- ad earnings
- page views
- article generation count
- cache hit rate
- revenue per thousand pageviews

Simple profitability rule:

```text
daily ad revenue > daily AWS + model + domain amortized cost
```

If revenue is lower:

- reduce model calls
- increase cache TTL
- use demo mode for anonymous traffic
- require regenerate for expensive model calls
- prewarm only high-value topics
- pause ads or traffic experiments until content quality improves

## Implementation Backlog

Priority order:

1. Add static public pages: About, Contact, Privacy, Terms, AI Content Policy.
2. Add ad slot abstraction with `none|placeholder|adsense`.
3. Add source mode config, starting with `none`.
4. Add Wikimedia metadata-only fetcher with User-Agent, cache, attribution, and rate limiting.
5. Add quality critic/self-refinement pass.
6. Add persistent cache only after public traffic proves the need.
7. Add analytics that do not violate privacy promises.
8. Add real ad provider after policy pages and content quality are ready.

## Stop Conditions

Pause public monetization if:

- AWS/model costs exceed a budget alert
- pages show false medical/legal/financial confidence
- ad provider flags invalid traffic or policy issues
- Wikimedia usage lacks attribution/source tracking
- generated pages become thin, repetitive, or spam-like at scale
