package lastfm

import (
	"strconv"
)

// AlbumClient
// Collection of methods that correspond to most of
// LastFM's album\.(.+) methods.
// Where the name of the method is \1 in CamelCase.
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

// Get full information for Album.
// Returns AlbumInforResponse or error.
// There may be an error returned from the parser/decoder as well.
func (c *AlbumClient) GetInfo(artist, album, mbid, username string, autocorrect int) (response *AlbumInfoResponse, err error) {
	response = new(AlbumInfoResponse)
	query := c.prepareQuery(artist, album, mbid, username, autocorrect, true)
	query["method"] = "album.getInfo"
	err = c.lfm.getResponse(query, response)

	return
}

// Get Tags for some Album, that are added by some user.
// Returns TagsResponse or error.
func (c *AlbumClient) GetTags(artist, album, mbid, user string, autocorrect int) (response *TagsResponse, err error) {
	response = new(TagsResponse)
	query := c.prepareQuery(artist, album, mbid, user, autocorrect, false)
	query["method"] = "album.getTags"
	err = c.lfm.getResponse(query, response)

	return
}

// Get Top Tags for Album.
// Returns TopTagsResponse or error.
func (c *AlbumClient) GetTopTags(artist, album, mbid string, autocorrect int) (response *TopTagsResponse, err error) {
	response = new(TopTagsResponse)
	query := c.prepareQuery(artist, album, mbid, "", autocorrect, false)
	query["method"] = "album.getTopTags"
	err = c.lfm.getResponse(query, response)

	return
}

// Search album by given name. You can specify page and limit also.
// Default values are as stated in lastfm's api documentation.
// Returns AlbumSearchResponse.
func (c *AlbumClient) Search(album string, page, limit int) (response *AlbumSearchResponse, err error) {
	response = new(AlbumSearchResponse)
	query := make(map[string]string)
	query["album"] = album
	query["method"] = "album.search"

	if page != 0 {
		query["page"] = strconv.Itoa(page)
	}

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	err = c.lfm.getResponse(query, response)

	return
}
