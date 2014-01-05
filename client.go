package lastfm

import (
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

var (
	apiRootURL = url.URL{Scheme: "http", Host: "ws.audioscrobbler.com", Path: "/2.0/"}
)

type client interface {
	Get(uri string) (resp *http.Response, err error)
}

type LastFM struct {
	apiKey string
	client client
}

func New(apiKey string) LastFM {
	if apiKey == "api_key_for_testing" {
		return LastFM{apiKey: apiKey, client: &DummyClient{}}
	} else {
		return LastFM{apiKey: apiKey, client: http.DefaultClient}
	}
}

type DummyClient struct{}

func (c *DummyClient) Get(uri string) (resp *http.Response, err error) {
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

func (c *DummyClient) buildFilename(values url.Values) string {
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
		parts = append(parts, strings.Join([]string{key, strings.Replace(strings.Join(values[key], ","), " ", ".", -1)}, "="))
	}
	return "testing/" + strings.Join(parts, "-") + ".xml"
}
