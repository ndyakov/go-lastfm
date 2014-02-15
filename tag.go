package lastfm

import (
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

// Get full information for some tag.
// Returns TagInfoResponse or error.
func (c *TagClient) GetInfo(tag string) (response *TagInfoResponse, err error) {
	response = new(TagInfoResponse)
	method := "tag.getInfo"
	query := make(map[string]string)
	query["tag"] = tag
	err = c.lfm.getResponse(method, query, response)

	return
}

// Get similar tags to some tag.
// Returns TagSimilarResponse or error.
func (c *TagClient) GetSimilar(tag string) (response *TagSimilarResponse, err error) {
	response = new(TagSimilarResponse)
	method := "tag.getSimilar"
	query := make(map[string]string)
	query["tag"] = tag
	err = c.lfm.getResponse(method, query, response)

	return
}

// Get top Albums with some tag.
// Returns TopAlbumsResponse or error.
func (c *TagClient) GetTopAlbums(tag string, page, limit int) (response *TopAlbumsResponse, err error) {
	response = new(TopAlbumsResponse)
	method := "tag.getTopAlbums"
	query := c.prepareQuery(tag, page, limit)
	err = c.lfm.getResponse(method, query, response)

	return
}

// Get top Artists with some tag.
// Returns TopArtistsResponse or error.
func (c *TagClient) GetTopArtists(tag string, page, limit int) (response *TopArtistsResponse, err error) {
	response = new(TopArtistsResponse)
	method := "tag.getTopArtists"
	query := c.prepareQuery(tag, page, limit)
	err = c.lfm.getResponse(method, query, response)

	return
}

// Get Top Tags in Lastfm.
// Returns TopTagsResponse or error.
func (c *TagClient) GetTopTags() (response *TopTagsResponse, err error) {
	response = new(TopTagsResponse)
	method := "tag.getTopTags"
	query := make(map[string]string)
	err = c.lfm.getResponse(method, query, response)

	return
}

// Get Top Tracks with some tag.
// Returns TopTracksResponse or error.
func (c *TagClient) GetTopTracks(tag string, page, limit int) (response *TopTracksResponse, err error) {
	response = new(TopTracksResponse)
	method := "tag.getTopTracks"
	query := c.prepareQuery(tag, page, limit)
	err = c.lfm.getResponse(method, query, response)

	return
}

// Search tag by some string.
// Returns TagSearchResponse or error.
func (c *TagClient) Search(tag string, page, limit int) (response *TagSearchResponse, err error) {
	response = new(TagSearchResponse)
	method := "tag.search"
	query := c.prepareQuery(tag, page, limit)
	err = c.lfm.getResponse(method, query, response)

	return
}

//TODO : getWeeklyArtistChart
//TODO : getWeeklyChartList
