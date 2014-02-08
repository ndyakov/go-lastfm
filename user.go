package lastfm

import (
	"encoding/xml"
)

// UserClients
// Collection of methods that correspond to part of
// the LastFM's user\.(.+) methods, where the name
// of the Go method is \1 in CamelCase.
type UserClient struct {
	Client
}

func (c *UserClient) GetInfo(user string) (response UserInfoResponse, err error) {
	method := "user.getInfo"
	query := make(map[string]string)
	query["user"] = user
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
	return
}

func (c *UserClient) GetNeighbours(user string, limit int) (response NeighboursResponse, err error) {
	return
}

func (c *UserClient) GetRecentTracks(user string, page, limit int, from, to int64, extended int) (response RecentTracksResponse, err error) {
	return
}

func (c *UserClient) GetTopAlbums(user, period string, page, limit int) (response TopAlbumsResponse, err error) {
	return
}

func (c *UserClient) GetTopArtists(user, period string, page, limit int) (response TopArtistsResponse, err error) {
	return
}

func (c *UserClient) GetTopTags(user string, limit int) (response TopTagsResponse, err error) {
	return
}
func (c *UserClient) GetTopTracks(user, period string, page, limit int) (response TopTracksResponse, err error) {
	return
}
