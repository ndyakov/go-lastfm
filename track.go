package lastfm

import (
	"strconv"
)

// TrackClient
// Collection of methods that correspond to most of
// LastFM's track\.(.+) methods.
// Where the name of the method is \` is CamelCase.
type TrackClient struct {
	Client
}

// Prepares query for most of the Track requests.
// MBID is with higher priority, so if present track
// and artist are igrnored, otherwise track and artist are used.
// Returns map[string]string that can be used for lastfm makeRequest.
func (c *TrackClient) prepareQuery(track, artist, mbid string, autocorrect int) (query map[string]string) {
	query = make(map[string]string)

	if mbid == "" {
		query["track"] = track
		query["artist"] = artist
	} else {
		query["mbid"] = mbid
	}

	query["autocorrect"] = strconv.Itoa(autocorrect)

	return
}

// Get full information for some track.
// Returns TrackInfoResponse or error.
// Be awere there may be an error from the xml decoding.
func (c *TrackClient) GetInfo(track, artist, mbid, user string, autocorrect int) (response *TrackInfoResponse, err error) {
	response = new(TrackInfoResponse)
	query := c.prepareQuery(track, artist, mbid, autocorrect)
	query["method"] = "track.getInfo"

	if user != "" {
		query["username"] = user
	}

	err = c.lfm.getResponse(query, response)

	return
}

// Get similar tracks to some track.
// Returns SimilarTracksResponse or error.
func (c *TrackClient) GetSimilar(track, artist, mbid string, autocorrect int) (response *SimilarTracksResponse, err error) {
	response = new(SimilarTracksResponse)
	query := c.prepareQuery(track, artist, mbid, autocorrect)
	query["method"] = "track.getSimilar"
	err = c.lfm.getResponse(query, response)

	return
}

// Get tags for some track tagged by some user.
// Returns TagsResponse or error.
func (c *TrackClient) GetTags(track, artist, mbid, user string, autocorrect int) (response *TagsResponse, err error) {
	response = new(TagsResponse)
	query := c.prepareQuery(track, artist, mbid, autocorrect)
	query["method"] = "track.getTags"

	if user != "" {
		query["user"] = user
	}

	err = c.lfm.getResponse(query, response)

	return
}

// Get top fans for some track.
// Returns TopFansResponse or error.
func (c *TrackClient) GetTopFans(track, artist, mbid string, autocorrect int) (response *TopFansResponse, err error) {
	response = new(TopFansResponse)
	query := c.prepareQuery(track, artist, mbid, autocorrect)
	query["method"] = "track.getTopFans"
	err = c.lfm.getResponse(query, response)

	return
}

// Get top tags in lastfm for some track.
// Returns TopTagsResponse or error.
func (c *TrackClient) GetTopTags(track, artist, mbid string, autocorrect int) (response *TopTagsResponse, err error) {
	response = new(TopTagsResponse)
	query := c.prepareQuery(track, artist, mbid, autocorrect)
	query["method"] = "track.getTopTags"
	err = c.lfm.getResponse(query, response)

	return
}

// Get correction for some track and artist.
// Returns TrackCorrectionResponse or error.
func (c *TrackClient) GetCorrection(track, artist string) (response *TrackCorrectionResponse, err error) {
	response = new(TrackCorrectionResponse)
	query := make(map[string]string)
	query["method"] = "track.getCorrection"
	query["track"] = track
	query["artist"] = artist
	err = c.lfm.getResponse(query, response)

	return
}

// Search for some track.
// artist parameter is optional, you may specify it to narrow your search
// otherwise pass empty string.
func (c *TrackClient) Search(track, artist string, page, limit int) (response *TrackSearchResponse, err error) {
	response = new(TrackSearchResponse)
	query := make(map[string]string)
	query["method"] = "track.search"
	query["track"] = track

	if artist != "" {
		query["artist"] = artist
	}

	if page != 0 {
		query["page"] = strconv.Itoa(page)
	}

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	err = c.lfm.getResponse(query, response)

	return
}
