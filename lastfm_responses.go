package lastfm

type LastfmResponse interface{}

// LastfmStatusResponse, used to parse the status
// of any response from lastfm.
type LastfmStatusResponse struct {
	Status string              `xml:"status,attr"`
	Error  LastfmErrorResponse `xml:"error"`
}

// LastfmErrorResponse
// Holds errors in the status response,
// also implements error interface, so
// it can be returned as an error.
type LastfmErrorResponse struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}

// Return the error message.
func (e *LastfmErrorResponse) Error() string {
	return e.Message
}

// LastfmDateResponse to parse dates
// from the xml responses. Holds date as string
// and also UTS unix timestamp.
type LastfmDateResponse struct {
	Date string `xml:",chardata"`
	UTS  int64  `xml:"uts,attr"`
}

// LastfmImageResponse, usefull where multiple
// image tags are present.
type LastfmImageResponse struct {
	Size string `xml:"size,attr"`
	URL  string `xml:",chardata"`
}

// LastfmStatsResponse is used to parse
// some of the artist and album stats.
type LastfmStatsResponse struct {
	Listeners int64 `xml:"listeners"`
	Playcount int64 `xml:"playcount"`
}

// LastfmWikiResponse is used for any of the
// xxx.getInfo's responses. Holds data from the
// <bio> or <wiki> tags.
type LastfmWikiResponse struct {
	Published string `xml:"published"`
	Summary   string `xml:"summary"`
	Content   string `xml:"content"`
}

// LastfmOpenSearchResponse is used for any of the
// xxx.search's responses. Holds data about the search
// request - query, total results, start index, item per page.
type LastfmOpenSearchResponse struct {
	Query        OpenSearchQueryResponse `xml:"results>Query"`
	TotalResults int                     `xml:"results>totalResults"`
	StartIndex   int                     `xml:"results>startIndex"`
	ItemsPerPage int                     `xml:"results>itemsPerPage"`
}

// OpenSearchQueryResponse is part of
// LastfmOpenSearchResponse, holds data from the
// <opensearch:Query> tag.
type OpenSearchQueryResponse struct {
	Role        string `xml:"role,attr"`
	SearchTerms string `xml:"searchTerms,attr"`
	StartPage   int    `xml:"startPage,attr"`
}
