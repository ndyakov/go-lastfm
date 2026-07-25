package lastfm

import (
	"context"
	"net/url"
	"strings"
)

type AlbumClient struct{ client *Client }

type AlbumInfoParams struct {
	AlbumRef
	Username string
}

func (c AlbumClient) GetInfo(ctx context.Context, p AlbumInfoParams) (*AlbumInfoResponse, error) {
	v := url.Values{}
	p.AlbumRef.values(v)
	set(v, "username", p.Username)
	out := new(AlbumInfoResponse)
	return out, c.client.call(ctx, "album.getInfo", v, out)
}

type AlbumTagsParams struct {
	AlbumRef
	User string
}

func (c AlbumClient) GetTags(ctx context.Context, p AlbumTagsParams) (*TagsResponse, error) {
	v := url.Values{}
	p.AlbumRef.values(v)
	set(v, "user", p.User)
	out := new(TagsResponse)
	return out, c.client.call(ctx, "album.getTags", v, out)
}

func (c AlbumClient) GetTopTags(ctx context.Context, p AlbumRef) (*TopTagsResponse, error) {
	v := url.Values{}
	p.values(v)
	out := new(TopTagsResponse)
	return out, c.client.call(ctx, "album.getTopTags", v, out)
}

type AlbumSearchParams struct {
	Album string
	Pagination
}

func (c AlbumClient) Search(ctx context.Context, p AlbumSearchParams) (*AlbumSearchResponse, error) {
	v := url.Values{}
	set(v, "album", p.Album)
	p.Pagination.values(v)
	out := new(AlbumSearchResponse)
	return out, c.client.call(ctx, "album.search", v, out)
}

func (c AlbumClient) AddTags(ctx context.Context, artist, album string, tags []string) error {
	v := url.Values{
		"artist": {artist},
		"album":  {album},
		"tags":   {strings.Join(tags, ",")},
	}
	return c.client.call(ctx, "album.addTags", v, new(StatusResponse))
}

func (c AlbumClient) RemoveTag(ctx context.Context, artist, album, tag string) error {
	v := url.Values{
		"artist": {artist},
		"album":  {album},
		"tag":    {tag},
	}
	return c.client.call(ctx, "album.removeTag", v, new(StatusResponse))
}
