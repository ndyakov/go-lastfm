package lastfm

// TrackStreamableResponse is part of TrackResponse.
type TrackStreamableResponse struct {
	FullTrack  int `xml:"fulltrack,attr"`
	Streamable int `xml:",chardata"`
}

// TrackResponse is used where <track> tag is present
// and there is a tree of tags for <artist>.
type TrackResponse struct {
	Name       string                  `xml:"name"`
	ID         int64                   `xml:"id"`
	MBID       string                  `xml:"mbid"`
	URL        string                  `xml:"url"`
	Duration   int64                   `xml:"duration"`
	Streamable TrackStreamableResponse `xml:"streamable"`
	Listeners  int64                   `xml:"listeners"`
	Playcount  int64                   `xml:"playcount"`
	Artist     ArtistResponse          `xml:"artist"`
	Image      []LastfmImageResponse   `xml:"image"`
	UserLoved  int                     `xml:"userloved"`
	Album      AlbumResponse           `xml:"album"`
	TopTags    TopTagsResponse         `xml:"toptags"`
	Wiki       LastfmWikiResponse      `xml:"wiki"`
}

// TrackResponseNoArtistStruct is the same as TracKResponse
// but it is used where the <artist> tag contains only the
// artist's name.
type TrackResponseNoArtistStruct struct {
	Name       string                  `xml:"name"`
	ID         int64                   `xml:"id"`
	MBID       string                  `xml:"mbid"`
	URL        string                  `xml:"url"`
	Duration   int64                   `xml:"duration"`
	Streamable TrackStreamableResponse `xml:"streamable"`
	Listeners  int64                   `xml:"listeners"`
	Playcount  int64                   `xml:"playcount"`
	Artist     string                  `xml:"artist"`
	Image      []LastfmImageResponse   `xml:"image"`
	UserLoved  int                     `xml:"userloved"`
	Album      AlbumResponse           `xml:"album"`
	TopTags    TopTagsResponse         `xml:"toptags"`
	Wiki       LastfmWikiResponse      `xml:"wiki"`
}

// TrackInfoResponse is used for the track.getInfo's response.
type TrackInfoResponse struct {
	LastfmStatusResponse
	Track TrackResponse `xml:"track"`
}

// TopTracksResponse is used where <toptracks> tag is present.
type TopTracksResponse struct {
	LastfmStatusResponse
	TopTracks []TrackResponse `xml:"toptracks>track"`
}

// SimilarTracksResponse is used where <similartracks> tag is present.
type SimilarTracksResponse struct {
	LastfmStatusResponse
	SimilarTracks []TrackResponse `xml:"similartracks>track"`
}

// TrackSingleCorrectionResponse is part of TrackCorrectionResponse.
type TrackSingleCorrectionResponse struct {
	Index           int           `xml:"index,attr"`
	ArtistCorrected int           `xml:"artistcorrected,attr"`
	TrackCorrected  int           `xml:"trackcorrected,attr"`
	Track           TrackResponse `xml:"track"`
}

// TrackCorrectioNnResponse is used for the track.getCorrection's response.
type TrackCorrectionResponse struct {
	LastfmStatusResponse
	TrackCorrections []TrackSingleCorrectionResponse `xml:"corrections>correction"`
}

// TrackSearchResponse is used for the track.search's response.
type TrackSearchResponse struct {
	LastfmStatusResponse
	LastfmOpenSearchResponse
	TrackMatches []TrackResponseNoArtistStruct `xml:"results>trackmatches>track"`
}

// LovedTracksResponse is used where <lovedtracks> is present.
type LovedTracksResponse struct {
	LastfmStatusResponse
	LovedTracks []TrackResponse `xml:"lovedtracks>track"`
}

// RecentTrackResponse is part of RecentTracksResponse.
type RecentTrackResponse struct {
	TrackResponse
	NowPlaying string             `xml:"nowplaying,attr"`
	Date       LastfmDateResponse `xml:"date"`
}

// RecentTracksResponse is used where <recenttracks> is present.
type RecentTracksResponse struct {
	LastfmStatusResponse
	RecentTracks []RecentTrackResponse `xml:"recenttracks>track"`
}
