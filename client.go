package lastfm

import (
	"context"
	"crypto/md5" // #nosec G501 -- required by the Last.fm protocol.
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	defaultAPIURL    = "https://ws.audioscrobbler.com/2.0/"
	defaultAuthURL   = "https://www.last.fm/api/auth/"
	defaultUserAgent = "go-lastfm/v2"
)

type requester interface {
	Do(*http.Request) (*http.Response, error)
}

// Option configures Client.
type Option func(*Client) error

// WithHTTPClient replaces the default timeout-enabled HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) error {
		if client == nil {
			return errors.New("lastfm: nil HTTP client")
		}
		c.httpClient = client
		return nil
	}
}

// WithBaseURL replaces the API URL. It is useful for proxies and tests.
func WithBaseURL(rawURL string) Option {
	return func(c *Client) error {
		u, err := url.Parse(rawURL)
		if err != nil {
			return fmt.Errorf("lastfm: invalid base URL: %w", err)
		}
		if !u.IsAbs() || u.Host == "" {
			return errors.New("lastfm: base URL must be absolute")
		}
		c.baseURL = u
		return nil
	}
}

// WithUserAgent sets the User-Agent header. An empty value suppresses the header.
func WithUserAgent(userAgent string) Option {
	return func(c *Client) error { c.userAgent = userAgent; return nil }
}

// Client is a concurrency-safe Last.fm API client. Configure credentials before
// sharing it when SetSessionKey will be used.
type Client struct {
	apiKey, apiSecret, sessionKey, userAgent string
	sessionMu                                sync.RWMutex
	httpClient                               requester
	baseURL                                  *url.URL
	Album                                    AlbumClient
	Artist                                   ArtistClient
	Auth                                     AuthClient
	Chart                                    ChartClient
	Experimental                             ExperimentalClient
	Geo                                      GeoClient
	Library                                  LibraryClient
	Tag                                      TagClient
	Track                                    TrackClient
	User                                     UserClient
}

// New constructs a v2 client.
func New(apiKey, apiSecret string, options ...Option) (*Client, error) {
	c := &Client{apiKey: apiKey, apiSecret: apiSecret, userAgent: defaultUserAgent,
		httpClient: &http.Client{Timeout: 30 * time.Second}, baseURL: mustURL(defaultAPIURL)}
	for _, option := range options {
		if option == nil {
			continue
		}
		if err := option(c); err != nil {
			return nil, err
		}
	}
	c.Album = AlbumClient{c}
	c.Artist = ArtistClient{c}
	c.Auth = AuthClient{c}
	c.Chart = ChartClient{c}
	c.Experimental = ExperimentalClient{c}
	c.Geo = GeoClient{c}
	c.Library = LibraryClient{c}
	c.Tag = TagClient{c}
	c.Track = TrackClient{c}
	c.User = UserClient{c}
	return c, nil
}

// MustNew is New for callers whose configuration is static and known-valid.
func MustNew(apiKey, apiSecret string, options ...Option) *Client {
	c, err := New(apiKey, apiSecret, options...)
	if err != nil {
		panic(err)
	}
	return c
}

func (c *Client) SetSessionKey(key string) {
	c.sessionMu.Lock()
	c.sessionKey = key
	c.sessionMu.Unlock()
}

func (c *Client) SessionKey() string {
	c.sessionMu.RLock()
	defer c.sessionMu.RUnlock()
	return c.sessionKey
}

func (c *Client) UserAgent() string {
	return c.userAgent
}

