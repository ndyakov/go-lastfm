package lastfm

import (
	"encoding/xml"
	"strconv"
)

// Type TasteometerClient holds Compare method
// for user to user comparison.
type TasteometerClient struct {
	Client
}

// Compare two users or two lists of artists (up to 100 artists per list, comma separated).
// Returns TasteometerCompareResponse or error
func (c *TasteometerClient) Compare(type1, value1, type2, value2 string, limit int) (response TasteometerCompareResponse, err error) {
	method := "tasteometer.compare"
	query := make(map[string]string)
	query["type1"] = type1
	query["value1"] = value1
	query["type2"] = type2
	query["value2"] = value2

	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
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
