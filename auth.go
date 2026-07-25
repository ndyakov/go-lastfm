package lastfm

import (
	"context"
	"net/url"
)

type AuthClient struct{ client *Client }

func (c AuthClient) GetToken(ctx context.Context) (*AuthTokenResponse, error) {
	out := new(AuthTokenResponse)
	return out, c.client.call(ctx, "auth.getToken", nil, out)
}

func (c AuthClient) GetSession(ctx context.Context, token string) (*AuthSessionResponse, error) {
	out := new(AuthSessionResponse)
	err := c.client.call(ctx, "auth.getSession", url.Values{"token": {token}}, out)
	if err == nil {
		c.client.SetSessionKey(out.Session.Key)
	}
	return out, err
}

func (c AuthClient) GetMobileSession(ctx context.Context, username, password string) (*AuthSessionResponse, error) {
	v := url.Values{"username": {username}, "password": {password}}
	out := new(AuthSessionResponse)
	err := c.client.call(ctx, "auth.getMobileSession", v, out)
	if err == nil {
		c.client.SetSessionKey(out.Session.Key)
	}
	return out, err
}
