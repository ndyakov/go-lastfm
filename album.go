package lastfm

import (
	"encoding/xml"
	"strconv"
)

// AlbumClient
// Collection of methods that coresponde to most of
// LastFM's album\.(.+) methods.
// Where the name of the method is \1 in CammelCase.
type AlbumClient struct {
	Client
}

// Prepare query for most of the Album Requests.
// MBID has higher priority than Artist's and Album's name, so if MBID is present
// the names are ignored. Provide empty strings for missing data.
// Returns map[string]string that can be used in LastFM.makeRequest.
func (c *AlbumClient) prepareQuery(artist, album, mbid, username string, autocorrect int) (query map[string]string) {
	query = make(map[string]string)

	if mbid == "" {
		query["artist"] = artist
		query["album"] = album
	} else {
		query["mbid"] = mbid
	}

	if username != "" {
		query["username"] = username
	}

	query["autocorrect"] = strconv.Itoa(autocorrect)

	return
}

func (c *AlbumClient) GetInfo(artist, album, mbid, username string, autocorrect int) (response AlbumInfoResponse, err error) {
	method := "album.getInfo"
	query := c.prepareQuery(artist, album, mbid, username, autocorrect)
	body, _, err := c.lfm.makeRequest(method, query)

	if err != nil {
		return
	}

	defer body.Close()
	err = xml.NewDecoder(body).Decode(&response)

	if err != nil {
		return
	}

	if response.Error.Code != 0 {
		err = &response.Error
		return
	}

	return
}

func (c *AlbumClient) GetTags(artist, album, mbid, username string, autocorrect int) (response TagsResponse, err error) {
	return
}

func (c *AlbumClient) GetTopTags(artist, album, mbid, username string, autocorrect int) (response TopTagsResponse, err error) {
	return
}

func (c *AlbumClient) Search(album string, limit, page int) (response AlbumSearchResponse, err error) {
	return
}
