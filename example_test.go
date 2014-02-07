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

// Get Top Tracks for Artist.
func ExampleArtistClient_GetTopTracks() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetTopTracks("Ogonek", "", 0, 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	track := response.TopTracks[0]
	fmt.Printf("Top Tracks By : %v\n", track.Artist.Name)
	fmt.Printf("Top Track #1 : %v\n", track.Name)
	fmt.Printf("URL : %v\n", track.URL)
	fmt.Printf("Playcount: %v\n", track.Playcount)

	// Output:
	// Top Tracks By : Ogonek
	// Top Track #1 : Starlight
	// URL : http://www.last.fm/music/Ogonek/_/Starlight
	// Playcount: 109
}

// Search for Artist.
func ExampleArtistClient_Search() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.Search("Ogo", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	var artist lastfm.ArtistResponse

	fmt.Printf("Search request for : %v\n", response.Query.SearchTerms)
	fmt.Printf("Total results : %v\n", response.TotalResults)
	if response.TotalResults > 0 {
		artist = response.ArtistMatches[0]
	}
	fmt.Printf("First result :\n")
	fmt.Printf("  Artist Name : %v\n", artist.Name)
	fmt.Printf("  Artist MBID : %v\n", artist.MBID)
	fmt.Printf("  Artist Image[%v] : %v\n", artist.Image[0].Size, artist.Image[0].URL)

	// Output:
	// Search request for : Ogo
	// Total results : 22
	// First result :
	//   Artist Name : Ogonek
	//   Artist MBID : 971762d6-c851-499c-a656-7fa4fa055f47
	//   Artist Image[small] : http://userserve-ak.last.fm/serve/34/26437217.jpg
}

// Get full info for Artist.
func ExampleArtistClient_GetInfo() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetInfo("Gorillaz", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("Artist : %v\n", response.Artist.Name)
	fmt.Printf("MBID : %v\n", response.Artist.MBID)
	fmt.Printf("URL : %v\n", response.Artist.URL)
	fmt.Printf("Listeners : %v\n", response.Artist.Stats.Listeners)
	fmt.Printf("Playcount : %v\n", response.Artist.Stats.Playcount)
	// Output:
	// Artist : Gorillaz
	// MBID : e21857d5-3256-4547-afb3-4b6ded592596
	// URL : http://www.last.fm/music/Gorillaz
	// Listeners : 3091710
	// Playcount : 124649955
}

// Get full info for Album.
func ExampleAlbumClient_GetInfo() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Album.GetInfo("Ogonek", "Drum And Bass Massacre", "", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("Artist : %v\n", response.Album.Artist)
	fmt.Printf("Album : %v\n", response.Album.Name)
	fmt.Printf("ID : %v\n", response.Album.ID)
	fmt.Printf("URL : %v\n", response.Album.URL)
	fmt.Printf("Listeners : %v\n", response.Album.Listeners)
	fmt.Printf("Playcount : %v\n", response.Album.Playcount)
	fmt.Printf("First track :\n")
	track := response.Album.Tracks[0]
	fmt.Printf("  Title : %v\n", track.Name)
	fmt.Printf("  Duration : %v\n", track.Duration)
	// Output:
	// Artist : Ogonek
	// Album : Drum And Bass Massacre
	// ID : 3668131
	// URL : http://www.last.fm/music/Ogonek/Drum+And+Bass+Massacre
	// Listeners : 2163
	// Playcount : 4288
	// First track :
	//   Title : luk i praz
	//   Duration : 302
}

// Get Tags for Album.
func ExampleAlbumClient_GetTags() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Album.GetTags("Guano Apes", "Bel Air", "", "n3mo-", 1)
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
	// Name : guano apes v2
	// URL  : http://www.last.fm/tag/guano%20apes%20v2
}

// Get Top Tags for Album.
func ExampleAlbumClient_GetTopTags() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Album.GetTopTags("Guano Apes", "Bel Air", "", 1)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("Number of Top Tags : %v\n", len(response.TopTags))
	fmt.Printf("Tag #1 :\n")
	fmt.Printf("Name : %v\n", response.TopTags[0].Name)
	fmt.Printf("URL  : %v\n", response.TopTags[0].URL)
	fmt.Printf("Count : %v\n", response.TopTags[0].Count)
	// Output:
	// Number of Top Tags : 14
	// Tag #1 :
	// Name : rock
	// URL  : http://www.last.fm/tag/rock
	// Count : 100
}
