package lastfm

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

type requesterFunc func(*http.Request) (*http.Response, error)

func (f requesterFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func testClient(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	s := httptest.NewServer(handler)
	c, err := New("key", "secret", WithBaseURL(s.URL+"/2.0/"), WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}
	return c, s
}

func TestOptionsStateAndAuthorization(t *testing.T) {
	if _, err := New("", "", WithHTTPClient(nil)); err == nil {
		t.Fatal("nil client")
	}
	if _, err := New("", "", WithBaseURL(":")); err == nil {
		t.Fatal("bad URL")
	}
	if _, err := New("", "", WithBaseURL("relative")); err == nil {
		t.Fatal("relative URL")
	}
	want := errors.New("option")
	if _, err := New("", "", nil, func(*Client) error { return want }); !errors.Is(err, want) {
		t.Fatal(err)
	}
	c, err := New("key", "secret", WithUserAgent("agent"))
	if err != nil {
		t.Fatal(err)
	}
	c.SetSessionKey("session")
	if c.SessionKey() != "session" || c.UserAgent() != "agent" {
		t.Fatal("state")
	}
	u, err := url.Parse(c.AuthorizationURL("token", "https://app/cb"))
	if err != nil {
		t.Fatal(err)
	}
	if u.Query().Get("api_key") != "key" || u.Query().Get("token") != "token" || u.Query().Get("cb") == "" {
		t.Fatal(u)
	}
	if !strings.Contains(c.AuthorizationURL("", ""), "api_key=key") {
		t.Fatal("minimal auth URL")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("MustNew should panic")
		}
	}()
	_ = MustNew("", "", WithBaseURL(":"))
}

func TestMustURLPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("mustURL should panic")
		}
	}()
	mustURL(":")
}

func TestTransport(t *testing.T) {
	var got *http.Request
	c, s := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		got = r.Clone(r.Context())
		w.Header().Set("X-Test", "ok")
		_, _ = io.WriteString(w, `{}`)
	})
	defer s.Close()
	c.SetSessionKey("sk")
	v := url.Values{"artist": {"a"}}
	if err := c.call(context.Background(), "artist.getInfo", v, new(ArtistInfoResponse)); err != nil {
		t.Fatal(err)
	}
	if got.Method != http.MethodGet ||
		got.Header.Get("User-Agent") != defaultUserAgent ||
		got.Header.Get("Accept") != "application/json" ||
		got.URL.Query().Get("format") != "json" ||
		v.Get("api_key") != "" {
		t.Fatal("GET")
	}
	if err := c.call(context.Background(), "track.love", url.Values{"artist": {"a"}, "track": {"t"}}, new(StatusResponse)); err != nil {
		t.Fatal(err)
	}
	if got.Method != http.MethodPost || got.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Fatal("POST")
	}
	if c.signature(url.Values{"method": {"x"}, "format": {"json"}, "callback": {"cb"}}) != c.signature(url.Values{"method": {"x"}}) {
		t.Fatal("signature exclusions")
	}
	signingClient := MustNew("api_key_for_testing", "api_secret_for_testing")
	if got := signingClient.signature(url.Values{"api_key": {"api_key_for_testing"}, "method": {"auth.getToken"}}); got != "67243a89bada3fd15e6d152c2ad5cc99" {
		t.Fatalf("signature = %s", got)
	}
	if len(cloneValues(url.Values{"x": {"1", "2"}})["x"]) != 2 {
		t.Fatal("clone")
	}
	if err := c.call(nil, "artist.getInfo", nil, new(ArtistInfoResponse)); err == nil {
		t.Fatal("nil context")
	}
	if err := c.call(context.Background(), "missing", nil, new(StatusResponse)); err == nil {
		t.Fatal("unknown method")
	}
	c.SetSessionKey("")
	if err := c.call(context.Background(), "track.love", nil, new(StatusResponse)); err == nil {
		t.Fatal("session")
	}
}

