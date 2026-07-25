# go-lastfm v2

A context-aware JSON client for the complete current [Last.fm API](https://www.last.fm/api).

Version 2 is a deliberate breaking release. It covers all 57 currently published
methods across Album, Artist, Auth, Chart, Geo, Library, Tag, Track, and User. Methods
that disappeared from Last.fm's catalog were removed.

## Install

```sh
go get github.com/ndyakov/go-lastfm/v2
```

## Usage

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

client, err := lastfm.New(
	"api-key",
	"shared-secret",
	lastfm.WithUserAgent("my-app/1.0 (contact@example.com)"),
)
if err != nil {
	log.Fatal(err)
}

result, err := client.Artist.GetInfo(ctx, lastfm.ArtistInfoParams{
	ArtistRef: lastfm.ArtistRef{Artist: "Björk", Autocorrect: true},
	Language:  "en",
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(result.Artist.Name)
```

Every network method accepts `context.Context`. Optional API parameters are represented
by request structs, while reusable selectors and pagination use `ArtistRef`, `AlbumRef`,
`TrackRef`, and `Pagination`.

## Authentication

Desktop flow:

1. `token, err := client.Auth.GetToken(ctx)`
2. Open `client.AuthorizationURL(token.Token, "")` and wait for the user to approve.
3. `session, err := client.Auth.GetSession(ctx, token.Token)`

`GetSession` and `GetMobileSession` store the returned session key. It can also be set
with `SetSessionKey`. All write calls are POSTed, signed, and rejected locally when no
session key is configured.

Batch scrobbling accepts 1–50 `Scrobble` values and uses Last.fm's required indexed
parameter names and ASCII-sorted signature calculation.

## Migrating from v1

- Change the import path to `github.com/ndyakov/go-lastfm/v2`.
- `New` now returns `(*Client, error)`; use `MustNew` only for static known-valid options.
- Pass a context as the first argument to every endpoint method.
- Replace positional and `map[string]string` arguments with typed parameter structs.
- Replace `GetSessionKey` with `SessionKey` and `AuthURL` with `AuthorizationURL`.
- Removed undocumented methods: artist/track top fans, user neighbours, tag search, and
  tasteometer comparison.

## API coverage

- Album: 6 methods
- Artist: 10 methods
- Auth: 3 methods
- Chart: 3 methods
- Geo: 2 methods
- Library: 1 method
- Tag: 7 methods
- Track: 12 methods
- User: 13 methods

## Testing

Tests are credential-free and use local HTTP servers. The suite covers all endpoint
wrappers, authentication/signing, GET and POST requests, batch scrobbling, options,
context errors, flexible JSON scalar decoding, and API/HTTP/JSON failure handling.

```sh
go test -cover ./...
```

Additional ideas are recorded in [IMPROVEMENTS.md](IMPROVEMENTS.md).

## Experimental legacy methods

`Client.Experimental` exposes six historically documented methods that are absent from
Last.fm's current method index: tag search, metros, user neighbours, artist and track
top fans, and tasteometer comparison. Each call returns an `ExperimentalResult[T]`
containing typed `Data` and the original `Raw` JSON. These methods are not covered by
the v2 stability guarantee because Last.fm may change or disable them without notice.

Live probes are opt-in:

```sh
LASTFM_RUN_EXPERIMENTAL_TESTS=1 LASTFM_API_KEY=... go test -run TestExperimentalLiveContracts
```

## License

Apache-2.0. See [LICENSE](LICENSE).
