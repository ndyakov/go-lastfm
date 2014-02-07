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
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("Number of Top Tags : %v\n", len(response.TopTags))
	fmt.Printf("Tag #1 :\n")
	fmt.Printf("Name : %v\n", response.TopTags[0].Name)
	fmt.Printf("URL  : %v\n", response.TopTags[0].URL)
	// Output:
	// Number of Top Tags : 100
	// Tag #1 :
	// Name : electronic
	// URL  : http://www.last.fm/tag/electronic
}

// Get Tags for Artist.
func ExampleArtistClient_GetTags() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetTags("Red Hot Chili Peppers", "", 1, "RJ")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("Number of Tags by User : %v\n", len(response.Tags))
	fmt.Printf("Tag #1 :\n")
	fmt.Printf("Name : %v\n", response.Tags[0].Name)
	fmt.Printf("URL  : %v\n", response.Tags[0].URL)
	// Output:
	// Number of Tags by User : 2
	// Tag #1 :
	// Name : funky
	// URL  : http://www.last.fm/tag/funky
}

// Get Similar Artists.
func ExampleArtistClient_GetSimilar() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetSimilar("Cher", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	artist := response.SimilarArtists[0]
	fmt.Printf("Number of matches : %v\n", len(response.SimilarArtists))
	fmt.Printf("First match :\n")
	fmt.Printf("Artist : %v\n", artist.Name)
	fmt.Printf("MBID   : %v\n", artist.MBID)
	fmt.Printf("Match  : %v\n", artist.Match)
	// Output:
	// Number of matches : 100
	// First match :
	// Artist : Sonny & Cher
	// MBID   : 3d6e4b6d-2700-458c-9722-9021965a8164
	// Match  : 1
}

// Get Top Albums for Artist.
func ExampleArtistClient_GetTopAlbums() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetTopAlbums("Ogonek", "", 0, 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	album := response.TopAlbums[0]
	fmt.Printf("Top Albums By : %v\n", album.Artist.Name)
	fmt.Printf("Top Album #1 : %v\n", album.Name)
	fmt.Printf("URL : %v\n", album.URL)
	fmt.Printf("Playcount: %v\n", album.Playcount)
	fmt.Printf("Image [%v] : %v\n", album.Image[0].Size, album.Image[0].URL)

	// Output:
	// Top Albums By : Ogonek
	// Top Album #1 : Drum And Bass Massacre
	// URL : http://www.last.fm/music/Ogonek/Drum+And+Bass+Massacre
	// Playcount: 26
	// Image [small] : http://cdn.last.fm/flatness/catalogue/noimage/2/default_album_medium.png
}

// Get Top Fans for Artist.
func ExampleArtistClient_GetTopFans() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetTopFans("Ogonek", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	user := response.TopFans[0]
	fmt.Printf("Top Fan #1 : %v\n", user.Name)
	fmt.Printf("URL : %v\n", user.URL)
	fmt.Printf("Image [%v] : %v\n", user.Image[0].Size, user.Image[0].URL)

	// Output:
	// Top Fan #1 : rikardo_83
	// URL : http://www.last.fm/user/rikardo_83
	// Image [small] : http://userserve-ak.last.fm/serve/34/79926153.gif
}
