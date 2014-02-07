package lastfm

type LastfmResponse interface{}

type LastfmStatusResponse struct {
	Status string              `xml:"status,attr"`
	Error  LastfmErrorResponse `xml:"error"`
}

type LastfmErrorResponse struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}

func (e *LastfmErrorResponse) Error() string {
	return e.Message
}

type LastfmDateResponse struct {
	Date string `xml:",chardata"`
	UTS  int64  `xml:"uts,attr"`
}

type LastfmImageResponse struct {
	Size string `xml:"size,attr"`
	URL  string `xml:",chardata"`
}

type LastfmStatsResponse struct {
	Listeners int64 `xml:"listeners"`
	Playcount int64 `xml:"playcount"`
}

type LastfmWikiResponse struct {
	Published string `xml:"published"`
	Summary   string `xml:"summary"`
	Content   string `xml:"content"`
}

type LastfmOpenSearchResponse struct {
	Query        OpenSearchQueryResponse `xml:"results>Query"`
	TotalResults int                     `xml:"results>totalResults"`
	StartIndex   int                     `xml:"results>startIndex"`
	ItemsPerPage int                     `xml:"results>itemsPerPage"`
}

type OpenSearchQueryResponse struct {
	Role        string `xml:"role,attr"`
	SearchTerms string `xml:"searchTerms,attr"`
	StartPage   int    `xml:"startPage,attr"`
}
