# AIWIKI Next Steps

## Ordered Plan

1. Harden article normalization.
   Ensure `sections`, `references`, `seeAlso`, `paragraphs`, `links`, and infobox rows always become arrays, not `null`. Add tests for missing and partial JSON fields.

2. Fix mobile overflow.
   Make the sticky topbar, search forms, article actions, headings, and footer fit cleanly at 390px. Re-check home, loading, error, and article states.

3. Smoke test Docker once Docker Desktop is available.
   Run `docker compose up --build`, verify `http://localhost:8088`, and test both host Ollama and the optional Ollama Compose override.

4. Improve error states.
   Show clearer messages for backend unreachable, Ollama offline, Ollama timeout, malformed JSON, and CORS/fetch failures. Keep the retry button.

5. Add service-level malformed JSON tests.
   Cover repair success, repair failure, valid JSON with missing arrays, valid JSON with empty references, and invalid field shapes.

6. Reduce factual-trust signals.
   Rename rendered `References` to something like `Generated notes`, and make infobox/reference styling subtly reinforce that content is AI-generated.

7. Improve long-running actions.
   Keep the old article visible while regeneration is pending, and add loading/error feedback for `Random article`.

8. Polish article anchors and footnotes.
   Make section IDs unique and either map references intentionally or use a single generated-note citation model.

9. Harden local HTTP basics.
   Add request body limits, route-level method tests, and clearer status codes for backend errors.

10. Re-run a visual smoke pass.
    Verify desktop and mobile screenshots for home, article, loading, error, dark mode, cached badge, generated links, and regenerate behavior.
