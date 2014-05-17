package lastfm

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

// lastfm's api root url
var (
	apiRootURL     = url.URL{Scheme: "https", Host: "ws.audioscrobbler.com", Path: "/2.0/"}
	apiWebURL      = url.URL{Scheme: "https", Host: "www.last.fm", Path: "/api/auth/"}
	needsSignature = map[string]bool{
		"auth.getSession":       true,
		"auth.getMobileSession": true,
		"auth.getToken":         true,
	}
)

// getter interface present
// so we can swap getters while testing
type getter interface {
	Get(uri string) (resp *http.Response, err error)
}

// Client struct that holds pointer
// to the LastFM object.
type Client struct {
	lfm *LastFM
}

// LastFM struct with all available clients.
type LastFM struct {
	apiKey      string
	apiSecret   string
	sessionKey  string
	userAgent   string
	token       string
	getter      getter
	Album       AlbumClient
	Artist      ArtistClient
	Tag         TagClient
	Track       TrackClient
	User        UserClient
	Tasteometer TasteometerClient
	Auth        AuthClient
}

// Initialize all clients and
// chooses which getter to use.
func New(apiKey, apiSecret string) *LastFM {
	lfm := new(LastFM)
	lfm.apiKey = apiKey
	lfm.apiSecret = apiSecret
	lfm.userAgent = "go-lastfm"
	lfm.Album = AlbumClient{Client: Client{lfm}}
	lfm.Artist = ArtistClient{Client: Client{lfm}}
	lfm.Tag = TagClient{Client: Client{lfm}}
	lfm.Track = TrackClient{Client: Client{lfm}}
	lfm.User = UserClient{Client: Client{lfm}}
	lfm.Tasteometer = TasteometerClient{Client: Client{lfm}}
	lfm.Auth = AuthClient{Client: Client{lfm}}

	if apiKey == "api_key_for_testing" {
		lfm.getter = new(dummyGetter)
	} else {
		lfm.getter = http.DefaultClient
	}

	return lfm
}

func (lfm *LastFM) SetSessionKey(sessionKey string) {
	lfm.sessionKey = sessionKey
}

func (lfm *LastFM) GetSessionKey() string {
	return lfm.sessionKey
}

func (lfm *LastFM) SetUserAgent(userAgent string) {
	lfm.userAgent = userAgent
}

func (lfm *LastFM) GetUserAgent() string {
	return lfm.userAgent
}

func (lfm *LastFM) SetToken(token string) {
	lfm.token = token
}

func (lfm *LastFM) GetToken() string {
	return lfm.token
}

func (lfm *LastFM) getSignature(params map[string]string) string {
	var plainSignature string
	var keys []string

	for key := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		plainSignature += key + params[key]
	}

	plainSignature += lfm.apiSecret

	hasher := md5.New()
	hasher.Write([]byte(plainSignature))

	return hex.EncodeToString(hasher.Sum(nil))
}

// Build url for the request.
func (lfm *LastFM) buildURL(query map[string]string) string {
	values := url.Values{}

	for key, value := range query {
		values.Add(key, value)
	}

	uri := apiRootURL
	uri.RawQuery = values.Encode()

	return uri.String()
}

func (lfm *LastFM) getResponse(query map[string]string, response LastfmResponse) (err error) {
	body, _, err := lfm.makeRequest(query)

	if err != nil {
		return
	}

	defer body.Close()
	err = xml.NewDecoder(body).Decode(response)

	if err != nil {
		return
	}

	if response.getErrorCode() != 0 {
		err = response.getError()
		return
	}
	return
}

// Add metthod and api_key as part of the query parameters.
// Call getter to get the requests body.
func (lfm *LastFM) makeRequest(params map[string]string) (body io.ReadCloser, hdr http.Header, err error) {
	queryParams := make(map[string]string, len(params)+2)
	queryParams["api_key"] = lfm.apiKey

	for key, value := range params {
		queryParams[key] = value
	}

	if _, ok := needsSignature[queryParams["method"]]; ok {
		queryParams["api_sig"] = lfm.getSignature(queryParams)
	}

	response, err := lfm.getter.Get(lfm.buildURL(queryParams))

	if err != nil {

		if response != nil && response.Body != nil {
			response.Body.Close()
		}

		return
	}

	return response.Body, response.Header, err
}

// dummyGetter used for testing.
// Implements getter interface.
type dummyGetter struct{}

// Getter that reads from some file and then returns it
// as body of http.Response.
func (c *dummyGetter) Get(uri string) (resp *http.Response, err error) {
	u, err := url.Parse(uri)

	if err != nil {
		return
	}

	fh, err := os.Open(c.buildFilename(u.Query()))

	if err != nil {
		return
	}

	resp = &http.Response{Body: fh}

	return
}

// Construct unique filename per request.
func (c *dummyGetter) buildFilename(values url.Values) string {
	var parts []string
	var keys []string
	parts = append(parts, values.Get("method"))

	for key := range values {

		if key == "method" || key == "api_key" {
			continue
		}

		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		parts = append(parts, strings.Join([]string{key, strings.Replace(strings.Join(values[key], ","), " ", "-", -1)}, "="))
	}

	return "testing/" + strings.Join(parts, "-") + ".xml"
}
