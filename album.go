package lastfm

type Album struct {
	Name        string        `xml:"name"`
	ArtistName  string        `xml:"artist"`
	ID          int64         `xml:"id"`
	MBID        string        `xml:"mbid"`
	ReleaseDate string        `xml:"releasedate"`
	Images      []LastfmImage `xml:"image"`
	Listeners   int64         `xml:"listeners"`
	Playcount   int64         `xml:"playcount"`
	TopTags     []Tag         `xml:"toptags"`
	Tracks      []Track       `xml:"tracks"`
}