func TestResponseShapes(t *testing.T) {
	var artists ChartArtistsResponse
	if err := json.Unmarshal([]byte(`{"artists":{"artist":[{"name":"A","playcount":"12"}]}}`), &artists); err != nil {
		t.Fatal(err)
	}
	if len(artists.Artists.Artists) != 1 || artists.Artists.Artists[0].Name != "A" || artists.Artists.Artists[0].Playcount != 12 {
		t.Fatalf("artists: %#v", artists.Artists)
	}
	var tags PersonalTagsResponse
	if err := json.Unmarshal([]byte(`{"taggings":{"artists":{"artist":{"name":"A"}}}}`), &tags); err != nil {
		t.Fatal(err)
	}
	if len(tags.Taggings.Artists.Artists) != 1 || tags.Taggings.Artists.Artists[0].Name != "A" {
		t.Fatalf("tags: %#v", tags.Taggings.Artists.Artists)
	}
	var charts WeeklyChartListResponse
	if err := json.Unmarshal([]byte(`{"weeklychartlist":{"chart":[{"from":"1","to":"2"}]}}`), &charts); err != nil {
		t.Fatal(err)
	}
	if len(charts.Charts.Charts) != 1 || charts.Charts.Charts[0].From != 1 || charts.Charts.Charts[0].To != 2 {
		t.Fatalf("charts: %#v", charts.Charts)
	}
}

func TestFlexibleJSONTypes(t *testing.T) {
	for _, input := range []string{`null`, `""`, `12`, `"13"`} {
		var value Integer
		if err := json.Unmarshal([]byte(input), &value); err != nil {
			t.Fatalf("Integer %s: %v", input, err)
		}
	}
	var integer Integer
	if err := json.Unmarshal([]byte(`"nope"`), &integer); err == nil {
		t.Fatal("expected integer error")
	}

	for _, input := range []string{`null`, `""`, `1.5`, `"2.5"`} {
		var value Float
		if err := json.Unmarshal([]byte(input), &value); err != nil {
			t.Fatalf("Float %s: %v", input, err)
		}
	}
	var floating Float
	if err := json.Unmarshal([]byte(`"nope"`), &floating); err == nil {
		t.Fatal("expected float error")
	}

	var list List[Tag]
	if err := json.Unmarshal([]byte(`null`), &list); err != nil || list != nil {
		t.Fatalf("null list: %#v, %v", list, err)
	}
	if err := list.UnmarshalJSON([]byte(`{bad`)); err == nil {
		t.Fatal("expected list error")
	}
}

func TestExperimentalMethods(t *testing.T) {
	payload := `{
		"results":{"tagmatches":{"tag":{"name":"rock"}}},
		"metros":{"metro":{"name":"Sofia","country":"Bulgaria"}},
		"neighbours":{"user":{"name":"listener"}},
		"topfans":{"user":{"name":"fan","weight":"1"}},
		"comparison":{"result":{"score":"0.5","artists":{"artist":{"name":"A"}}}}
	}`
	client, server := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, payload)
	})
	defer server.Close()
	ctx := context.Background()

	tagResult, err := client.Experimental.SearchTags(ctx, TagSearchParams{
		Tag:        "rock",
		Pagination: Pagination{Page: 1, Limit: 10},
	})
	if err != nil || len(tagResult.Raw) == 0 || len(tagResult.Data.Results.Matches.Tags) != 1 {
		t.Fatalf("tag search: %#v, %v", tagResult, err)
	}

	metros, err := client.Experimental.GetMetros(ctx, "Bulgaria")
	if err != nil || len(metros.Data.Metros.Metros) != 1 {
		t.Fatalf("metros: %#v, %v", metros, err)
	}

	neighbours, err := client.Experimental.GetNeighbours(ctx, NeighboursParams{User: "listener", Limit: 10})
	if err != nil || len(neighbours.Data.Neighbours.Users) != 1 {
		t.Fatalf("neighbours: %#v, %v", neighbours, err)
	}

	artistFans, err := client.Experimental.GetArtistTopFans(ctx, TopFansParams{Artist: "A"})
	if err != nil || len(artistFans.Data.TopFans.Fans) != 1 {
		t.Fatalf("artist fans: %#v, %v", artistFans, err)
	}

	trackFans, err := client.Experimental.GetTrackTopFans(ctx, TopFansParams{Artist: "A", Track: "T", MBID: "m"})
	if err != nil || len(trackFans.Data.TopFans.Fans) != 1 {
		t.Fatalf("track fans: %#v, %v", trackFans, err)
	}

	taste, err := client.Experimental.CompareTaste(ctx, TasteometerParams{
		Type1: "user", Value1: "a", Type2: "user", Value2: "b", Limit: 10,
	})
	if err != nil || taste.Data.Comparison.Result.Score != 0.5 {
		t.Fatalf("tasteometer: %#v, %v", taste, err)
	}

	var malformed ExperimentalResult[TagSearchResponse]
	if err := malformed.UnmarshalJSON([]byte(`{bad`)); err == nil {
		t.Fatal("expected malformed experimental response error")
	}
}

