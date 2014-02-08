package lastfm

import (
	"encoding/xml"
	"strconv"
)

// UserClients
// Collection of methods that correspond to part of
// the LastFM's user\.(.+) methods, where the name
// of the Go method is \1 in CamelCase.
type UserClient struct {
	Client
}

// Used for most of the UserClient methods to prepare the query.
// Returns map[string]string that can be used with lastfm's makeRequest method.
func (c *UserClient) prepareQuery(user string, page, limit int) (query map[string]string) {
	query = make(map[string]string)
	query["user"] = user

	if page != 0 {
		query["page"] = strconv.Itoa(page)
	}

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	return
}

// Get full user information for some user.
// Returns UserInfoResponse or error.
func (c *UserClient) GetInfo(user string) (response UserInfoResponse, err error) {
	method := "user.getInfo"
	query := c.prepareQuery(user, 0, 0)
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

// Get loved tracks for some user.
// You can specify the limit and the page for the result.
// Returns LovedTracksResponse or error.
func (c *UserClient) GetLovedTracks(user string, page, limit int) (response LovedTracksResponse, err error) {
	method := "user.getLovedTracks"
	query := c.prepareQuery(user, page, limit)
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

// Get neighbours for some user. You may specify the limit of
// returned users.
// Returns NeighboursResponse or error.
func (c *UserClient) GetNeighbours(user string, limit int) (response NeighboursResponse, err error) {
	method := "user.getNeighbours"
	query := c.prepareQuery(user, 0, limit)
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

// Ger recent tracks listened by some user.
// You may specify the period (from, to) as UNIX timestamp.
// Also you may pass page and limit for the request.
// Returns RecentTracksResponse or error.
func (c *UserClient) GetRecentTracks(user string, page, limit int, from, to int64) (response RecentTracksResponse, err error) {
	method := "user.getRecentTracks"
	query := c.prepareQuery(user, page, limit)
	query["extended"] = "1"

	if from != 0 {
		query["from"] = strconv.FormatInt(from, 10)
	}

	if to != 0 {
		query["to"] = strconv.FormatInt(to, 10)
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

// Get top albums for some user.
// You may specify page and limit for the request.
// Also you may pass the period of the search.
// For more informacion about valid periods refer to
// lastfm's api documentation.
// Returns TopAlbumsResponse or error.
func (c *UserClient) GetTopAlbums(user, period string, page, limit int) (response TopAlbumsResponse, err error) {
	method := "user.getTopAlbums"
	query := c.prepareQuery(user, page, limit)

	if period != "" {
		query["period"] = period
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

// Get top artists for some user.
// You may specify page and limit for the request.
// Also you may pass the period of the search.
// For more informacion about valid periods refer to
// lastfm's api documentation.
// Returns TopArtistsResponse or error.
func (c *UserClient) GetTopArtists(user, period string, page, limit int) (response TopArtistsResponse, err error) {
	method := "user.getTopArtists"
	query := c.prepareQuery(user, page, limit)

	if period != "" {
		query["period"] = period
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

// Get top tracks listened by some user.
// You may specify page and limit for the request.
// Also you may pass the period of the search.
// For more informacion about valid periods refer to
// lastfm's api documentation.
// Returns TopTracksResponse or error.
func (c *UserClient) GetTopTracks(user, period string, page, limit int) (response TopTracksResponse, err error) {
	method := "user.getTopTracks"
	query := c.prepareQuery(user, page, limit)

	if period != "" {
		query["period"] = period
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

// Get top tags by some user.
// You may specify the limit of tags to be requested.
// Returns TopTagsResponse or error.
func (c *UserClient) GetTopTags(user string, limit int) (response TopTagsResponse, err error) {
	method := "user.getTopTags"
	query := c.prepareQuery(user, 0, limit)
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
