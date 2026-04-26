# AIWIKI License Recommendation

This is not legal advice. Do not add a license until the project owner chooses the intended rights posture.

## Short Recommendation

For a public GitHub repo where others may run, study, and contribute to AIWIKI, use **Apache-2.0**. It is permissive like MIT but includes an explicit patent grant and clearer contribution/notice terms.

If the goal is only public viewing while preserving nearly all rights, use **no open-source license yet** or a custom source-available license reviewed by counsel. A public GitHub repo without a license can be viewed and forked under GitHub's platform terms, but it is not open source and does not clearly allow reuse.

## Options

| Option | What It Allows | Pros | Cons | Fit for AIWIKI |
| --- | --- | --- | --- | --- |
| MIT | Broad use, copy, modify, distribute, commercial use | Very simple, familiar, low friction | No explicit patent grant; minimal attribution/notice language | Good if maximum adoption matters |
| Apache-2.0 | Broad permissive use with patent grant and notices | Better rights clarity for public/commercial projects | Longer and slightly more formal | Best default if AIWIKI should be open source |
| PolyForm Noncommercial | Public source with noncommercial use only | Preserves commercial rights | Not open source; can deter contributors and businesses | Good only if commercial exclusivity matters |
| Custom source-available | Whatever the custom terms say | Can match exact business goals | Needs attorney drafting; unfamiliar to users | Use if "viewable but not reusable" is the real goal |
| All rights reserved / no license | No reuse rights granted beyond platform viewing/forking terms | Maximum default control | Not contributor-friendly; unclear for users who want to run or adapt it | Acceptable temporarily before launch decision |

## Recommended Next Step

Decide whether AIWIKI is meant to be:

- **Open-source demo:** add Apache-2.0.
- **Personal/public portfolio code:** keep no license temporarily and add a README note that no reuse license has been granted yet.
- **Commercial source-available project:** have counsel draft or review source-available terms.

Whatever repo license is chosen, it does not grant users rights to Ollama, Meta Llama models, hosted model APIs, third-party package dependencies, generated outputs, or Wikimedia marks.

