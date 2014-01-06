package lastfm

type UserResponse struct {
	ID         int64                 `xml:"id"`
	Name       string                `xml:"name"`
	RealName   string                `xml:"realname"`
	URL        string                `xml:"url"`
	Image      []LastfmImageResponse `xml:"image"`
	Country    string                `xml:"country"`
	Age        int                   `xml:"age"`
	Gender     rune                  `xml:"gender"`
	Subscriber int                   `xml:"subscriber"`
	Playcount  int64                 `xml:"playcount"`
	Playlists  int                   `xml:"playlists"`
	Bootstrap  int                   `xml:"bootstrap"`
	Registered LastfmDateResponse    `xml:"registered"`
}

type UserInfoResponse struct {
	LastfmStatusResponse
	UserResponse
}

type FanResponse struct {
	UserResponse
	Weight int64 `xml:"weight"`
}

type TopFansResponse struct {
	LastfmStatusResponse
	TopFans []FanResponse `xml:"topfans>user"`
}
