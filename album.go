package lastfm

type AlbumClient struct {
	Client
}

func (c *AlbumClient) prepareQuery(artist, album, mbid, username string, autocorrect int) (query map[string]string) {
	return
}

func (c *AlbumClient) GetInfo(artist, album, mbid, username string, autocorrect int) (response AlbumInfoResponse, err error) {
	return
}

func (c *AlbumClient) GetTags(artist, album, mbid, username string, autocorrect int) (response TagsResponse, err error) {
	return
}

func (c *AlbumClient) GetTopTags(artist, album, mbid, username string, autocorrect int) (response TopTagsResponse, err error) {
	return
}

func (c *AlbumClient) Search(album string, limit, page int) (response AlbumSearchResponse, err error) {
	return
}
