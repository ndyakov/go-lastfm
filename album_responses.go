package lastfm

// AlbumResponse, used where <album> tag is present
// and there is a tree of tags in <artist></artist> pair.
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

// AlbumResponseNoArtistStruct, used where <album> tag is
// present but there is no subtags in <artist></artist>.
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

// AlbumInlineResponse used where <album mbid="xx">album name</album>
// tag is present. For example in user.getRecentTracks
type AlbumInlineResponse struct {
	Name string `xml:",chardata"`
	MBID string `xml:"mbid,attr"`
}

// AlbumInfoResponse, used for album.getInfo request.
type AlbumInfoResponse struct {
	LastfmStatusResponse
	Album AlbumResponseNoArtistStruct `xml:"album"`
}

// TopAlbumResponse, used where <topalbums> tag is present.
type TopAlbumsResponse struct {
	LastfmStatusResponse
	TopAlbums []AlbumResponse `xml:"topalbums>album"`
}

// AlbumSearchResponse, used for album.search request.
type AlbumSearchResponse struct {
	LastfmStatusResponse
	LastfmOpenSearchResponse
	AlbumMatches []AlbumResponseNoArtistStruct `xml:"results>albummatches>album"`
}
