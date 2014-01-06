package lastfm

import (
	"math"
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

func TestArtistSimilar(t *testing.T) {
	lfm := New("api_key_for_testing")
	response, err := lfm.Artist.GetSimilar("Cher", "", 0)
	if err != nil {
		t.Error(err)
	}
	if math.Abs((float64)(response.SimilarArtists[0].Match-1)) > 0.01 {
		t.Errorf("Expexted SimilarArtists[0].Match to be closer to 1, but it is %v",
			response.SimilarArtists[0].Match)
	}
	if response.SimilarArtists[0].MBID != "3d6e4b6d-2700-458c-9722-9021965a8164" {
		t.Errorf("Expexted SimilarArtists[0] to be with MBID 3d6e4b6d-2700-458c-9722-9021965a8164, but got %v",
			response.SimilarArtists[0].MBID)
	}
}
