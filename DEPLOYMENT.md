# AIWIKI Deployment

AIWIKI is packaged as two containers:

- `backend`: Go API server, default container port `8080`.
- `frontend`: Nginx static frontend, default container port `80`, proxying `/api/*` to the backend.

Ollama is not bundled into either app container. For local demos, run Ollama separately on a host/service that the backend can reach. For public free-tier demos, use `AIWIKI_MODEL_PROVIDER=demo` or `AIWIKI_MODEL_PROVIDER=hosted` instead of running Ollama on the web host.

## Environment Variables

Backend:

| Name | Required | Default | Notes |
| --- | --- | --- | --- |
| `PORT` | No | `8080` | HTTP listen port used by containers and AWS platforms. |
| `AIWIKI_ADDR` | No | unset | Legacy listen address fallback. `PORT` wins when set. |
| `AIWIKI_MODEL_PROVIDER` | No | `ollama` | `ollama`, `hosted`, or `demo`. Use `demo` for instant free-tier UI demos and `hosted` for OpenAI-compatible APIs. |
| `OLLAMA_BASE_URL` | No | `http://localhost:11434` | URL for Ollama or an Ollama-compatible service. |
| `OLLAMA_MODEL` | No | `llama3.1` | Model name. If unset and Ollama is reachable, backend tries to discover an installed model. |
| `OLLAMA_WORKER_BASE_URLS` | No | empty | Comma-separated Ollama worker URLs for parallel section generation. If empty, section workers use `OLLAMA_BASE_URL`. |
| `AIWIKI_GENERATION_MODE` | No | `pipeline` | `pipeline` creates outline then sections in parallel; `single` uses the old one-shot prompt; `pipeline-strict` disables fallback. |
| `AIWIKI_GENERATION_CONCURRENCY` | No | `2` | Number of sections to generate concurrently. Tune for local GPU/CPU or remote workers. |
| `HOSTED_MODEL_BASE_URL` | No | `https://api.openai.com/v1` | OpenAI-compatible API base URL used when `AIWIKI_MODEL_PROVIDER=hosted`. |
| `HOSTED_MODEL_NAME` | No | `gpt-4o-mini` | Hosted provider model id. Override for OpenRouter or another OpenAI-compatible provider. |
| `HOSTED_MODEL_API_KEY` | Hosted only | empty | Backend-only API key. Never expose this to the frontend. |
| `HOSTED_MODEL_JSON_MODE` | No | `true` | Sends `response_format: {"type":"json_object"}` to compatible chat-completion APIs. Set `false` if your provider rejects it. |
| `HOSTED_MODEL_TIMEOUT_SECONDS` | No | `90` | HTTP timeout for hosted provider calls. |
| `HOSTED_MODEL_HTTP_REFERER` | No | empty | Optional provider header, useful for OpenRouter-style provider attribution. |
| `HOSTED_MODEL_APP_TITLE` | No | `AIWIKI` | Optional provider title header. |
| `AIWIKI_ALLOWED_ORIGINS` | No | local Vite origins | Comma-separated CORS allowlist. Use only when frontend calls backend as a separate origin. |
| `AIWIKI_RATE_LIMIT_PER_MINUTE` | No | `12` | Per-IP limit for article generation requests. |
| `AIWIKI_RATE_LIMIT_BURST` | No | `4` | Burst capacity for article generation requests. |

Frontend:

| Name | Required | Default | Notes |
| --- | --- | --- | --- |
| `PORT` | No | `80` | Nginx listen port inside the frontend container. |
| `BACKEND_UPSTREAM` | No | `http://backend:8080` | Nginx upstream for same-origin `/api` proxying. |
| `AIWIKI_API_BASE_URL` | No | empty | Runtime browser API base URL. Leave empty for same-origin `/api`. |
| `VITE_API_BASE_URL` | No | empty | Build/dev-time fallback. Prefer runtime `AIWIKI_API_BASE_URL` in containers. |

## Local Docker Run

Use this when Ollama is already running on your machine:

```powershell
cd C:\Users\wirsi\OneDrive\Desktop\git\aiwiki
copy .env.example .env
docker compose up --build
```

Open:

```text
http://localhost:8088
```

Health check:

```text
http://localhost:8080/api/health
```

The default `.env.example` points the backend at:

```text
http://host.docker.internal:11434
```

That works for Docker Desktop when Ollama runs on the host.

## Local Docker With Ollama Container

For an all-Compose demo stack, start Ollama first and pull the model:

```powershell
docker compose -f docker-compose.yml -f docker-compose.ollama.yml up -d ollama
docker compose -f docker-compose.yml -f docker-compose.ollama.yml exec ollama ollama pull llama3.1
```

Then start the app:

```powershell
docker compose -f docker-compose.yml -f docker-compose.ollama.yml up --build
```

This is convenient for demos, but model downloads are large and CPU-only generation can be slow.

## Build Images

```powershell
docker build -t aiwiki-backend:local .\backend
docker build -t aiwiki-frontend:local .\frontend
```

## AWS Lightsail Container Option

Lightsail Containers is the simplest public demo path.

Suggested shape:

1. Push `aiwiki-frontend` and `aiwiki-backend` images to a registry.
2. Create a Lightsail container service.
3. Expose only the frontend container publicly.
4. Set frontend `BACKEND_UPSTREAM` to the internal backend service URL if Lightsail service networking supports it for your setup.
5. Leave `AIWIKI_API_BASE_URL` empty when frontend Nginx can proxy `/api`.
6. Set backend `OLLAMA_BASE_URL` to a reachable Ollama host.

