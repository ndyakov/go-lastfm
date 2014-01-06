package lastfm

type ArtistBioResponse struct {
	LastfmWikiResponse
}

type ArtistResponse struct {
	Name       string                `xml:"name"`
	MBID       string                `xml:"mbid"`
	URL        string                `xml:"url"`
	Image      []LastfmImageResponse `xml:"image"`
	Streamable int                   `xml:"streamable"`
	Match      float32               `xml:"match"`
}

type ArtistInfoResponse struct {
	LastfmStatusResponse
	ArtistResponse
	Stats          LastfmStatsResponse `xml:"stats"`
	SimilarArtists []ArtistResponse    `xml:"similar>artist"`
	Bio            ArtistBioResponse   `xml:"bio"`
}

type ArtistCorrectionResponse struct {
	LastfmStatusResponse
	Corrections []ArtistResponse `xml:"corrections>correction>artist"`
}

type ArtistSimilarResponse struct {
	LastfmStatusResponse
	SimilarArtists []ArtistResponse `xml:"similarartists>artist"`
}

type ArtistSearchResponse struct {
	LastfmStatusResponse
	LastfmOpenSearchResponse
	ArtistMatches []ArtistResponse `xml:"artistmaches>artist"`
}
