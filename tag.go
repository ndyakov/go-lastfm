package lastfm

import (
	"encoding/xml"
	"strconv"
)

// TagClient
// Collection of methods that correspond to most of
// LastFM's tag\.(.+) methods.
// Where the name of the method is \1 in CamelCase.
type TagClient struct {
	Client
}

// Prepares query for most of the Tag Requests.
// Returns map[string]string that can be used with LastFM's makeRequest method.
func (c *TagClient) prepareQuery(tag string, page, limit int) (query map[string]string) {
	query = make(map[string]string)
	query["tag"] = tag

	if page != 0 {
		query["page"] = strconv.Itoa(page)
	}

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	return
}

func (c *TagClient) GetInfo(tag string) (response TagInfoResponse, err error) {
	method := "tag.getInfo"
	query := make(map[string]string)
	query["tag"] = tag
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

func (c *TagClient) GetSimilar(tag string) (response TagSimilarResponse, err error) {
	method := "tag.getSimilar"
	query := make(map[string]string)
	query["tag"] = tag
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

func (c *TagClient) GetTopAlbums(tag string, page, limit int) (response TopAlbumsResponse, err error) {
	return
}

func (c *TagClient) GetTopArtists(tag string, page, limit int) (response TopArtistsResponse, err error) {
	return
}

func (c *TagClient) GetTopTags(tag string, page, limit int) (response TopTagsResponse, err error) {
	return
}

func (c *TagClient) GetTopTracks(tag string, page, limit int) (response TopTracksResponse, err error) {
	return
}

func (c *TagClient) Search(tag string, page, limit int) (response TagSearchResponse, err error) {
	return
}

//TODO : getWeeklyArtistChart
//TODO : getWeeklyChartList
