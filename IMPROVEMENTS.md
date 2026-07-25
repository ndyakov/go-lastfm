# Improvement ideas after v2

1. Add an opt-in retry policy for API errors 11/16/29 and HTTP 429/5xx, respecting
   context cancellation and avoiding unsafe automatic scrobble retries.
2. Add credential-gated live contract tests to detect response-shape drift while
   keeping the ordinary suite deterministic and offline.
3. Extend the flexible JSON scalar types if Last.fm introduces additional inconsistent
   representations for dates, booleans, or nullable values.
4. Expose pagination metadata consistently across every paged response without
   breaking the convenient item slices.
5. Add validation helpers for documented constraints such as positive page/limit,
   non-empty required selectors, and the maximum of ten tags per write call.
6. Add observability hooks for request duration, method name, response status, and
   Last.fm error code without exposing credentials or signed parameters.
7. Establish CI gates for 100% statement coverage, race detection, `go vet`,
   staticcheck, govulncheck, and supported Go versions.
