package lastfm

// CompareResultResponse is the first half
// of TasteometerCompareResponse.
type CompareResultResponse struct {
	Score float64 `xml:"score"`
	// As this doesn't work in go for now, we will skip it
	// untill they fix the issue. Check here:
	// http://code.google.com/p/go/issues/detail?id=3688
	//Matches int              `xml:"artists>matches,attr"`
	Artists []ArtistResponse `xml:"artists>artist"`
}

// CompareInputResponse is the second half
// of the TasteometerCompareResponse.
type CompareInputResponse struct {
	Users []UserResponse `xml:"user"`
}

// TasteometerCompareResponse is used for the
// tasteometer.compare's response.
type TasteometerCompareResponse struct {
	LastfmStatusResponse
	Result CompareResultResponse `xml:"comparison>result"`
	Input  CompareInputResponse  `xml:"comparison>input"`
}
