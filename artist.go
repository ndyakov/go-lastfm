package lastfm

import (
	"context"
	"net/url"
	"strings"
)

type ArtistClient struct{ client *Client }

func (c ArtistClient) GetCorrection(ctx context.Context, artist string) (*ArtistCorrectionResponse, error) {
	v := url.Values{"artist": {artist}}
	out := new(ArtistCorrectionResponse)
	return out, c.client.call(ctx, "artist.getCorrection", v, out)
}

type ArtistInfoParams struct {
	ArtistRef
	Username, Language string
}

func (c ArtistClient) GetInfo(ctx context.Context, p ArtistInfoParams) (*ArtistInfoResponse, error) {
	v := url.Values{}
	p.ArtistRef.values(v)
	set(v, "username", p.Username)
	set(v, "lang", p.Language)
	out := new(ArtistInfoResponse)
	return out, c.client.call(ctx, "artist.getInfo", v, out)
}

type ArtistSimilarParams struct {
	ArtistRef
	Limit int
}

func (c ArtistClient) GetSimilar(ctx context.Context, p ArtistSimilarParams) (*ArtistSimilarResponse, error) {
	v := url.Values{}
	p.ArtistRef.values(v)
	setInt(v, "limit", p.Limit)
	out := new(ArtistSimilarResponse)
	return out, c.client.call(ctx, "artist.getSimilar", v, out)
}

type ArtistTagsParams struct {
	ArtistRef
	User string
}

func (c ArtistClient) GetTags(ctx context.Context, p ArtistTagsParams) (*TagsResponse, error) {
	v := url.Values{}
	p.ArtistRef.values(v)
	set(v, "user", p.User)
	out := new(TagsResponse)
	return out, c.client.call(ctx, "artist.getTags", v, out)
}

type ArtistPagedParams struct {
	ArtistRef
	Pagination
}

func (c ArtistClient) GetTopAlbums(ctx context.Context, p ArtistPagedParams) (*TopAlbumsResponse, error) {
	v := url.Values{}
	p.ArtistRef.values(v)
	p.Pagination.values(v)
	out := new(TopAlbumsResponse)
	return out, c.client.call(ctx, "artist.getTopAlbums", v, out)
}

func (c ArtistClient) GetTopTracks(ctx context.Context, p ArtistPagedParams) (*TopTracksResponse, error) {
	v := url.Values{}
	p.ArtistRef.values(v)
	p.Pagination.values(v)
	out := new(TopTracksResponse)
	return out, c.client.call(ctx, "artist.getTopTracks", v, out)
}

func (c ArtistClient) GetTopTags(ctx context.Context, p ArtistRef) (*TopTagsResponse, error) {
	v := url.Values{}
	p.values(v)
	out := new(TopTagsResponse)
	return out, c.client.call(ctx, "artist.getTopTags", v, out)
}

type ArtistSearchParams struct {
	Artist string
	Pagination
}

func (c ArtistClient) Search(ctx context.Context, p ArtistSearchParams) (*ArtistSearchResponse, error) {
	v := url.Values{}
	set(v, "artist", p.Artist)
	p.Pagination.values(v)
	out := new(ArtistSearchResponse)
	return out, c.client.call(ctx, "artist.search", v, out)
}

func (c ArtistClient) AddTags(ctx context.Context, artist string, tags []string) error {
	v := url.Values{
		"artist": {artist},
		"tags":   {strings.Join(tags, ",")},
	}
	return c.client.call(ctx, "artist.addTags", v, new(StatusResponse))
}

func (c ArtistClient) RemoveTag(ctx context.Context, artist, tag string) error {
	v := url.Values{
		"artist": {artist},
		"tag":    {tag},
	}
	return c.client.call(ctx, "artist.removeTag", v, new(StatusResponse))
}
