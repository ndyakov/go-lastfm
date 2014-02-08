package lastfm

// UserResponse
// Used where <user> tag is present.
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

// UserInfoResponse
// Used for the user.getInfo's response.
type UserInfoResponse struct {
	LastfmStatusResponse
	User UserResponse `xml:"user"`
}

// FanResponse
// Part of TopFansResponse.
type FanResponse struct {
	UserResponse
	Weight int64 `xml:"weight"`
}

// TopFansResponse
// Used where <topfans> tag is present.
type TopFansResponse struct {
	LastfmStatusResponse
	TopFans []FanResponse `xml:"topfans>user"`
}

// NeighbourResponse
// Part of NeighboursResponse.
type NeighbourResponse struct {
	UserResponse
	Match float64 `xml:"match"`
}

// NeighboursResponse
// Used where <neighbours> tag is present.
type NeighboursResponse struct {
	LastfmStatusResponse
	Neighbours []NeighbourResponse `xml:"neighbours>user"`
}
