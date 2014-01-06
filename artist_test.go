package lastfm

import (
	"testing"
)

func TestArtistTopTags(t *testing.T) {
	lfm := New("api_key_for_testing")
	response, err := lfm.Artist.GetTopTags("Daft Punk", "", 0)
	if err != nil {
		t.Error(err)
	}
	if len(response.TopTags) != 100 {
		t.Error("Expexted 100 top tags")
	}
}
