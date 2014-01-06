package lastfm

import (
	"encoding/xml"
	"strconv"
)

// ArtistClient
// Collection of methods that coresponde to most of
// LastFM's artist\.(.+) methods.
// Where the name of the method is \1 in CammelCase.
type ArtistClient struct {
	Client
}

// Prepares query for few of the Artist Requests.
// MBID has higher priority than Artist's name, so if MBID is present
// the name is ignored. Provide empty strings for missing data.
// Returns map[string]string that can be parsed to LastFM.makeRequest.
func (c *ArtistClient) prepareQuery(name, mbid string, autocorrect int) (query map[string]string) {
	query = make(map[string]string)
	if mbid == "" {
		query["artist"] = name
	}
	if mbid != "" {
		query["mbid"] = mbid
	}
	query["autocorrect"] = strconv.Itoa(autocorrect)
	return
}

// Get top tags for Artist.
// Returns TopTagsResponse structure or error.
func (c *ArtistClient) GetTopTags(name, mbid string, autocorrect int) (response TopTagsResponse, err error) {
	method := "artist.getTopTags"
	query := c.prepareQuery(name, mbid, autocorrect)
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

func (c *ArtistClient) GetTags(name, mbid string, autocorrect int) (response TagsResponse, err error) {
	return
}

func (c *ArtistClient) GetTopAlbums(name, mbid string, autocorrect, page, limit int) (response TopAlbumsResponse, err error) {
	return
}

func (c *ArtistClient) GetTopFans(name, mbid string, autocorrect int) (response TopFansResponse, err error) {
	return
}

func (c *ArtistClient) GetTopTracks(name, mbid string, autocorrect, page, limit int) (response TopTracksResponse, err error) {
	return
}

func (c *ArtistClient) Search(name string, limit, page int) (response ArtistSearchResponse, err error) {
	return
}

// Get similar artists.
// Returns ArtistSimilarResponse or error.
func (c *ArtistClient) GetSimilar(name string, mbid string, autocorrect int) (response ArtistSimilarResponse, err error) {
	method := "artist.getsimilar"
	query := c.prepareQuery(name, mbid, autocorrect)
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

func (c *ArtistClient) GetInfo(name, mbid string, autocorrect int) (response ArtistInfoResponse, err error) {
	return
}
