package lastfm

type ArtistBio struct {
	LastfmWiki
}

type Artist struct {
	Rank       int           `xml:"rank,attr"`
	Name       string        `xml:"name"`
	Playcount  int64         `xml:"playcount"`
	MBID       string        `xml:"mbid"`
	URL        string        `xml:"url"`
	Streamable int           `xml:"streamable"`
	Image      []LastfmImage `xml:"image"`
	Stats      LastfmStats   `xml:"stats"`
	Similar    []Artist      `xml:"similar"`
	Bio        ArtistBio     `xml:"bio"`
}
