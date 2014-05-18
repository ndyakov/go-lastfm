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
	apiRootURL        = url.URL{Scheme: "https", Host: "ws.audioscrobbler.com", Path: "/2.0/"}
	apiWebURL         = url.URL{Scheme: "https", Host: "www.last.fm", Path: "/api/auth/"}
	requiresSignature = map[string]bool{
		"auth.getSession":       true,
		"auth.getMobileSession": true,
		"auth.getToken":         true,
	}

	requiresPost = map[string]bool{
		"auth.getMobileSession": true,
	}
)

// getter interface present
// so we can swap getters while testing
type getter interface {
	Get(uri string) (resp *http.Response, err error)
	PostForm(uri string, data url.Values) (resp *http.Response, err error)
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
		lfm.getter = new(dummyClient)
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

func (lfm *LastFM) simpleMerge(destination, source map[string]string) map[string]string {
	for key, value := range source {
		destination[key] = value
	}

	return destination
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
	values := lfm.buildValues(query)

	uri := apiRootURL
	uri.RawQuery = values.Encode()

	return uri.String()
}

func (lfm *LastFM) buildValues(query map[string]string) url.Values {
	values := url.Values{}

	for key, value := range query {
		values.Add(key, value)
	}

	return values
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

// Add api_key as part of the query parameters.
// Also add api_sig if requered.
// Call getter to get the requests body.
func (lfm *LastFM) makeRequest(params map[string]string) (body io.ReadCloser, hdr http.Header, err error) {
	var response *http.Response
	params["api_key"] = lfm.apiKey

	for key, value := range params {
		params[key] = value
	}

	if _, ok := requiresSignature[params["method"]]; ok {
		params["api_sig"] = lfm.getSignature(params)
	}

	if _, ok := requiresPost[params["method"]]; ok {
		response, err = lfm.getter.PostForm(apiRootURL.String(), lfm.buildValues(params))
	} else {
		response, err = lfm.getter.Get(lfm.buildURL(params))
	}

	if err != nil {

		if response != nil && response.Body != nil {
			response.Body.Close()
		}

		return
	}

	return response.Body, response.Header, err
}

// dummyClient used for testing.
// Implements getter interface.
type dummyClient struct{}

func (c *dummyClient) PostForm(uri string, params url.Values) (resp *http.Response, err error) {
	fh, err := os.Open(c.buildFilename(params))

	if err != nil {
		return
	}

	resp = &http.Response{Body: fh}

	return
}

// Getter that reads from some file and then returns it
// as body of http.Response.
func (c *dummyClient) Get(uri string) (resp *http.Response, err error) {
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
func (c *dummyClient) buildFilename(values url.Values) string {
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
