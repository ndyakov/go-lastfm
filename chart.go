package lastfm

import (
	"context"
	"net/url"
)

type ChartClient struct{ client *Client }

func (c ChartClient) GetTopArtists(ctx context.Context, p Pagination) (*ChartArtistsResponse, error) {
	v := url.Values{}
	p.values(v)
	out := new(ChartArtistsResponse)
	return out, c.client.call(ctx, "chart.getTopArtists", v, out)
}

func (c ChartClient) GetTopTags(ctx context.Context, p Pagination) (*ChartTagsResponse, error) {
	v := url.Values{}
	p.values(v)
	out := new(ChartTagsResponse)
	return out, c.client.call(ctx, "chart.getTopTags", v, out)
}

func (c ChartClient) GetTopTracks(ctx context.Context, p Pagination) (*ChartTracksResponse, error) {
	v := url.Values{}
	p.values(v)
	out := new(ChartTracksResponse)
	return out, c.client.call(ctx, "chart.getTopTracks", v, out)
}
