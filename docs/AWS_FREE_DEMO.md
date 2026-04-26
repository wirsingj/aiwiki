# AIWIKI Free-Tier AWS Demo Plan

This plan is optimized for "show the website publicly without paying for local model inference."

## Reality Check

AWS free-tier compute can host the AIWIKI web app, but it is not enough for comfortable LLM inference. Running Ollama on a free-tier EC2 instance will be slow, memory-constrained, and unreliable for public traffic.

Use one of these modes:

| Mode | Cost shape | Good for | Tradeoff |
| --- | --- | --- | --- |
| `AIWIKI_MODEL_PROVIDER=demo` | AWS hosting only | Free/near-free public UI demo | Not real AI generation |
| `AIWIKI_MODEL_PROVIDER=hosted` | AWS hosting plus model API cost/credits | Real public AI demo | Requires backend API key and provider terms |
| `AIWIKI_MODEL_PROVIDER=ollama` | Local GPU or paid EC2/GPU | Local realism, private demos | Not suitable for AWS free-tier public site |

## Recommended Free AWS Shape

Use one small EC2 instance that is marked Free Tier eligible in your account and region.

1. Run backend and frontend containers on the instance.
2. Do not run Ollama on the instance.
3. Start with `AIWIKI_MODEL_PROVIDER=demo`.
4. Switch to `hosted` when you want real generation.
5. Put a domain/HTTPS proxy in front later only after the demo works.

AWS changed Free Tier benefits for accounts created on or after July 15, 2025. Verify your account's exact eligible instance types and credits in the AWS console before leaving anything running.

## Environment: Instant Demo Mode

```env
AIWIKI_MODEL_PROVIDER=demo
AIWIKI_ALLOWED_ORIGINS=https://your-domain.example
AIWIKI_RATE_LIMIT_PER_MINUTE=12
AIWIKI_RATE_LIMIT_BURST=4
```

Demo mode returns structured AIWIKI article JSON immediately using a deterministic template. It exercises the real frontend, routes, cache, safety banner, infobox, links, and article rendering without model cost.

## Environment: Hosted Model Mode

```env
AIWIKI_MODEL_PROVIDER=hosted
HOSTED_MODEL_BASE_URL=https://api.openai.com/v1
HOSTED_MODEL_NAME=gpt-4o-mini
HOSTED_MODEL_API_KEY=replace-with-backend-only-secret
HOSTED_MODEL_JSON_MODE=true
HOSTED_MODEL_TIMEOUT_SECONDS=90
```

For OpenRouter-compatible use:

```env
AIWIKI_MODEL_PROVIDER=hosted
HOSTED_MODEL_BASE_URL=https://openrouter.ai/api/v1
HOSTED_MODEL_NAME=replace-with-provider-model-id
HOSTED_MODEL_API_KEY=replace-with-backend-only-secret
HOSTED_MODEL_HTTP_REFERER=https://your-domain.example
HOSTED_MODEL_APP_TITLE=AIWIKI
```

Never expose hosted model API keys to the frontend.

## Local Fast Path

For your machine, prefer native dev mode or backend container pointed at a GPU-enabled host Ollama. The all-Compose Ollama container is convenient but currently CPU-only in WSL on this machine.

Fastest local workflow:

```powershell
ollama serve
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\dev.ps1
```

If you want Docker frontend/backend with host Ollama later, configure Ollama to listen on a network address that containers can reach and set:

```env
AIWIKI_MODEL_PROVIDER=ollama
OLLAMA_BASE_URL=http://<host-reachable-ip>:11434
OLLAMA_MODEL=llama3.2:3b
```

## Do Not Do This on Free AWS

- Do not run the Ollama Compose override for the public free-tier demo.
- Do not expose Ollama directly to the internet.
- Do not put API keys in `VITE_*`, frontend env vars, Dockerfiles, or committed `.env` files.
- Do not assume Fargate is free; price it before using it.

