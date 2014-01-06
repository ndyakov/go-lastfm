package lastfm

type TagClient struct {
	Client
}

func (c *TagClient) GetInfo(tag string) (response TagInfoResponse, err error) {
	return
}

func (c *TagClient) GetSimilar(tag string) (response TagSimilarResponse, err error) {
	return
}

func (c *TagClient) GetTopAlbums(tag string, limit, page int) (response TopAlbumsResponse, err error) {
	return
}

func (c *TagClient) GetTopArtists(tag string, limit, page int) (response TopArtistsResponse, err error) {
	return
}

func (c *TagClient) GetTopTags(tag string, limit, page int) (response TopTagsResponse, err error) {
	return
}

func (c *TagClient) GetTopTracks(tag string, limit, page int) (response TopTracksResponse, err error) {
	return
}

func (c *TagClient) Search(tag string, limit, page int) (response TagSearchResponse, err error) {
	return
}

//TODO : getWeeklyArtistChart
//TODO : getWeeklyChartList
