package lastfm

type CompareResultResponse struct {
	Score float64 `xml:"score"`
	//TODO: how to catch <artists matches="xx" >
	Artists []ArtistResponse `xml:"artists>artist"`
}

type CompareInputResponse struct {
	Users []UserResponse `xml:"input>user"`
}
type TasteometerCompareResponse struct {
	LastfmStatusResponse
	Result CompareResultResponse `xml:"comparison>result"`
	Input  CompareInputResponse  `xml:"comparison>input"`
}
