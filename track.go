package lastfm

type TrackStreamable struct {
	FullTrack  int `xml:"fulltrack,attr"`
	Streamable int `xml",chardata"`
}
type Track struct {
	ID         int64           `xml:"id"`
	Name       string          `xml:"name"`
	MBID       string          `xml:"mbid"`
	URL        string          `xml:"url"`
	Duration   int64           `xml:"duration"`
	Streamable TrackStreamable `xml:"streamable"`
	Listeners  int64           `xml:"listeners"`
	Playcount  int64           `xml:"playcount"`
	Artist     Artist          `xml:"artist"`
	Album      Album           `xml:"album"`
	TopTags    []Tag           `xml:"toptags"`
	Wiki       LastfmWiki      `xml:"wiki"`
}
