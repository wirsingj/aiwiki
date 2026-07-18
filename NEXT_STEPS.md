# AIWIKI Next Steps

## Ordered Plan

1. Re-run a visual mobile smoke pass.
   Verify the 390px fixes for home, loading, error, article, sticky topbar, search forms, article actions, headings, footer, dark mode, cached badge, generated links, and regenerate behavior.

2. Smoke test Docker once Docker Desktop is available.
   Run `docker compose up --build`, verify `http://localhost:8088`, and test both host Ollama and the optional Ollama Compose override.

3. Improve error states.
   Show clearer messages for backend unreachable, Ollama offline, Ollama timeout, malformed JSON, and CORS/fetch failures. Keep the retry button.

4. Expand service-level malformed JSON tests.
   Missing and partial generated arrays are covered. Add remaining coverage for repair success, repair failure, empty references, and invalid field shapes.

5. Reduce factual-trust signals.
   Rename rendered `References` to something like `Generated notes`, and make infobox/reference styling subtly reinforce that content is AI-generated.

6. Improve long-running actions.
   Keep the old article visible while regeneration is pending, and add loading/error feedback for `Random article`.

7. Polish article anchors and footnotes.
   Make section IDs unique and either map references intentionally or use a single generated-note citation model.

8. Harden remaining local HTTP basics.
   POST request body limits are implemented and tested. Add route-level method tests and clearer status codes for backend errors.

## Recently Completed

- 2026-07-16: Hardened article normalization so `sections`, `references`, `seeAlso`, section `paragraphs`, section `links`, and infobox `rows` do not encode as `null`; added missing/partial array tests.
- 2026-07-16: Added CSS guardrails for narrow topbar/search/action/error/footer overflow; frontend build passes, but visual screenshot verification is still pending.
