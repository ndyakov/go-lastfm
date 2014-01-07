package lastfm

type UserClient struct {
	Client
}

func (c *UserClient) GetInfo(user string) (response UserInfoResponse, err error) {
	return
}

func (c *UserClient) GetLovedTracks(user string, limit, page int) (response TracksLovedResponse, err error) {
	return
}

func (c *UserClient) GetNeighbours(user string, limit int) (response NeighboursResponse, err error) {
	return
}

func (c *UserClient) GetRecentTracks(user string, limit, page int, from, to int64, extended int) (response TracksRecentResponse, err error) {
	return
}

func (c *UserClient) GetTopAlbums(user, period string, limit, page int) (response TopAlbumsResponse, err error) {
	return
}

func (c *UserClient) GetTopArtists(user, period string, limit, page int) (response TopArtistsResponse, err error) {
	return
}

func (c *UserClient) GetTopTags(user string, limit int) (response TopTagsResponse, err error) {
	return
}
func (c *UserClient) GetTopTracks(user, period string, limit, page int) (response TopTracksResponse, err error) {
	return
}
