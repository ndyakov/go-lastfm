package lastfm

import (
	"encoding/json"
	"strings"
	"testing"
)

// The fixtures under testing/ are XML from the v1 API, so the JSON response
// structs had no coverage against real Last.fm JSON. These tests pin the
// quirks of that JSON: keys that change shape between endpoints, "#text"
// fields that arrive as bare numbers, and names carried in "#text".

func TestDateAcceptsNumericText(t *testing.T) {
	// user.getInfo returns registered."#text" as a bare JSON number, while
	// every other endpoint quotes it.
	for name, input := range map[string]string{
		"number": `{"unixtime":"1354320000","#text":1354320000}`,
		"string": `{"uts":"1354320000","#text":"12 Dec 2012, 00:00"}`,
	} {
		t.Run(name, func(t *testing.T) {
			var date Date
			if err := json.Unmarshal([]byte(input), &date); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if date.UTS != 1354320000 {
				t.Errorf("UTS = %d, want 1354320000", date.UTS)
			}
		})
	}
}

func TestDateAcceptsUnixtimeKey(t *testing.T) {
	// user.getInfo spells the registration timestamp "unixtime"; the rest of
	// the API uses "uts". Without both, registration time decodes to zero.
	var date Date
	if err := json.Unmarshal([]byte(`{"unixtime":"1354320000","#text":1354320000}`), &date); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if date.UTS != 1354320000 {
		t.Fatalf("UTS = %d, want 1354320000 (unixtime key ignored)", date.UTS)
	}
}

func TestStreamableAcceptsBothForms(t *testing.T) {
	// track.getInfo returns an object; user.getRecentTracks returns "0".
	tests := map[string]struct {
		input string
		want  Integer
	}{
		"object": {`{"fulltrack":"0","#text":"1"}`, 1},
		"string": {`"0"`, 0},
		"number": {`1`, 1},
		"null":   {`null`, 0},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var s Streamable
			if err := json.Unmarshal([]byte(tc.input), &s); err != nil {
				t.Fatalf("decode %s: %v", tc.input, err)
			}
			if s.Text != tc.want {
				t.Errorf("Text = %d, want %d", s.Text, tc.want)
			}
		})
	}
}

func TestListReportsElementErrors(t *testing.T) {
	// A decode failure inside an element must surface as itself. Falling back
	// to the single-object form on any array error reports the misleading
	// "cannot unmarshal array into Go value of type T" instead.
	var list List[Track]
	err := json.Unmarshal([]byte(`[{"duration":{"bad":"shape"}}]`), &list)
	if err == nil {
		t.Fatal("want error for malformed element, got nil")
	}
	if strings.Contains(err.Error(), "cannot unmarshal array") {
		t.Errorf("element error masked by single-object fallback: %v", err)
	}
}

func TestListAcceptsArrayAndSingleObject(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"array":  {`[{"name":"a"},{"name":"b"}]`, 2},
		"single": {`{"name":"a"}`, 1},
		"null":   {`null`, 0},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var list List[Tag]
			if err := json.Unmarshal([]byte(tc.input), &list); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if len(list) != tc.want {
				t.Errorf("len = %d, want %d", len(list), tc.want)
			}
		})
	}
}

func TestArtistAndAlbumNameFromText(t *testing.T) {
	// Nested in user.getRecentTracks entries, artist and album carry their
	// name in "#text" rather than "name".
	var artist Artist
	if err := json.Unmarshal([]byte(`{"mbid":"abc","#text":"Magdalena Bay"}`), &artist); err != nil {
		t.Fatalf("decode artist: %v", err)
	}
	if artist.Name != "Magdalena Bay" {
		t.Errorf("Artist.Name = %q, want %q", artist.Name, "Magdalena Bay")
	}

	var album Album
	if err := json.Unmarshal([]byte(`{"mbid":"def","#text":"Imaginal Disk"}`), &album); err != nil {
		t.Fatalf("decode album: %v", err)
	}
	if album.Name != "Imaginal Disk" {
		t.Errorf("Album.Name = %q, want %q", album.Name, "Imaginal Disk")
	}

	// The canonical "name" key must still win where both could appear.
	var named Artist
	if err := json.Unmarshal([]byte(`{"name":"Clairo","#text":"ignored"}`), &named); err != nil {
		t.Fatalf("decode artist: %v", err)
	}
	if named.Name != "Clairo" {
		t.Errorf("Artist.Name = %q, want %q", named.Name, "Clairo")
	}
}

