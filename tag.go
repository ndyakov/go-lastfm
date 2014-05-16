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
	query := make(map[string]string)
	query["method"] = "tag.getInfo"
	query["tag"] = tag
	err = c.lfm.getResponse(query, response)

	return
}

// Get similar tags to some tag.
// Returns TagSimilarResponse or error.
func (c *TagClient) GetSimilar(tag string) (response *TagSimilarResponse, err error) {
	response = new(TagSimilarResponse)
	query := make(map[string]string)
	query["method"] = "tag.getSimilar"
	query["tag"] = tag
	err = c.lfm.getResponse(query, response)

	return
}

// Get top Albums with some tag.
// Returns TopAlbumsResponse or error.
func (c *TagClient) GetTopAlbums(tag string, page, limit int) (response *TopAlbumsResponse, err error) {
	response = new(TopAlbumsResponse)
	query := c.prepareQuery(tag, page, limit)
	query["method"] = "tag.getTopAlbums"
	err = c.lfm.getResponse(query, response)

	return
}

// Get top Artists with some tag.
// Returns TopArtistsResponse or error.
func (c *TagClient) GetTopArtists(tag string, page, limit int) (response *TopArtistsResponse, err error) {
	response = new(TopArtistsResponse)
	query := c.prepareQuery(tag, page, limit)
	query["method"] = "tag.getTopArtists"
	err = c.lfm.getResponse(query, response)

	return
}

// Get Top Tags in Lastfm.
// Returns TopTagsResponse or error.
func (c *TagClient) GetTopTags() (response *TopTagsResponse, err error) {
	response = new(TopTagsResponse)
	query := make(map[string]string)
	query["method"] = "tag.getTopTags"
	err = c.lfm.getResponse(query, response)

	return
}

// Get Top Tracks with some tag.
// Returns TopTracksResponse or error.
func (c *TagClient) GetTopTracks(tag string, page, limit int) (response *TopTracksResponse, err error) {
	response = new(TopTracksResponse)
	query := c.prepareQuery(tag, page, limit)
	query["method"] = "tag.getTopTracks"
	err = c.lfm.getResponse(query, response)

	return
}

// Search tag by some string.
// Returns TagSearchResponse or error.
func (c *TagClient) Search(tag string, page, limit int) (response *TagSearchResponse, err error) {
	response = new(TagSearchResponse)
	query := c.prepareQuery(tag, page, limit)
	query["method"] = "tag.search"
	err = c.lfm.getResponse(query, response)

	return
}

//TODO : getWeeklyArtistChart
//TODO : getWeeklyChartList
