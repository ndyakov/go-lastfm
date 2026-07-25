# Contributing to go-lastfm

Thank you for helping improve go-lastfm.

## Before opening an issue

- Search existing issues to avoid duplicates.
- Confirm the behavior against the current
  [Last.fm API documentation](https://www.last.fm/api).
- Use the bug or feature issue form and include the affected Last.fm method.
- Report security vulnerabilities privately as described in [SECURITY.md](SECURITY.md).

Questions about Last.fm accounts, API keys, rate limits, or service availability are
usually better directed to Last.fm support. This repository covers the Go client.

## Development setup

The module requires Go 1.22 or newer.

```sh
git clone https://github.com/ndyakov/go-lastfm.git
cd go-lastfm
go test ./...
```

The ordinary tests are offline and do not require Last.fm credentials. Experimental
live contract tests are opt-in and must never contain committed credentials.

## Making changes

1. Fork the repository and create a focused branch.
2. Keep public methods context-aware and preserve the JSON transport.
3. Add or update typed request and response models when changing an endpoint.
4. Add deterministic tests using a local HTTP server.
5. Keep statement coverage at 100%.
6. Format and verify the change:

```sh
gofmt -w *.go
go test -cover ./...
go vet ./...
```

Do not commit API keys, shared secrets, session keys, captured private responses, or
generated coverage files.

## API compatibility

Stable clients cover methods in Last.fm's current public catalog. Historically
documented or unlisted methods belong under `Client.Experimental` and must return both
typed data and raw JSON. Changes to the stable public API require a clear compatibility
assessment and, when breaking, a major-version plan.

## Pull requests

- Keep each pull request limited to one logical change.
- Explain the problem, the approach, and any compatibility impact.
- Link related issues.
- Update `README.md` and `CHANGELOG.md` when behavior visible to users changes.
- Ensure all pull-request checklist items are complete.
- Be responsive to review feedback.

By contributing, you agree that your contribution is licensed under the repository's
Apache-2.0 license and that you will follow the [Code of Conduct](CODE_OF_CONDUCT.md).
