package lastfm

import (
	"encoding/xml"
	"strconv"
)

type artistClient struct {
	client
}

func (c *artistClient) GetTopTags(name string, mbid string, autocorrect int) (response TopTagsResponse, err error) {
	method := "artist.getTopTags"
	query := make(map[string]string)
	if mbid == "" {
		query["artist"] = name
	}
	if mbid != "" {
		query["mbid"] = mbid
	}
	query["autocorrect"] = strconv.Itoa(autocorrect)

	body, _, err := c.lfm.makeRequest(method, query)
	if err != nil {
		return
	}
	defer body.Close()
	response = TopTagsResponse{}
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
