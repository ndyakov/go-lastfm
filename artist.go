package lastfm

import (
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
func (c *ArtistClient) GetTopTags(name, mbid string, autocorrect int) (response *TopTagsResponse, err error) {
	response = new(TopTagsResponse)
	query := c.prepareQuery(name, mbid, autocorrect)
	query["method"] = "artist.getTopTags"
	err = c.lfm.getResponse(query, response)

	return
}

// Get tags for Artist who is in users directory.
// Returns TagsResponse structure or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTags(name, mbid, user string, autocorrect int) (response *TagsResponse, err error) {
	response = new(TagsResponse)
	query := c.prepareQuery(name, mbid, autocorrect)
	query["method"] = "artist.getTags"
	query["user"] = user
	err = c.lfm.getResponse(query, response)

	return
}

// Get top Albums for Artist.
// Returns TopAlbumsResponse or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTopAlbums(name, mbid string, autocorrect, page, limit int) (response *TopAlbumsResponse, err error) {
	response = new(TopAlbumsResponse)
	query := c.prepareQuery(name, mbid, autocorrect)
	query["method"] = "artist.getTopAlbums"

	if page != 0 {
		query["page"] = strconv.Itoa(page)
	}

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	err = c.lfm.getResponse(query, response)

	return
}

// Get top fans for Artist.
// Returns TopFansResponse or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTopFans(name, mbid string, autocorrect int) (response *TopFansResponse, err error) {
	response = new(TopFansResponse)
	query := c.prepareQuery(name, mbid, autocorrect)
	query["method"] = "artist.getTopFans"
	err = c.lfm.getResponse(query, response)

	return
}

// Get top tracks for Artist.
// Returns TopTracksResponse or error.
// Be careful, there may be error returned from the parsing as well.
func (c *ArtistClient) GetTopTracks(name, mbid string, autocorrect, page, limit int) (response *TopTracksResponse, err error) {
	response = new(TopTracksResponse)
	query := c.prepareQuery(name, mbid, autocorrect)
	query["method"] = "artist.getTopTracks"

	if page != 0 {
		query["page"] = strconv.Itoa(page)
	}

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	err = c.lfm.getResponse(query, response)

	return
}

// Search for Artist by given name.
// Returns ArtistSearchResponse or error.
func (c *ArtistClient) Search(name string, page, limit int) (response *ArtistSearchResponse, err error) {
	response = new(ArtistSearchResponse)
	query := make(map[string]string)
	query["method"] = "artist.Search"
	query["artist"] = name

	if page != 0 {
		query["page"] = strconv.Itoa(page)
	}

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	err = c.lfm.getResponse(query, response)

	return
}

// Get similar artists.
// Returns ArtistSimilarResponse or error.
func (c *ArtistClient) GetSimilar(name string, mbid string, autocorrect int) (response *ArtistSimilarResponse, err error) {
	response = new(ArtistSimilarResponse)
	query := c.prepareQuery(name, mbid, autocorrect)
	query["method"] = "artist.getsimilar"
	err = c.lfm.getResponse(query, response)

	return
}

// Get info for Artist with given name.
// Returns ArtistInfoResponse or error.
// Response data is actually in response.Artist.
func (c *ArtistClient) GetInfo(name, mbid string, autocorrect int) (response *ArtistInfoResponse, err error) {
	response = new(ArtistInfoResponse)
	query := c.prepareQuery(name, mbid, autocorrect)
	query["method"] = "artist.getInfo"
	err = c.lfm.getResponse(query, response)

	return
}
