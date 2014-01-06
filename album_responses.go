package lastfm

type AlbumResponse struct {
	Name      string                `xml:"name"`
	Rank      int                   `xml:"rank,attr"`
	MBID      string                `xml:"mbid"`
	Listeners int64                 `xml:"listeners"`
	URL       string                `xml:"url"`
	Image     []LastfmImageResponse `xml:"image"`
}

type AlbumInfoResponse struct {
	LastfmStatusResponse
	AlbumResponse
	ArtistName  string          `xml:"artist"`
	ID          int64           `xml:"id"`
	ReleaseDate string          `xml:"releasedate"`
	Playcount   int64           `xml:"playcount"`
	TopTags     TopTagsResponse `xml:"toptags"`
	Tracks      []TrackResponse `xml:"tracks>track"`
}

type TopAlbumsResponse struct {
	LastfmStatusResponse
	TopAlbums []AlbumResponse `xml:"topalbums>album"`
}
