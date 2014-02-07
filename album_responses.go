package lastfm

type AlbumResponse struct {
	Name        string                `xml:"name"`
	Artist      ArtistResponse        `xml:"artist"`
	ID          int64                 `xml:"id"`
	Rank        int                   `xml:"rank,attr"`
	MBID        string                `xml:"mbid"`
	Listeners   int64                 `xml:"listeners"`
	URL         string                `xml:"url"`
	Playcount   int64                 `xml:"playcount"`
	Image       []LastfmImageResponse `xml:"image"`
	ReleaseDate string                `xml:"releasedate"`
	TopTags     TopTagsResponse       `xml:"toptags"`
	Tracks      []TrackResponse       `xml:"tracks>track"`
}

type AlbumResponseNoArtistStruct struct {
	Name        string                `xml:"name"`
	Artist      string                `xml:"artist"`
	ID          int64                 `xml:"id"`
	Rank        int                   `xml:"rank,attr"`
	MBID        string                `xml:"mbid"`
	Listeners   int64                 `xml:"listeners"`
	URL         string                `xml:"url"`
	Playcount   int64                 `xml:"playcount"`
	Image       []LastfmImageResponse `xml:"image"`
	ReleaseDate string                `xml:"releasedate"`
	TopTags     TopTagsResponse       `xml:"toptags"`
	Tracks      []TrackResponse       `xml:"tracks>track"`
}

type AlbumInfoResponse struct {
	LastfmStatusResponse
	Album AlbumResponseNoArtistStruct `xml:"album"`
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
