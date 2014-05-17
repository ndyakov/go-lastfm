package lastfm

type AuthClient struct {
	Client
}

func (c *AuthClient) GetSession() (response *AuthSessionResponse, err error) {
	response = new(AuthSessionResponse)
	query := make(map[string]string)
	query["method"] = "auth.getSession"
	tokenResponse, err := c.GetToken()

	if err == nil {
		query["token"] = tokenResponse.Token
	} else {
		return
	}

	err = c.lfm.getResponse(query, response)

	return
}

func (c *AuthClient) GetMobileSession(password, username string) (response *AuthSessionResponse, err error) {
	response = new(AuthSessionResponse)
	query := make(map[string]string)
	query["method"] = "auth.getMobileSession"
	query["username"] = username
	query["password"] = password
	err = c.lfm.getResponse(query, response)

	return
}

func (c *AuthClient) GetToken() (response *AuthTokenResponse, err error) {
	response = new(AuthTokenResponse)
	query := make(map[string]string)
	query["method"] = "auth.getToken"
	err = c.lfm.getResponse(query, response)

	if err == nil {
		c.lfm.SetToken(response.Token)
	}

	return
}
