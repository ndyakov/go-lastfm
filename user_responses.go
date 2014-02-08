package lastfm

type UserResponse struct {
	ID         int64                 `xml:"id"`
	Name       string                `xml:"name"`
	RealName   string                `xml:"realname"`
	URL        string                `xml:"url"`
	Image      []LastfmImageResponse `xml:"image"`
	Country    string                `xml:"country"`
	Age        int                   `xml:"age"`
	Gender     string                `xml:"gender"`
	Subscriber int                   `xml:"subscriber"`
	Playcount  int64                 `xml:"playcount"`
	Playlists  int                   `xml:"playlists"`
	Bootstrap  int                   `xml:"bootstrap"`
	Registered LastfmDateResponse    `xml:"registered"`
}

type UserInfoResponse struct {
	LastfmStatusResponse
	User UserResponse `xml:"user"`
}

type FanResponse struct {
	UserResponse
	Weight int64 `xml:"weight"`
}

type TopFansResponse struct {
	LastfmStatusResponse
	TopFans []FanResponse `xml:"topfans>user"`
}

type NeighbourResponse struct {
	UserResponse
	Match float64 `xml:"match"`
}
type NeighboursResponse struct {
	LastfmStatusResponse
	Neighbours []NeighbourResponse `xml:"neighbours>user"`
}
