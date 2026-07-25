package lastfm

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type TrackClient struct{ client *Client }

func (c TrackClient) GetCorrection(ctx context.Context, artist, track string) (*TrackCorrectionResponse, error) {
	v := url.Values{
		"artist": {artist},
		"track":  {track},
	}
	out := new(TrackCorrectionResponse)
	return out, c.client.call(ctx, "track.getCorrection", v, out)
}

type TrackInfoParams struct {
	TrackRef
	Username string
}

func (c TrackClient) GetInfo(ctx context.Context, p TrackInfoParams) (*TrackInfoResponse, error) {
	v := url.Values{}
	p.TrackRef.values(v)
	set(v, "username", p.Username)
	out := new(TrackInfoResponse)
	return out, c.client.call(ctx, "track.getInfo", v, out)
}

type TrackSimilarParams struct {
	TrackRef
	Limit int
}

func (c TrackClient) GetSimilar(ctx context.Context, p TrackSimilarParams) (*SimilarTracksResponse, error) {
	v := url.Values{}
	p.TrackRef.values(v)
	setInt(v, "limit", p.Limit)
	out := new(SimilarTracksResponse)
	return out, c.client.call(ctx, "track.getSimilar", v, out)
}

type TrackTagsParams struct {
	TrackRef
	User string
}

func (c TrackClient) GetTags(ctx context.Context, p TrackTagsParams) (*TagsResponse, error) {
	v := url.Values{}
	p.TrackRef.values(v)
	set(v, "user", p.User)
	out := new(TagsResponse)
	return out, c.client.call(ctx, "track.getTags", v, out)
}

func (c TrackClient) GetTopTags(ctx context.Context, p TrackRef) (*TopTagsResponse, error) {
	v := url.Values{}
	p.values(v)
	out := new(TopTagsResponse)
	return out, c.client.call(ctx, "track.getTopTags", v, out)
}

type TrackSearchParams struct {
	Track, Artist string
	Pagination
}

func (c TrackClient) Search(ctx context.Context, p TrackSearchParams) (*TrackSearchResponse, error) {
	v := url.Values{}
	set(v, "track", p.Track)
	set(v, "artist", p.Artist)
	p.Pagination.values(v)
	out := new(TrackSearchResponse)
	return out, c.client.call(ctx, "track.search", v, out)
}

func (c TrackClient) AddTags(ctx context.Context, artist, track string, tags []string) error {
	v := url.Values{
		"artist": {artist},
		"track":  {track},
		"tags":   {strings.Join(tags, ",")},
	}
	return c.client.call(ctx, "track.addTags", v, new(StatusResponse))
}

func (c TrackClient) RemoveTag(ctx context.Context, artist, track, tag string) error {
	v := url.Values{
		"artist": {artist},
		"track":  {track},
		"tag":    {tag},
	}
	return c.client.call(ctx, "track.removeTag", v, new(StatusResponse))
}

func (c TrackClient) Love(ctx context.Context, artist, track string) error {
	v := url.Values{
		"artist": {artist},
		"track":  {track},
	}
	return c.client.call(ctx, "track.love", v, new(StatusResponse))
}

func (c TrackClient) Unlove(ctx context.Context, artist, track string) error {
	v := url.Values{
		"artist": {artist},
		"track":  {track},
	}
	return c.client.call(ctx, "track.unlove", v, new(StatusResponse))
}

type NowPlayingParams struct {
	Artist, Track, Album, AlbumArtist, MBID string
	TrackNumber, Duration                   int
}

func (c TrackClient) UpdateNowPlaying(ctx context.Context, p NowPlayingParams) (*TrackNowPlayingResponse, error) {
	v := url.Values{}
	set(v, "artist", p.Artist)
	set(v, "track", p.Track)
	set(v, "album", p.Album)
	set(v, "albumArtist", p.AlbumArtist)
	set(v, "mbid", p.MBID)
	setInt(v, "trackNumber", p.TrackNumber)
	setInt(v, "duration", p.Duration)
	out := new(TrackNowPlayingResponse)
	return out, c.client.call(ctx, "track.updateNowPlaying", v, out)
}

type Scrobble struct {
	Artist, Track, Album, AlbumArtist, MBID string
	Timestamp                               time.Time
	ChosenByUser                            *bool
	TrackNumber, Duration                   int
}

func (c TrackClient) Scrobble(ctx context.Context, entries []Scrobble) (*TrackScrobbleResponse, error) {
	if len(entries) == 0 || len(entries) > 50 {
		return nil, fmt.Errorf("lastfm: scrobble batch must contain 1 to 50 tracks")
	}
	v := url.Values{}
	for i, e := range entries {
		suffix := "[" + strconv.Itoa(i) + "]"
		v.Set("artist"+suffix, e.Artist)
		v.Set("track"+suffix, e.Track)
		v.Set("timestamp"+suffix, strconv.FormatInt(e.Timestamp.Unix(), 10))
		set(v, "album"+suffix, e.Album)
		set(v, "albumArtist"+suffix, e.AlbumArtist)
		set(v, "mbid"+suffix, e.MBID)
		setInt(v, "trackNumber"+suffix, e.TrackNumber)
		setInt(v, "duration"+suffix, e.Duration)
		if e.ChosenByUser != nil {
			if *e.ChosenByUser {
				v.Set("chosenByUser"+suffix, "1")
			} else {
				v.Set("chosenByUser"+suffix, "0")
			}
		}
	}
	out := new(TrackScrobbleResponse)
	return out, c.client.call(ctx, "track.scrobble", v, out)
}
