# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- `Track.Date`, so scrobble timestamps from `user.getRecentTracks` are retained.
- `String` and `Streamable` types decoding the shapes Last.fm varies between
  endpoints.

### Fixed

- `user.getInfo` failed to decode because `registered."#text"` arrives as a bare
  JSON number.
- `user.getRecentTracks` failed to decode because `streamable` arrives as a bare
  string rather than the object `track.getInfo` returns.
- Decode errors inside list elements were reported as `cannot unmarshal array
  into T`, hiding the underlying cause.
- `Artist.Name` and `Album.Name` decoded empty when nested in
  `user.getRecentTracks` entries, where the name arrives in `#text`.
- `User.Registered` decoded to zero because `user.getInfo` uses `unixtime` where
  other methods use `uts`.

## [2.0.0] - 2026-07-21

Version 2 is a complete, breaking redesign of go-lastfm for the current Last.fm API.

### Added

- Coverage of all 57 methods in Last.fm's current public API catalog.
- Context support on every network operation.
- Typed request parameters including `ArtistRef`, `AlbumRef`, `TrackRef`, and
  `Pagination`.
- Dedicated Album, Artist, Auth, Chart, Geo, Library, Tag, Track, and User clients.
- Authenticated tag writes, love and unlove, now-playing updates, and scrobbling.
- Batch scrobbling for 1–50 tracks using Last.fm's indexed parameter and signature
  rules.
- Configurable HTTP client, base URL, and User-Agent.
- Structured HTTP and Last.fm API errors with raw response preservation.
- Flexible decoding for Last.fm JSON fields that vary between quoted and unquoted
  numbers or between a single object and an array.
- Thread-safe session-key access.
- An experimental namespace for six historically documented methods that are absent
  from Last.fm's current public method index:
  - `tag.search`
  - `geo.getMetros`
  - `user.getNeighbours`
  - `artist.getTopFans`
  - `track.getTopFans`
  - `tasteometer.compare`
- Credential-gated live contract probes for experimental methods.
- A credential-free test suite with 100% statement coverage.

### Changed

- Changed the module path to `github.com/ndyakov/go-lastfm/v2`.
- Replaced XML requests and responses with JSON.
- Changed `New` to return `(*Client, error)` and accept functional options.
- Changed every endpoint method to take `context.Context` as its first argument.
- Replaced positional and `map[string]string` optional arguments with typed parameter
  structures.
- Replaced the original response structures with JSON-oriented models.
- Renamed `GetSessionKey` to `SessionKey`.
- Renamed `AuthURL` to `AuthorizationURL`.

### Removed

- Removed undocumented methods from the stable endpoint clients.
- Removed the production test fixture transport selected by a magic API key.
- Removed the legacy XML response models.

### Migration example

```go
client, err := lastfm.New(
	"api-key",
	"shared-secret",
	lastfm.WithUserAgent("my-app/1.0"),
)
if err != nil {
	return err
}

result, err := client.Artist.GetInfo(ctx, lastfm.ArtistInfoParams{
	ArtistRef: lastfm.ArtistRef{
		Artist:      "Björk",
		Autocorrect: true,
	},
	Language: "en",
})
```

### Experimental live probes

```sh
LASTFM_RUN_EXPERIMENTAL_TESTS=1 \
LASTFM_API_KEY=... \
go test -run TestExperimentalLiveContracts
```

[Unreleased]: https://github.com/ndyakov/go-lastfm/compare/v2.0.0...HEAD
[2.0.0]: https://github.com/ndyakov/go-lastfm/releases/tag/v2.0.0
