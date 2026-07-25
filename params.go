package lastfm

import (
	"net/url"
	"strconv"
)

type Pagination struct{ Page, Limit int }

func (p Pagination) values(v url.Values) {
	setInt(v, "page", p.Page)
	setInt(v, "limit", p.Limit)
}

type ArtistRef struct {
	Artist, MBID string
	Autocorrect  bool
}

func (p ArtistRef) values(v url.Values) {
	set(v, "artist", p.Artist)
	set(v, "mbid", p.MBID)
	setBool(v, "autocorrect", p.Autocorrect)
}

type AlbumRef struct {
	Artist, Album, MBID string
	Autocorrect         bool
}

func (p AlbumRef) values(v url.Values) {
	set(v, "artist", p.Artist)
	set(v, "album", p.Album)
	set(v, "mbid", p.MBID)
	setBool(v, "autocorrect", p.Autocorrect)
}

type TrackRef struct {
	Artist, Track, MBID string
	Autocorrect         bool
}

func (p TrackRef) values(v url.Values) {
	set(v, "artist", p.Artist)
	set(v, "track", p.Track)
	set(v, "mbid", p.MBID)
	setBool(v, "autocorrect", p.Autocorrect)
}

func set(v url.Values, key, value string) {
	if value != "" {
		v.Set(key, value)
	}
}

func setInt(v url.Values, key string, value int) {
	if value != 0 {
		v.Set(key, strconv.Itoa(value))
	}
}

func setInt64(v url.Values, key string, value int64) {
	if value != 0 {
		v.Set(key, strconv.FormatInt(value, 10))
	}
}

func setBool(v url.Values, key string, value bool) {
	if value {
		v.Set(key, "1")
	}
}
