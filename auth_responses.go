package lastfm

type AuthSessionResponse struct {
	LastfmStatusResponse
	Session LastfmSessionResponse `xml:"session"`
}

type AuthTokenResponse struct {
	LastfmStatusResponse
	Token string `xml:"token"`
}
