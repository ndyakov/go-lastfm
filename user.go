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

func (c *UserClient) GetTopTags(user string, limit int) (response TopTagsResponse, err error) {
	return
}
