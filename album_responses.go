package lastfm

type AlbumResponse struct {
	Name        string                `xml:"name"`
	ArtistName  string                `xml:"artist"`
	ID          int64                 `xml:"id"`
	MBID        string                `xml:"mbid"`
	ReleaseDate string                `xml:"releasedate"`
	Images      []LastfmImageResponse `xml:"image"`
	Listeners   int64                 `xml:"listeners"`
	Playcount   int64                 `xml:"playcount"`
	TopTags     []TagResponse         `xml:"toptags"`
	Tracks      []TrackResponse       `xml:"tracks"`
}
