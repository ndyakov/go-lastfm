package lastfm

import (
	"context"
	"net/url"
)

type TagClient struct{ client *Client }

type TagInfoParams struct {
	Tag      string
	Language string
}

func (c TagClient) GetInfo(ctx context.Context, p TagInfoParams) (*TagInfoResponse, error) {
	v := url.Values{}
	set(v, "tag", p.Tag)
	set(v, "lang", p.Language)
	out := new(TagInfoResponse)
	return out, c.client.call(ctx, "tag.getInfo", v, out)
}

func (c TagClient) GetSimilar(ctx context.Context, tag string) (*TagSimilarResponse, error) {
	out := new(TagSimilarResponse)
	return out, c.client.call(ctx, "tag.getSimilar", url.Values{"tag": {tag}}, out)
}

type TagPagedParams struct {
	Tag string
	Pagination
}

func (c TagClient) GetTopAlbums(ctx context.Context, p TagPagedParams) (*TopAlbumsResponse, error) {
	v := url.Values{}
	set(v, "tag", p.Tag)
	p.Pagination.values(v)
	out := new(TopAlbumsResponse)
	return out, c.client.call(ctx, "tag.getTopAlbums", v, out)
}

func (c TagClient) GetTopArtists(ctx context.Context, p TagPagedParams) (*TopArtistsResponse, error) {
	v := url.Values{}
	set(v, "tag", p.Tag)
	p.Pagination.values(v)
	out := new(TopArtistsResponse)
	return out, c.client.call(ctx, "tag.getTopArtists", v, out)
}

func (c TagClient) GetTopTracks(ctx context.Context, p TagPagedParams) (*TopTracksResponse, error) {
	v := url.Values{}
	set(v, "tag", p.Tag)
	p.Pagination.values(v)
	out := new(TopTracksResponse)
	return out, c.client.call(ctx, "tag.getTopTracks", v, out)
}

func (c TagClient) GetTopTags(ctx context.Context) (*TopTagsResponse, error) {
	out := new(TopTagsResponse)
	return out, c.client.call(ctx, "tag.getTopTags", nil, out)
}

func (c TagClient) GetWeeklyChartList(ctx context.Context, tag string) (*WeeklyChartListResponse, error) {
	out := new(WeeklyChartListResponse)
	return out, c.client.call(ctx, "tag.getWeeklyChartList", url.Values{"tag": {tag}}, out)
}
