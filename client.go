package lastfm

import (
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
	apiRootURL = url.URL{Scheme: "http", Host: "ws.audioscrobbler.com", Path: "/2.0/"}
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
	getter      getter
	Album       AlbumClient
	Artist      ArtistClient
	Tag         TagClient
	Track       TrackClient
	User        UserClient
	Tasteometer TasteometerClient
}

// Initialize all clients and
// chooses which getter to use.
func New(apiKey string) *LastFM {
	lfm := new(LastFM)
	lfm.apiKey = apiKey
	lfm.Album = AlbumClient{Client: Client{lfm}}
	lfm.Artist = ArtistClient{Client: Client{lfm}}
	lfm.Tag = TagClient{Client: Client{lfm}}
	lfm.Track = TrackClient{Client: Client{lfm}}
	lfm.User = UserClient{Client: Client{lfm}}
	lfm.Tasteometer = TasteometerClient{Client: Client{lfm}}

	if apiKey == "api_key_for_testing" {
		lfm.getter = new(dummyGetter)
	} else {
		lfm.getter = http.DefaultClient
	}

	return lfm
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

func (lfm *LastFM) getResponse(method string, query map[string]string, response LastfmResponse) (err error) {
	body, _, err := lfm.makeRequest(method, query)

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
func (lfm *LastFM) makeRequest(method string, params map[string]string) (body io.ReadCloser, hdr http.Header, err error) {
	queryParams := make(map[string]string, len(params)+2)
	queryParams["api_key"] = lfm.apiKey
	queryParams["method"] = method

	for key, value := range params {
		queryParams[key] = value
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

	for key, _ := range values {

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
