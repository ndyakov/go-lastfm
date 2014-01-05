package lastfm

type TrackStreamableResponse struct {
	FullTrack  int `xml:"fulltrack,attr"`
	Streamable int `xml",chardata"`
}
type TrackResponse struct {
	ID         int64                   `xml:"id"`
	Name       string                  `xml:"name"`
	MBID       string                  `xml:"mbid"`
	URL        string                  `xml:"url"`
	Duration   int64                   `xml:"duration"`
	Streamable TrackStreamableResponse `xml:"streamable"`
	Listeners  int64                   `xml:"listeners"`
	Playcount  int64                   `xml:"playcount"`
	Artist     ArtistResponse          `xml:"artist"`
	Album      AlbumResponse           `xml:"album"`
	TopTags    []TagResponse           `xml:"toptags"`
	Wiki       LastfmWikiResponse      `xml:"wiki"`
}
