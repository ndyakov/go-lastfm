package lastfm

import (
	"encoding/xml"
	"strconv"
)

type ArtistClient struct {
	Client
}

func (c *ArtistClient) prepareQuery(name, mbid string, autocorrect int) (query map[string]string){
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

func (c *ArtistClient) GetTopTags(name string, mbid string, autocorrect int) (response TopTagsResponse, err error) {
	method := "artist.getTopTags"
	query := c.prepareQuery(name,mbid,autocorrect)
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

func (c *ArtistClient) GetSimilar(name string, mbid string, autocorrect int) (response ArtistSimilarResponse, err error) {
	method := "artist.getsimilar"
	query := c.prepareQuery(name,mbid,autocorrect)
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