// The decoders are reached through encoding/json in normal use, which only
// hands them syntactically valid JSON. Call them directly to reach the error
// paths for malformed or wrongly-typed input.
func TestUnmarshalErrorPaths(t *testing.T) {
	tests := map[string]struct {
		target json.Unmarshaler
		input  string
	}{
		"string/unterminated":   {new(String), `"unterminated`},
		"date/bad uts":          {new(Date), `{"uts":{}}`},
		"date/malformed":        {new(Date), `{`},
		"artist/name not text":  {new(Artist), `{"name":123}`},
		"album/name not text":   {new(Album), `{"name":123}`},
		"streamable/bad object": {new(Streamable), `{"fulltrack":"abc"}`},
		"streamable/bad scalar": {new(Streamable), `"abc"`},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if err := tc.target.UnmarshalJSON([]byte(tc.input)); err == nil {
				t.Errorf("UnmarshalJSON(%s) = nil, want error", tc.input)
			}
		})
	}
}

func TestStringAcceptsEmptyAndNull(t *testing.T) {
	for name, input := range map[string]string{
		"null":  `null`,
		"empty": ``,
	} {
		t.Run(name, func(t *testing.T) {
			s := String("stale")
			if err := s.UnmarshalJSON([]byte(input)); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if s != "" {
				t.Errorf("String = %q, want empty", s)
			}
		})
	}
}

const userInfoPayload = `{"user":{
  "name":"rj","age":"0","subscriber":"1","realname":"","bootstrap":"0",
  "playcount":"8797","artist_count":"1103","playlists":"0",
  "image":[{"size":"small","#text":"https://example.invalid/s.png"}],
  "registered":{"unixtime":"1354320000","#text":1354320000},
  "country":"None","gender":"n","url":"https://www.last.fm/user/rj","type":"user"}}`

func TestUserInfoResponse(t *testing.T) {
	var resp UserInfoResponse
	if err := json.Unmarshal([]byte(userInfoPayload), &resp); err != nil {
		t.Fatalf("decode user.getInfo: %v", err)
	}
	if resp.User.Name != "rj" {
		t.Errorf("Name = %q, want %q", resp.User.Name, "rj")
	}
	if resp.User.Playcount != 8797 {
		t.Errorf("Playcount = %d, want 8797", resp.User.Playcount)
	}
	if resp.User.Registered.UTS != 1354320000 {
		t.Errorf("Registered.UTS = %d, want 1354320000", resp.User.Registered.UTS)
	}
}

const recentTracksPayload = `{"recenttracks":{"track":[
  {"artist":{"mbid":"68c261d5","#text":"Magdalena Bay"},
   "streamable":"0",
   "image":[{"size":"small","#text":"https://example.invalid/s.png"}],
   "mbid":"","album":{"mbid":"","#text":"Imaginal Disk"},
   "name":"Death & Romance","url":"https://example.invalid/track",
   "date":{"uts":"1754900000","#text":"11 Aug 2025, 12:00"}}],
  "@attr":{"user":"rj","totalPages":"1100","page":"1","perPage":"1","total":"8797"}}}`

func TestRecentTracksResponse(t *testing.T) {
	var resp RecentTracksResponse
	if err := json.Unmarshal([]byte(recentTracksPayload), &resp); err != nil {
		t.Fatalf("decode user.getRecentTracks: %v", err)
	}
	if len(resp.RecentTracks.Tracks) != 1 {
		t.Fatalf("len(Tracks) = %d, want 1", len(resp.RecentTracks.Tracks))
	}
	track := resp.RecentTracks.Tracks[0]
	if track.Name != "Death & Romance" {
		t.Errorf("Name = %q, want %q", track.Name, "Death & Romance")
	}
	if track.Artist.Name != "Magdalena Bay" {
		t.Errorf("Artist.Name = %q, want %q", track.Artist.Name, "Magdalena Bay")
	}
	if track.Album.Name != "Imaginal Disk" {
		t.Errorf("Album.Name = %q, want %q", track.Album.Name, "Imaginal Disk")
	}
	if track.Date.UTS != 1754900000 {
		t.Errorf("Date.UTS = %d, want 1754900000", track.Date.UTS)
	}
	if resp.RecentTracks.Attr.Total != 8797 {
		t.Errorf("Attr.Total = %d, want 8797", resp.RecentTracks.Attr.Total)
	}
}
