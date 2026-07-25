## Description

<!-- Explain the problem and the solution. -->

## Related issue

<!-- Use "Fixes #123" when applicable. -->

## Compatibility

<!-- Describe public API, behavior, or response-model changes. State "None" if applicable. -->

## Checklist

- [ ] I kept this pull request focused on one logical change.
- [ ] I formatted Go code with `gofmt`.
- [ ] I added or updated deterministic tests.
- [ ] `go test -cover ./...` passes with 100% statement coverage.
- [ ] `go vet ./...` passes.
- [ ] Public network methods accept `context.Context` and use the JSON transport.
- [ ] I updated documentation and `CHANGELOG.md` when user-visible behavior changed.
- [ ] I did not include API keys, shared secrets, session keys, or private response data.
- [ ] I considered whether unlisted Last.fm methods belong under `Client.Experimental`.
