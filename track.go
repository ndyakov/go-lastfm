package lastfm

type TrackClient struct {
	Client
}

func (c *TrackClient) prepareQuery(track, artist, mbid string, autocorrect int) (query map[string]string) {
	return
}

func (c *TrackClient) GetInfo(track, artist, mbid, username string, autocorrect int) (response TrackInfoResponse, err error) {
	return
}

func (c *TrackClient) GetSimilar(track, artist, mbid string, autocorrect int) (response TrackSimilarResponse, err error) {
	return
}

func (c *TrackClient) GetTags(track, artist, mbid string, autocorrect int) (response TagsResponse, err error) {
	return
}

func (c *TrackClient) GetTopFans(track, artist, mbid string, autocorrect int) (response TopFansResponse, err error) {
	return
}

func (c *TrackClient) GetTopTags(track, artist, mbid string, autocorrect int) (response TopTagsResponse, err error) {
	return
}

func (c *TrackClient) GetCorrection(track, artist string) (response TrackCorrectionResponse, err error) {
	return
}

func (c *TrackClient) Search(track, artist string, limit, page int) (response TrackSearchResponse, err error) {
	return
}