func TestTransportErrors(t *testing.T) {
	for _, tc := range []struct {
		name, body string
		status     int
		want       string
	}{{"api", `{"error":29,"message":"rate"}`, 200, "API error 29"}, {"json", `{bad`, 200, "decode"}, {"http", "down", 503, "HTTP 503"}, {"http-empty", "", 500, "HTTP 500"}} {
		t.Run(tc.name, func(t *testing.T) {
			c, s := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tc.status)
				_, _ = io.WriteString(w, tc.body)
			})
			defer s.Close()
			err := c.call(context.Background(), "artist.getInfo", nil, new(ArtistInfoResponse))
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("%v", err)
			}
		})
	}
	client, server := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, `{"error":29,"message":"rate"}`)
	})
	defer server.Close()
	err := client.call(context.Background(), "artist.getInfo", nil, new(ArtistInfoResponse))
	var apiError *APIError
	if !errors.As(err, &apiError) || len(apiError.Raw) == 0 {
		t.Fatalf("API error raw response: %v", err)
	}
	c := MustNew("", "")
	c.httpClient = requesterFunc(func(*http.Request) (*http.Response, error) { return nil, errors.New("offline") })
	if err := c.call(context.Background(), "artist.getInfo", nil, new(ArtistInfoResponse)); err == nil {
		t.Fatal("transport")
	}
	c.httpClient = requesterFunc(func(*http.Request) (*http.Response, error) { return nil, nil })
	if err := c.call(context.Background(), "artist.getInfo", nil, new(ArtistInfoResponse)); err == nil {
		t.Fatal("nil response")
	}
	c.httpClient = requesterFunc(func(*http.Request) (*http.Response, error) { return &http.Response{StatusCode: 200}, nil })
	if err := c.call(context.Background(), "artist.getInfo", nil, new(ArtistInfoResponse)); err == nil {
		t.Fatal("nil body")
	}
	c.httpClient = requesterFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(failingReader{}),
		}, nil
	})
	if err := c.call(context.Background(), "artist.getInfo", nil, new(ArtistInfoResponse)); err == nil {
		t.Fatal("read body")
	}
	c.httpClient = requesterFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	var unsupported chan int
	if err := c.call(context.Background(), "artist.getInfo", nil, &unsupported); err == nil {
		t.Fatal("decode target")
	}
	c.baseURL = &url.URL{Scheme: ":"}
	if err := c.call(context.Background(), "artist.getInfo", nil, new(ArtistInfoResponse)); err == nil {
		t.Fatal("request creation")
	}
	if !strings.Contains((&HTTPError{StatusCode: 500, Status: "bad"}).Error(), "500") {
		t.Fatal("HTTP error")
	}
	if !strings.Contains((&APIError{Code: 7}).Error(), "7") {
		t.Fatal("API error")
	}
}

