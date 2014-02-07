package lastfm

type AlbumResponse struct {
	Name      string                `xml:"name"`
	Artist    ArtistResponse        `xml:"artist"`
	ID        int64                 `xml:"id"`
	Rank      int                   `xml:"rank,attr"`
	MBID      string                `xml:"mbid"`
	Listeners int64                 `xml:"listeners"`
	URL       string                `xml:"url"`
	Playcount int64                 `xml:"playcount"`
	Image     []LastfmImageResponse `xml:"image"`
}

type AlbumInfoResponse struct {
	LastfmStatusResponse
	AlbumResponse
	ReleaseDate string          `xml:"releasedate"`
	TopTags     TopTagsResponse `xml:"toptags"`
	Tracks      []TrackResponse `xml:"tracks>track"`
}

type TopAlbumsResponse struct {
	LastfmStatusResponse
	TopAlbums []AlbumResponse `xml:"topalbums>album"`
}

type AlbumSearchResponse struct {
	LastfmStatusResponse
	LastfmOpenSearchResponse
	AlbumMatches []AlbumResponse `xml:"albummatches>album"`
}
