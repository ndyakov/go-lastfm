package lastfm

import (
	"encoding/xml"
	"strconv"
)

// ArtistClient
// Collection of methods that correspond to most of
// LastFM's artist\.(.+) methods.
// Where the name of the method is \1 in CamelCase.
type ArtistClient struct {
	Client
}

// Prepares query for few of the Artist Requests.
// MBID has higher priority than Artist's name, so if MBID is present
// the name is ignored. Provide empty strings for missing data.
// Returns map[string]string that can be parsed to LastFM.makeRequest.
func (c *ArtistClient) prepareQuery(name, mbid string, autocorrect int) (query map[string]string) {
	query = make(map[string]string)

	if mbid == "" {
		query["artist"] = name
	} else {
		query["mbid"] = mbid
	}

	query["autocorrect"] = strconv.Itoa(autocorrect)

	return
}

// Get top tags for Artist.
// Returns TopTagsResponse structure or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTopTags(name, mbid string, autocorrect int) (response TopTagsResponse, err error) {
	method := "artist.getTopTags"
	query := c.prepareQuery(name, mbid, autocorrect)
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

// Get tags for Artist who is in users directory.
// Returns TagsResponse structure or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTags(name, mbid string, autocorrect int, user string) (response TagsResponse, err error) {
	method := "artist.getTags"
	query := c.prepareQuery(name, mbid, autocorrect)
	query["user"] = user
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

// Get top Albums for Artist.
// Returns TopAlbumsResponse or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTopAlbums(name, mbid string, autocorrect, page, limit int) (response TopAlbumsResponse, err error) {
	method := "artist.getTopAlbums"
	query := c.prepareQuery(name, mbid, autocorrect)

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

// Get top fans for Artist.
// Returns TopFansResponse or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTopFans(name, mbid string, autocorrect int) (response TopFansResponse, err error) {
	method := "artist.getTopFans"
	query := c.prepareQuery(name, mbid, autocorrect)
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

// Get top tracks for Artist.
// Returns TopTracksResponse or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTopTracks(name, mbid string, autocorrect, page, limit int) (response TopTracksResponse, err error) {
	method := "artist.getTopTracks"
	query := c.prepareQuery(name, mbid, autocorrect)

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

// Search for Artist by given name.
// Returns ArtistSearchResponse or error.
func (c *ArtistClient) Search(name string, page, limit int) (response ArtistSearchResponse, err error) {
	method := "artist.Search"
	query := make(map[string]string)
	query["artist"] = name

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

// Get similar artists.
// Returns ArtistSimilarResponse or error.
func (c *ArtistClient) GetSimilar(name string, mbid string, autocorrect int) (response ArtistSimilarResponse, err error) {
	method := "artist.getsimilar"
	query := c.prepareQuery(name, mbid, autocorrect)
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

// Get info for Artist with given name.
// Returns ArtistInfoResponse or error.
// Response data is actually in response.Artist.
func (c *ArtistClient) GetInfo(name, mbid string, autocorrect int) (response ArtistInfoResponse, err error) {
	method := "artist.getInfo"
	query := c.prepareQuery(name, mbid, autocorrect)
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
