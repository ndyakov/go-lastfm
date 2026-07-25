package lastfm

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestExperimentalLiveContracts(t *testing.T) {
	if os.Getenv("LASTFM_RUN_EXPERIMENTAL_TESTS") != "1" {
		t.Skip("set LASTFM_RUN_EXPERIMENTAL_TESTS=1 to probe legacy methods")
	}
	apiKey := os.Getenv("LASTFM_API_KEY")
	if apiKey == "" {
		t.Fatal("LASTFM_API_KEY is required")
	}
	client := MustNew(apiKey, os.Getenv("LASTFM_API_SECRET"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("tag.search", func(t *testing.T) {
		_, err := client.Experimental.SearchTags(ctx, TagSearchParams{Tag: "rock"})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("geo.getMetros", func(t *testing.T) {
		_, err := client.Experimental.GetMetros(ctx, "Spain")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("user.getNeighbours", func(t *testing.T) {
		_, err := client.Experimental.GetNeighbours(ctx, NeighboursParams{User: "RJ", Limit: 1})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("artist.getTopFans", func(t *testing.T) {
		_, err := client.Experimental.GetArtistTopFans(ctx, TopFansParams{Artist: "Cher"})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("track.getTopFans", func(t *testing.T) {
		_, err := client.Experimental.GetTrackTopFans(ctx, TopFansParams{Artist: "Cher", Track: "Believe"})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("tasteometer.compare", func(t *testing.T) {
		_, err := client.Experimental.CompareTaste(ctx, TasteometerParams{
			Type1: "user", Value1: "RJ", Type2: "user", Value2: "lastfm",
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}
