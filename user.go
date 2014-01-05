package lastfm

type User struct {
	ID int64 `xml:"id"`
	Name string `xml:"name"`
	RealName string `xml:"realname"`
	URL string `xml:"url"`
	Image string `xml:"image"`
	Country string `xml:"country"`
	Age int `xml:"age"`
	Gender rune `xml:"gender"`
	Subscriber int `xml:"subscriber"`
	Playcount int64 `xml:"playcount"`
	Playlists int `xml:"playlists"`
	Bootstrap int `xml:"bootstrap"`
	Registered LastfmDate `xml:"registered"`
}
