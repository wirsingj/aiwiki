# AIWIKI Third-Party Notices

Date: 2026-04-26

This report was generated from `backend/go.mod`, `frontend/package.json`, `frontend/package-lock.json`, installed package metadata, and spot checks against npm registry license metadata for packages whose installed metadata was incomplete.

It is a practical release notice, not legal advice.

## Backend

The Go backend currently uses only the Go standard library. No third-party Go modules are listed in `backend/go.mod`.

## Frontend Direct Dependencies

| Package | Version | License |
| --- | ---: | --- |
| `react` | 19.2.5 | MIT |
| `react-dom` | 19.2.5 | MIT |
| `react-router-dom` | 7.14.2 | MIT |
| `lucide-react` | 1.11.0 | ISC |

## Frontend Development Dependencies

| Package | Version | License |
| --- | ---: | --- |
| `@types/react` | 19.2.14 | MIT |
| `@types/react-dom` | 19.2.3 | MIT |
| `@vitejs/plugin-react` | 6.0.1 | MIT |
| `typescript` | 6.0.3 | Apache-2.0 |
| `vite` | 8.0.10 | MIT |

## License Summary for Installed Frontend Dependency Tree

| License | Packages |
| --- | --- |
| MIT | `@emnapi/core`, `@emnapi/runtime`, `@emnapi/wasi-threads`, `@napi-rs/wasm-runtime`, `@oxc-project/types`, `@rolldown/binding-*`, `@rolldown/pluginutils`, `@types/react`, `@types/react-dom`, `@tybys/wasm-util`, `@vitejs/plugin-react`, `cookie`, `csstype`, `fdir`, `fsevents`, `nanoid`, `picomatch`, `postcss`, `react`, `react-dom`, `react-router`, `react-router-dom`, `rolldown`, `scheduler`, `set-cookie-parser`, `tinyglobby`, `vite` |
| Apache-2.0 | `detect-libc`, `typescript` |
| ISC | `lucide-react`, `picocolors` |
| BSD-3-Clause | `source-map-js` |
| MPL-2.0 | `lightningcss`, `lightningcss-*` platform packages |
| 0BSD | `tslib` |

## Copyleft / Compatibility Notes

- No GPL or AGPL dependencies were found.
- `lightningcss` is MPL-2.0. MPL-2.0 is file-level copyleft and generally compatible with use in proprietary or permissive applications when notices and license obligations are respected. Keep this notice if distributing builds or source.
- Optional native packages may vary by platform. Re-run the license check after dependency updates or when building on a different OS/CPU target.

## Model and Runtime Notices

AIWIKI does not grant rights to any model runtime or model weights.

- Ollama source code is MIT-licensed, but Ollama services/website/API terms may also apply depending on use.
- Llama 3.1 is governed by the Meta Llama 3.1 Community License and Acceptable Use Policy, not by this repo's future project license.
- If a deployment distributes or makes available Llama Materials, or a service/product containing them, verify attribution and notice obligations before launch.

## Release Recommendation

Before a tagged public release:

1. Run `npm install` or `npm ci` from a clean checkout.
2. Re-run the dependency/license report.
3. Confirm no GPL/AGPL or unlicensed packages were introduced.
4. Keep this notice updated with direct dependency versions.

