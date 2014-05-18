package lastfm

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
func (c *TrackClient) prepareQuery(track, artist string, optionalParams map[string]string) (query map[string]string) {
	query = optionalParams

	if _, ok := optionalParams["mbid"]; !ok {
		query["track"] = track
		query["artist"] = artist
	}

	return
}

// Get full information for some track.
// Returns TrackInfoResponse or error.
// Be awere there may be an error from the xml decoding.
func (c *TrackClient) GetInfo(track, artist string, optionalParams map[string]string) (response *TrackInfoResponse, err error) {
	response = new(TrackInfoResponse)
	query := c.prepareQuery(track, artist, optionalParams)
	query["method"] = "track.getInfo"
	err = c.lfm.getResponse(query, response)

	return
}

// Get similar tracks to some track.
// Returns SimilarTracksResponse or error.
func (c *TrackClient) GetSimilar(track, artist string, optionalParams map[string]string) (response *SimilarTracksResponse, err error) {
	response = new(SimilarTracksResponse)
	query := c.prepareQuery(track, artist, optionalParams)
	query["method"] = "track.getSimilar"
	err = c.lfm.getResponse(query, response)

	return
}

// Get tags for some track tagged by some user.
// Returns TagsResponse or error.
func (c *TrackClient) GetTags(track, artist string, optionalParams map[string]string) (response *TagsResponse, err error) {
	response = new(TagsResponse)
	query := c.prepareQuery(track, artist, optionalParams)
	query["method"] = "track.getTags"
	err = c.lfm.getResponse(query, response)

	return
}

// Get top fans for some track.
// Returns TopFansResponse or error.
func (c *TrackClient) GetTopFans(track, artist string, optionalParams map[string]string) (response *TopFansResponse, err error) {
	response = new(TopFansResponse)
	query := c.prepareQuery(track, artist, optionalParams)
	query["method"] = "track.getTopFans"
	err = c.lfm.getResponse(query, response)

	return
}

// Get top tags in lastfm for some track.
// Returns TopTagsResponse or error.
func (c *TrackClient) GetTopTags(track, artist string, optionalParams map[string]string) (response *TopTagsResponse, err error) {
	response = new(TopTagsResponse)
	query := c.prepareQuery(track, artist, optionalParams)
	query["method"] = "track.getTopTags"
	err = c.lfm.getResponse(query, response)

	return
}

// Get correction for some track and artist.
// Returns TrackCorrectionResponse or error.
func (c *TrackClient) GetCorrection(track, artist string) (response *TrackCorrectionResponse, err error) {
	response = new(TrackCorrectionResponse)
	query := c.prepareQuery(track, artist, make(map[string]string))
	query["method"] = "track.getCorrection"
	err = c.lfm.getResponse(query, response)

	return
}

// Search for some track.
// artist parameter is optional, you may specify it to narrow your search
// otherwise pass empty string.
func (c *TrackClient) Search(track string, optionalParams map[string]string) (response *TrackSearchResponse, err error) {
	response = new(TrackSearchResponse)
	query := optionalParams
	query["track"] = track
	query["method"] = "track.search"
	err = c.lfm.getResponse(query, response)

	return
}

func (c *TrackClient) Scrobble(artist, track, timestamp string, optionalParams map[string]string) (response *TrackScrobbleResponse, err error) {
	response = new(TrackScrobbleResponse)
	query := c.prepareQuery(track, artist, optionalParams)
	query["timestamp"] = timestamp
	query["method"] = "track.scrobble"
	err = c.lfm.getResponse(query, response)

	return
}
