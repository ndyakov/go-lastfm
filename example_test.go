package lastfm_test

import (
	"fmt"
	"github.com/ndyakov/go-lastfm"
	"reflect"
)
// Create new api client object
// with your api key as argument.
func ExampleLastFM_New() {
	lfm := lastfm.New("api_key_for_testing")
	fmt.Printf("The type is %v\n", reflect.TypeOf(lfm))
	// Output:
	// The type is *lastfm.LastFM
}

// Get Top Tags for Artist.
func ExampleArtistClient_GetTopTags() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetTopTags("Daft Punk", "", 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of Top Tags : %v\n", len(response.TopTags))
	fmt.Printf("Tag #1 :\n")
	fmt.Printf("Name : %v\n",response.TopTags[0].Name)
	fmt.Printf("URL  : %v\n", response.TopTags[0].URL)
	// Output:
	// Number of Top Tags : 100
	// Tag #1 :
	// Name : electronic
	// URL  : http://www.last.fm/tag/electronic
}

// Get Similar Artists.
func ExampleArtistClient_GetSimilar() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetSimilar("Cher","",0)
	if err != nil {
		fmt.Println(err)
	}
	artist := response.SimilarArtists[0]
	fmt.Printf("Number of matches : %v\n",len(response.SimilarArtists))
	fmt.Printf("First match :\n")
	fmt.Printf("Artist : %v\n",artist.Name)
	fmt.Printf("MBID   : %v\n",artist.MBID)
	fmt.Printf("Match  : %v\n",artist.Match)
	// Output:
	// Number of matches : 100
	// First match :
	// Artist : Sonny & Cher
	// MBID   : 3d6e4b6d-2700-458c-9722-9021965a8164
	// Match  : 1
}

