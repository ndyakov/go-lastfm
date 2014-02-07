package lastfm

// ArtistBioResponse part of ArtistResponse
// can be found in artist.getInfo`s response.
type ArtistBioResponse struct {
	LastfmWikiResponse
}

// ArtistResponse, used where <artist> tag is present.
type ArtistResponse struct {
	Name           string                `xml:"name"`
	MBID           string                `xml:"mbid"`
	URL            string                `xml:"url"`
	Image          []LastfmImageResponse `xml:"image"`
	Streamable     int                   `xml:"streamable"`
	Match          float64               `xml:"match"`
	Stats          LastfmStatsResponse   `xml:"stats"`
	SimilarArtists []ArtistResponse      `xml:"similar>artist"`
	Bio            ArtistBioResponse     `xml:"bio"`
}

// ArtistInfoResponse, used for artist.getInfo request.
type ArtistInfoResponse struct {
	LastfmStatusResponse
	Artist ArtistResponse `xml:"artist"`
}

type ArtistCorrectionResponse struct {
	LastfmStatusResponse
	Corrections []ArtistResponse `xml:"corrections>correction>artist"`
}

// ArtistSimilarResponse, used where <similarartists> is present.
type ArtistSimilarResponse struct {
	LastfmStatusResponse
	SimilarArtists []ArtistResponse `xml:"similarartists>artist"`
}

// ArtistSearchResponse, used for artist.Search request.
type ArtistSearchResponse struct {
	LastfmStatusResponse
	LastfmOpenSearchResponse
	ArtistMatches []ArtistResponse `xml:"results>artistmatches>artist"`
}

// TopArtistsResponse, used where <topartists> tag is present.
type TopArtistsResponse struct {
	LastfmStatusResponse
	TopArtists []ArtistResponse `xml:"topartists>artist"`
}
