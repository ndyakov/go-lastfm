package lastfm

type ArtistBioResponse struct {
	LastfmWikiResponse
}

type ArtistResponse struct {
	LastfmStatusResponse
	Rank       int                   `xml:"rank,attr"`
	Name       string                `xml:"name"`
	Playcount  int64                 `xml:"playcount"`
	MBID       string                `xml:"mbid"`
	URL        string                `xml:"url"`
	Streamable int                   `xml:"streamable"`
	Image      []LastfmImageResponse `xml:"image"`
	Stats      LastfmStatsResponse   `xml:"stats"`
	Similar    []ArtistResponse      `xml:"similar"`
	Bio        ArtistBioResponse     `xml:"bio"`
}