func TestAllMethods(t *testing.T) {
	c, s := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, `{"token":"token","session":{"key":"session"}}`)
	})
	defer s.Close()
	c.SetSessionKey("session")
	ctx := context.Background()
	yes := true
	no := false
	calls := []func() error{
		func() error {
			_, e := c.Album.GetInfo(ctx, AlbumInfoParams{AlbumRef: AlbumRef{Artist: "a", Album: "b", Autocorrect: true}, Username: "u"})
			return e
		}, func() error {
			_, e := c.Album.GetTags(ctx, AlbumTagsParams{AlbumRef: AlbumRef{MBID: "m"}, User: "u"})
			return e
		}, func() error {
			_, e := c.Album.GetTopTags(ctx, AlbumRef{Artist: "a", Album: "b"})
			return e
		}, func() error {
			_, e := c.Album.Search(ctx, AlbumSearchParams{Album: "b", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			return c.Album.AddTags(ctx, "a", "b", []string{"x", "y"})
		}, func() error {
			return c.Album.RemoveTag(ctx, "a", "b", "x")
		},
		func() error {
			_, e := c.Artist.GetCorrection(ctx, "a")
			return e
		}, func() error {
			_, e := c.Artist.GetInfo(ctx, ArtistInfoParams{ArtistRef: ArtistRef{Artist: "a", Autocorrect: true}, Username: "u", Language: "en"})
			return e
		}, func() error {
			_, e := c.Artist.GetSimilar(ctx, ArtistSimilarParams{ArtistRef: ArtistRef{MBID: "m"}, Limit: 10})
			return e
		}, func() error {
			_, e := c.Artist.GetTags(ctx, ArtistTagsParams{ArtistRef: ArtistRef{Artist: "a"}, User: "u"})
			return e
		}, func() error {
			_, e := c.Artist.GetTopAlbums(ctx, ArtistPagedParams{ArtistRef: ArtistRef{Artist: "a"}, Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.Artist.GetTopTags(ctx, ArtistRef{Artist: "a"})
			return e
		}, func() error {
			_, e := c.Artist.GetTopTracks(ctx, ArtistPagedParams{ArtistRef: ArtistRef{Artist: "a"}, Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.Artist.Search(ctx, ArtistSearchParams{Artist: "a", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			return c.Artist.AddTags(ctx, "a", []string{"x", "y"})
		}, func() error {
			return c.Artist.RemoveTag(ctx, "a", "x")
		},
		func() error {
			_, e := c.Auth.GetToken(ctx)
			return e
		}, func() error {
			_, e := c.Auth.GetSession(ctx, "token")
			return e
		}, func() error {
			_, e := c.Auth.GetMobileSession(ctx, "u", "p")
			return e
		},
		func() error {
			_, e := c.Chart.GetTopArtists(ctx, Pagination{Page: 2, Limit: 10})
			return e
		}, func() error {
			_, e := c.Chart.GetTopTags(ctx, Pagination{Page: 2, Limit: 10})
			return e
		}, func() error {
			_, e := c.Chart.GetTopTracks(ctx, Pagination{Page: 2, Limit: 10})
			return e
		},
		func() error {
			_, e := c.Geo.GetTopArtists(ctx, GeoParams{Country: "Bulgaria", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.Geo.GetTopTracks(ctx, GeoParams{Country: "Bulgaria", Location: "Sofia", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.Library.GetArtists(ctx, LibraryArtistsParams{User: "u", Pagination: Pagination{2, 10}})
			return e
		},
		func() error {
			_, e := c.Tag.GetInfo(ctx, TagInfoParams{Tag: "rock", Language: "en"})
			return e
		}, func() error {
			_, e := c.Tag.GetSimilar(ctx, "rock")
			return e
		}, func() error {
			_, e := c.Tag.GetTopAlbums(ctx, TagPagedParams{Tag: "rock", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.Tag.GetTopArtists(ctx, TagPagedParams{Tag: "rock", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.Tag.GetTopTags(ctx)
			return e
		}, func() error {
			_, e := c.Tag.GetTopTracks(ctx, TagPagedParams{Tag: "rock", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.Tag.GetWeeklyChartList(ctx, "rock")
			return e
		},
		func() error {
			_, e := c.Track.GetCorrection(ctx, "a", "t")
			return e
		}, func() error {
			_, e := c.Track.GetInfo(ctx, TrackInfoParams{TrackRef: TrackRef{Artist: "a", Track: "t", Autocorrect: true}, Username: "u"})
			return e
		}, func() error {
			_, e := c.Track.GetSimilar(ctx, TrackSimilarParams{TrackRef: TrackRef{MBID: "m"}, Limit: 10})
			return e
		}, func() error {
			_, e := c.Track.GetTags(ctx, TrackTagsParams{TrackRef: TrackRef{Artist: "a", Track: "t"}, User: "u"})
			return e
		}, func() error {
			_, e := c.Track.GetTopTags(ctx, TrackRef{Artist: "a", Track: "t"})
			return e
		}, func() error {
			_, e := c.Track.Search(ctx, TrackSearchParams{Track: "t", Artist: "a", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			return c.Track.AddTags(ctx, "a", "t", []string{"x", "y"})
		}, func() error {
			return c.Track.RemoveTag(ctx, "a", "t", "x")
		}, func() error {
			return c.Track.Love(ctx, "a", "t")
		}, func() error {
			return c.Track.Unlove(ctx, "a", "t")
		}, func() error {
			_, e := c.Track.UpdateNowPlaying(ctx, NowPlayingParams{Artist: "a", Track: "t", Album: "b", AlbumArtist: "aa", MBID: "m", TrackNumber: 1, Duration: 100})
			return e
		}, func() error {
			_, e := c.Track.Scrobble(ctx, []Scrobble{{Artist: "a", Track: "t", Timestamp: time.Unix(1, 0), ChosenByUser: &yes, TrackNumber: 1, Duration: 100}, {Artist: "b", Track: "u", Timestamp: time.Unix(2, 0), ChosenByUser: &no}})
			return e
		},
		func() error {
			_, e := c.User.GetFriends(ctx, UserPagedParams{User: "u", IncludeRecentTracks: true, Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.User.GetInfo(ctx, "u")
			return e
		}, func() error {
			_, e := c.User.GetLovedTracks(ctx, UserPagedParams{User: "u", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.User.GetPersonalTags(ctx, PersonalTagsParams{User: "u", Tag: "rock", TaggingType: "artist", Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.User.GetRecentTracks(ctx, RecentTracksParams{User: "u", Pagination: Pagination{2, 10}, From: 1, To: 2, Extended: true})
			return e
		}, func() error {
			_, e := c.User.GetTopAlbums(ctx, UserTopParams{User: "u", Period: Period7Day, Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.User.GetTopArtists(ctx, UserTopParams{User: "u", Period: Period7Day, Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.User.GetTopTags(ctx, UserTopTagsParams{User: "u", Limit: 10})
			return e
		}, func() error {
			_, e := c.User.GetTopTracks(ctx, UserTopParams{User: "u", Period: Period7Day, Pagination: Pagination{2, 10}})
			return e
		}, func() error {
			_, e := c.User.GetWeeklyAlbumChart(ctx, WeeklyChartParams{User: "u", From: 1, To: 2})
			return e
		}, func() error {
			_, e := c.User.GetWeeklyArtistChart(ctx, WeeklyChartParams{User: "u", From: 1, To: 2})
			return e
		}, func() error {
			_, e := c.User.GetWeeklyChartList(ctx, "u")
			return e
		}, func() error {
			_, e := c.User.GetWeeklyTrackChart(ctx, WeeklyChartParams{User: "u", From: 1, To: 2})
			return e
		},
	}
	for i, call := range calls {
		if err := call(); err != nil {
			t.Fatalf("method %d: %v", i, err)
		}
	}
	if _, err := c.Track.Scrobble(ctx, nil); err == nil {
		t.Fatal("empty batch")
	}
	if _, err := c.Track.Scrobble(ctx, make([]Scrobble, 51)); err == nil {
		t.Fatal("large batch")
	}
	_ = PeriodOverall
	_ = Period1Month
	_ = Period3Month
	_ = Period6Month
	_ = Period12Month
}
