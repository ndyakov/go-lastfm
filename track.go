package lastfm

import (
	"encoding/xml"
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
func (c *TrackClient) GetInfo(track, artist, mbid, user string, autocorrect int) (response TrackInfoResponse, err error) {
	method := "track.getInfo"
	query := c.prepareQuery(track, artist, mbid, autocorrect)

	if user != "" {
		query["username"] = user
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

// Get similar tracks to some track.
// Returns SimilarTracksResponse or error.
func (c *TrackClient) GetSimilar(track, artist, mbid string, autocorrect int) (response SimilarTracksResponse, err error) {
	method := "track.getSimilar"
	query := c.prepareQuery(track, artist, mbid, autocorrect)
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

// Get tags for some track tagged by some user.
// Returns TagsResponse or error.
func (c *TrackClient) GetTags(track, artist, mbid, user string, autocorrect int) (response TagsResponse, err error) {
	method := "track.getTags"
	query := c.prepareQuery(track, artist, mbid, autocorrect)

	if user != "" {
		query["user"] = user
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

// Get top fans for some track.
// Returns TopFansResponse or error.
func (c *TrackClient) GetTopFans(track, artist, mbid string, autocorrect int) (response TopFansResponse, err error) {
	method := "track.getTopFans"
	query := c.prepareQuery(track, artist, mbid, autocorrect)
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

// Get top tags in lastfm for some track.
// Returns TopTagsResponse or error.
func (c *TrackClient) GetTopTags(track, artist, mbid string, autocorrect int) (response TopTagsResponse, err error) {
	method := "track.getTopTags"
	query := c.prepareQuery(track, artist, mbid, autocorrect)
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

// Get correction for some track and artist.
// Returns TrackCorrectionResponse or error.
func (c *TrackClient) GetCorrection(track, artist string) (response TrackCorrectionResponse, err error) {
	method := "track.getCorrection"
	query := make(map[string]string)
	query["track"] = track
	query["artist"] = artist
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

// Search for some track.
// artist parameter is optional, you may specify it to narrow your search
// otherwise pass empty string.
func (c *TrackClient) Search(track, artist string, page, limit int) (response TrackSearchResponse, err error) {
	method := "track.search"
	query := make(map[string]string)
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
