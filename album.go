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
func (c *AlbumClient) prepareQuery(artist, album, mbid, user string, autocorrect int, userLongVariableName bool) (query map[string]string) {
	query = make(map[string]string)

	if mbid == "" {
		query["artist"] = artist
		query["album"] = album
	} else {
		query["mbid"] = mbid
	}

	if user != "" {
		if userLongVariableName {
			query["username"] = user
		} else {
			query["user"] = user
		}
	}

	query["autocorrect"] = strconv.Itoa(autocorrect)

	return
}

func (c *AlbumClient) GetInfo(artist, album, mbid, username string, autocorrect int) (response AlbumInfoResponse, err error) {
	method := "album.getInfo"
	query := c.prepareQuery(artist, album, mbid, username, autocorrect, true)
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

func (c *AlbumClient) GetTags(artist, album, mbid, user string, autocorrect int) (response TagsResponse, err error) {
	method := "album.getTags"
	query := c.prepareQuery(artist, album, mbid, user, autocorrect, false)
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

func (c *AlbumClient) GetTopTags(artist, album, mbid string, autocorrect int) (response TopTagsResponse, err error) {
	method := "album.getTopTags"
	query := c.prepareQuery(artist, album, mbid, "", autocorrect, false)
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

func (c *AlbumClient) Search(album string, page, limit int) (response AlbumSearchResponse, err error) {
	method := "album.search"
	query := make(map[string]string)
	query["album"] = album

	if page != 0 {
		query["page"] = strconv.Itoa(page)
	}

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}

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
