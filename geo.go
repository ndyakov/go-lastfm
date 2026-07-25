package lastfm

import (
	"context"
	"net/url"
)

type GeoClient struct{ client *Client }
type GeoParams struct {
	Country, Location string
	Pagination
}

func (c GeoClient) GetTopArtists(ctx context.Context, p GeoParams) (*TopArtistsResponse, error) {
	v := url.Values{}
	set(v, "country", p.Country)
	p.Pagination.values(v)
	out := new(TopArtistsResponse)
	return out, c.client.call(ctx, "geo.getTopArtists", v, out)
}

func (c GeoClient) GetTopTracks(ctx context.Context, p GeoParams) (*TopTracksResponse, error) {
	v := url.Values{}
	set(v, "country", p.Country)
	set(v, "location", p.Location)
	p.Pagination.values(v)
	out := new(TopTracksResponse)
	return out, c.client.call(ctx, "geo.getTopTracks", v, out)
}
