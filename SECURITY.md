# Security Policy

## Supported versions

| Version | Supported |
| ------- | --------- |
| 2.x     | Yes       |
| 1.x     | No        |

Security fixes are made against the latest v2 release. Users should upgrade to the
latest available patch version before reporting a problem that may already be fixed.

## Reporting a vulnerability

Do not open a public issue for a suspected vulnerability.

Use GitHub's private vulnerability reporting flow from the repository's **Security**
tab. Include:

- The affected package version and Go version.
- A clear description of the vulnerability and its impact.
- Reproduction steps or a minimal proof of concept.
- Any known mitigations or workarounds.
- Whether the report includes exposed Last.fm credentials.

Never include live API secrets, session keys, or other user credentials unless the
maintainer explicitly requests a secure transfer. Revoke exposed Last.fm credentials
immediately.

You should receive an acknowledgement within seven days. The maintainer will assess
severity, coordinate a fix and disclosure timeline, and credit reporters who want to
be credited. Please allow a reasonable remediation period before public disclosure.

## Scope

Relevant reports include credential exposure caused by this library, signature or
authentication flaws, unsafe handling of server responses, and vulnerabilities in the
library's code or release process. Problems in Last.fm's service should be reported to
Last.fm directly.
