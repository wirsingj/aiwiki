# AIWIKI Public Demo Runbook

Use this as the day-to-day launch checklist once AIWIKI moves from local demo to public AWS demo.

## Daily Before Launch

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\test.ps1
git status --short
```

Confirm:

- no `.env` files are tracked
- no API keys appear in docs, Dockerfiles, or compose files
- generated pages show AI and source disclaimers
- `/api/health` reports the expected provider

## AWS Console Checklist

1. Billing:
   - create budget alerts
   - confirm Free Tier / credit balance
   - check current month cost
2. Compute:
   - keep only one public demo instance/service running
   - delete failed experiments
   - avoid GPU, NAT Gateway, managed DB, and load balancer until needed
3. Networking:
   - expose only HTTP/HTTPS
   - do not expose Ollama
   - restrict SSH to your IP if possible
4. Secrets:
   - put hosted model keys only in server environment
   - never bake secrets into Docker images
5. Domain:
   - point DNS only after health checks pass
   - enable HTTPS before ad review

## Ad Provider Checklist

1. Confirm public pages:
   - About
   - Contact
   - Privacy Policy
   - Terms / Disclaimer
   - AI Content Policy
2. Confirm no deceptive ad placement.
3. Confirm content has real user value.
4. Apply to AdSense or another provider.
5. Add tax/payment info inside provider dashboard.
6. Enable `placeholder` ads first, then real ads after approval.

## Payment / Tax Checklist

Monthly:

- export AWS invoice
- export model/API invoice
- export domain/hosting invoices
- export ad provider earnings report
- update bookkeeping spreadsheet

At payment setup:

- complete ad provider tax interview
- add bank/payment method when eligible
- verify identity/address if requested
- save any 1099 or tax forms

Ask a CPA before treating this as a business deduction strategy.

## Incident Checklist

If AWS cost spikes:

1. Stop the instance/service.
2. Check Billing and Cost Explorer.
3. Disable hosted model calls.
4. Increase cache TTL and lower generation limits.
5. Review logs for abuse.

If ad provider flags the site:

1. Disable ads.
2. Screenshot the warning.
3. Review affected pages.
4. Remove risky placements or content.
5. Request review only after fixes.

If Wikimedia usage is questioned:

1. Disable source mode.
2. Preserve logs/source metadata.
3. Verify User-Agent, rate limits, attribution, and license notices.
4. Get legal review before re-enabling grounded mode.
