package lastfm

import (
	"context"
	"net/url"
)

type LibraryClient struct{ client *Client }
type LibraryArtistsParams struct {
	User string
	Pagination
}

func (c LibraryClient) GetArtists(ctx context.Context, p LibraryArtistsParams) (*LibraryArtistsResponse, error) {
	v := url.Values{}
	set(v, "user", p.User)
	p.Pagination.values(v)
	out := new(LibraryArtistsResponse)
	return out, c.client.call(ctx, "library.getArtists", v, out)
}
