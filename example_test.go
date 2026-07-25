package lastfm_test

import (
	"context"
	"fmt"
	"log"

	lastfm "github.com/ndyakov/go-lastfm/v2"
)

func ExampleClient() {
	client, err := lastfm.New("api-key", "shared-secret", lastfm.WithUserAgent("my-app/1.0"))
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.Artist.GetInfo(context.Background(), lastfm.ArtistInfoParams{
		ArtistRef: lastfm.ArtistRef{Artist: "Björk", Autocorrect: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Artist.Name)
}