// AuthorizationURL returns the browser URL for desktop or web authorization.
func (c *Client) AuthorizationURL(token, callback string) string {
	u := mustURL(defaultAuthURL)
	q := u.Query()
	q.Set("api_key", c.apiKey)
	if token != "" {
		q.Set("token", token)
	}
	if callback != "" {
		q.Set("cb", callback)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Client) call(ctx context.Context, method string, params url.Values, out any) error {
	if ctx == nil {
		return errors.New("lastfm: nil context")
	}
	values := cloneValues(params)
	values.Set("method", method)
	values.Set("api_key", c.apiKey)
	values.Set("format", "json")
	rule, ok := methodRules[method]
	if !ok {
		return fmt.Errorf("lastfm: unknown method %q", method)
	}
	if rule.session {
		sessionKey := c.SessionKey()
		if sessionKey == "" {
			return errors.New("lastfm: session key is required")
		}
		values.Set("sk", sessionKey)
	}
	if rule.signed {
		values.Set("api_sig", c.signature(values))
	}

	var req *http.Request
	var err error
	if rule.post {
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL.String(), strings.NewReader(values.Encode()))
		if err == nil {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	} else {
		u := *c.baseURL
		u.RawQuery = values.Encode()
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	}
	if err != nil {
		return fmt.Errorf("lastfm: create request: %w", err)
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("lastfm: request failed: %w", err)
	}
	if resp == nil || resp.Body == nil {
		return errors.New("lastfm: HTTP client returned an empty response")
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return &HTTPError{StatusCode: resp.StatusCode, Status: resp.Status, Body: strings.TrimSpace(string(body))}
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("lastfm: read response: %w", err)
	}
	var apiError APIError
	if err := json.Unmarshal(data, &apiError); err != nil {
		return fmt.Errorf("lastfm: decode response: %w", err)
	}
	apiError.Raw = append(apiError.Raw[:0], data...)
	if apiError.Code != 0 {
		return &apiError
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("lastfm: decode response: %w", err)
	}
	return nil
}

func (c *Client) signature(values url.Values) string {
	keys := make([]string, 0, len(values))
	for key := range values {
		if key != "format" && key != "callback" && key != "api_sig" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	var plain strings.Builder
	for _, key := range keys {
		plain.WriteString(key)
		plain.WriteString(strings.Join(values[key], ""))
	}
	plain.WriteString(c.apiSecret)
	sum := md5.Sum([]byte(plain.String())) // #nosec G401 -- Last.fm mandates MD5.
	return hex.EncodeToString(sum[:])
}

type methodRule struct{ post, signed, session bool }

var methodRules = buildMethodRules()

func buildMethodRules() map[string]methodRule {
	methods := []string{
		"album.getInfo", "album.getTags", "album.getTopTags", "album.search",
		"artist.getCorrection", "artist.getInfo", "artist.getSimilar", "artist.getTags", "artist.getTopAlbums", "artist.getTopTags", "artist.getTopTracks", "artist.search",
		"auth.getSession", "auth.getToken", "chart.getTopArtists", "chart.getTopTags", "chart.getTopTracks",
		"geo.getTopArtists", "geo.getTopTracks", "library.getArtists", "tag.getInfo", "tag.getSimilar", "tag.getTopAlbums", "tag.getTopArtists", "tag.getTopTags", "tag.getTopTracks", "tag.getWeeklyChartList",
		"track.getCorrection", "track.getInfo", "track.getSimilar", "track.getTags", "track.getTopTags", "track.search",
		"user.getFriends", "user.getInfo", "user.getLovedTracks", "user.getPersonalTags", "user.getRecentTracks", "user.getTopAlbums", "user.getTopArtists", "user.getTopTags", "user.getTopTracks", "user.getWeeklyAlbumChart", "user.getWeeklyArtistChart", "user.getWeeklyChartList", "user.getWeeklyTrackChart",
	}
	rules := make(map[string]methodRule, 49)
	for _, method := range methods {
		rules[method] = methodRule{}
	}
	rules["auth.getToken"] = methodRule{signed: true}
	rules["auth.getSession"] = methodRule{signed: true}
	rules["auth.getMobileSession"] = methodRule{post: true, signed: true}
	for _, method := range []string{"album.addTags", "album.removeTag", "artist.addTags", "artist.removeTag", "track.addTags", "track.removeTag", "track.love", "track.unlove", "track.scrobble", "track.updateNowPlaying"} {
		rules[method] = methodRule{post: true, signed: true, session: true}
	}
	for _, method := range []string{
		"artist.getTopFans",
		"geo.getMetros",
		"tag.search",
		"tasteometer.compare",
		"track.getTopFans",
		"user.getNeighbours",
	} {
		rules[method] = methodRule{}
	}
	return rules
}

func cloneValues(values url.Values) url.Values {
	copy := make(url.Values, len(values))
	for key, entries := range values {
		copy[key] = append([]string(nil), entries...)
	}
	return copy
}

// HTTPError reports a non-successful HTTP response.
type HTTPError struct {
	StatusCode   int
	Status, Body string
}

// APIError is an error response returned by Last.fm.
type APIError struct {
	Code    int             `json:"error"`
	Message string          `json:"message"`
	Raw     json.RawMessage `json:"-"`
}

func (e *APIError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("lastfm: API error %d", e.Code)
	}
	return fmt.Sprintf("lastfm: API error %d: %s", e.Code, e.Message)
}

func (e *HTTPError) Error() string {
	if e.Body == "" {
		return fmt.Sprintf("lastfm: HTTP %d %s", e.StatusCode, e.Status)
	}
	return fmt.Sprintf("lastfm: HTTP %d %s: %s", e.StatusCode, e.Status, e.Body)
}

func mustURL(raw string) *url.URL {
	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}
	return u
}