If Lightsail cannot cleanly route frontend-to-backend internally for your service shape, use one of these simpler alternatives:

- Deploy a single combined reverse-proxy task elsewhere, such as ECS.
- Temporarily expose the backend and set frontend `AIWIKI_API_BASE_URL` to that HTTPS backend URL.

For a public demo, do not expose Ollama directly to the internet.

## AWS Free-Tier Demo Option

For a truly cheap/free public demo, do not run Ollama on AWS. Use a small free-tier-eligible EC2 instance for the app and set:

```env
AIWIKI_MODEL_PROVIDER=demo
```

This gives an instant public UI demo with the real routes, cache, article layout, disclaimers, and generated JSON shape, but it does not perform live model inference.

For a real public AI demo on free-tier app hosting, use:

```env
AIWIKI_MODEL_PROVIDER=hosted
HOSTED_MODEL_BASE_URL=https://api.openai.com/v1
HOSTED_MODEL_NAME=gpt-4o-mini
HOSTED_MODEL_API_KEY=<backend-only-secret>
```

Or point `HOSTED_MODEL_BASE_URL` and `HOSTED_MODEL_NAME` at an OpenAI-compatible provider. The AWS host can remain small because the model runs elsewhere. See `docs/AWS_FREE_DEMO.md`.

## AWS ECS Fargate Option

ECS Fargate is the cleaner production-ish path.

Suggested shape:

1. Push images to ECR:
   - `aiwiki-backend`
   - `aiwiki-frontend`
2. Create an ECS cluster.
3. Run backend and frontend as services.
4. Put an Application Load Balancer in front of the frontend service.
5. Keep backend private inside the VPC/service network.
6. Configure frontend:
   - `BACKEND_UPSTREAM=http://<backend-service-discovery-name>:8080`
   - `AIWIKI_API_BASE_URL=` empty
7. Configure backend:
   - `PORT=8080`
   - `OLLAMA_BASE_URL=http://<private-ollama-host>:11434`
   - `OLLAMA_MODEL=llama3.1`

This gives the browser one public origin, so same-origin `/api` requests avoid CORS complexity.

## Ollama In AWS

Running Ollama in AWS is possible, but there are practical limits:

- CPU-only instances can be very slow for article generation.
- GPU instances are more expensive and require more setup.
- Model downloads are large and should use persistent storage.
- Fargate is usually not ideal for local LLM inference.
- A private EC2 instance with enough RAM/GPU is the most straightforward Ollama host.

For the first public demo, consider:

- Free/near-free UI demo: `AIWIKI_MODEL_PROVIDER=demo`.
- Real AI demo on small AWS host: `AIWIKI_MODEL_PROVIDER=hosted`.
- Local/private demo: native Ollama on your 3090, outside Docker.
- Paid AWS inference: Ollama on a private GPU EC2 instance.

The backend isolates generation behind the `article.Generator` interface and now includes Ollama, hosted OpenAI-compatible, and template demo providers without changing the frontend contract.

## Parallel Generation

AIWIKI defaults to managed pipeline generation:

1. Generate a broad article outline.
2. Generate a fuller lead summary.
3. Generate each major section independently.
4. Assemble the final article JSON.

Local 3090-style setup:

```text
AIWIKI_GENERATION_MODE=pipeline
AIWIKI_GENERATION_CONCURRENCY=2
OLLAMA_BASE_URL=http://localhost:11434
OLLAMA_WORKER_BASE_URLS=
```

If your local Ollama setup and GPU can handle more parallel requests, try:

```text
AIWIKI_GENERATION_CONCURRENCY=3
```

AWS/EC2 spread setup:

```text
OLLAMA_BASE_URL=http://ollama-outline.internal:11434
OLLAMA_WORKER_BASE_URLS=http://ollama-worker-a.internal:11434,http://ollama-worker-b.internal:11434,http://ollama-worker-c.internal:11434
AIWIKI_GENERATION_CONCURRENCY=3
```

The backend uses `OLLAMA_BASE_URL` for outline/repair work and distributes section calls round-robin across `OLLAMA_WORKER_BASE_URLS`.

## Domain And HTTPS

Recommended:

1. Register or use an existing domain in Route 53.
2. Point the domain to the Lightsail endpoint or Application Load Balancer.
3. Use AWS Certificate Manager for HTTPS when using ALB/CloudFront.
4. Put CloudFront in front later if you want caching and cleaner TLS/domain handling.

Keep `/api` same-origin under the website domain when possible:

```text
https://aiwiki.example.com/api/health
```

That avoids public CORS configuration and is easier to operate.

## Health Checks

Backend:

```text
GET /api/health
```

The response includes backend status, Ollama status, and selected model. The backend container also has a Docker healthcheck that calls this endpoint.

Frontend:

```text
GET /
```

The frontend container has a Docker healthcheck that requests its local Nginx root.

## Secrets

Current Ollama-only deployment does not require API secrets.

When using hosted providers:

- Do not put API keys in frontend env vars.
- Keep provider API keys in backend-only runtime environment variables.
- Use AWS Secrets Manager, SSM Parameter Store, or platform-managed secrets.
- Never commit `.env` files.

## Current Limitations

- In-memory cache is per backend container and disappears on restart.
- Long generation calls may need ALB/Nginx timeouts around 120 seconds.
- No auth, rate limiting, abuse controls, or queueing yet.
- Generated articles are not factual sources and must keep the visible AI disclaimer.
