package lastfm

type TrackStreamableResponse struct {
	FullTrack  int `xml:"fulltrack,attr"`
	Streamable int `xml",chardata"`
}

type TrackResponse struct {
	Name       string                  `xml:"name"`
	MBID       string                  `xml:"mbid"`
	URL        string                  `xml:"url"`
	Duration   int64                   `xml:"duration"`
	Streamable TrackStreamableResponse `xml:"streamable"`
	Listeners  int64                   `xml:"listeners"`
	Playcount  int64                   `xml:"playcount"`
	Artist     ArtistResponse          `xml:"artist"`
	Image      []LastfmImageResponse   `xml:"image"`
}

type TrackInfoResponse struct {
	LastfmStatusResponse
	ID int64 `xml:"id"`
	TrackResponse
	Album   AlbumResponse      `xml:"album"`
	TopTags TopTagsResponse    `xml:"toptags"`
	Wiki    LastfmWikiResponse `xml:"wiki"`
}

type TopTracksResponse struct {
	LastfmStatusResponse
	TopTracks []TrackResponse `xml:"toptracks>track"`
}
