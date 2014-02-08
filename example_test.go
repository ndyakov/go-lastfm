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
	response, err := lfm.Artist.GetTags("Red Hot Chili Peppers", "", "RJ", 1)
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

// Search for Album.
func ExampleAlbumClient_Search() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Album.Search("go", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	var album lastfm.AlbumResponseNoArtistStruct

	fmt.Printf("Search request for : %v\n", response.Query.SearchTerms)
	fmt.Printf("Total results : %v\n", response.TotalResults)
	if response.TotalResults > 0 {
		album = response.AlbumMatches[0]
	}
	fmt.Printf("First result :\n")
	fmt.Printf("  Album Name : %v\n", album.Name)
	fmt.Printf("  Album MBID : %v\n", album.MBID)
	fmt.Printf("  Artist Name : %v\n", album.Artist)
	fmt.Printf("  Album Image[%v] : %v\n", album.Image[0].Size, album.Image[0].URL)

	// Output:
	// Search request for : go
	// Total results : 31102
	// First result :
	//   Album Name : Demon Days
	//   Album MBID : 26193df8-54c5-4c75-80e3-84ddc9aa7379
	//   Artist Name : Gorillaz
	//   Album Image[small] : http://userserve-ak.last.fm/serve/34s/44425129.png
}

// Get full info for Tag.
func ExampleTagClient_GetInfo() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Tag.GetInfo("rock")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("Name : %v\n", response.Tag.Name)
	fmt.Printf("URL : %v\n", response.Tag.URL)
	fmt.Printf("Reach : %v\n", response.Tag.Reach)
	fmt.Printf("Taggings : %v\n", response.Tag.Taggings)
	// Output:
	// Name : rock
	// URL : http://www.last.fm/tag/rock
	// Reach : 374178
	// Taggings : 3979953
}

// Get Similar Tags.
func ExampleTagClient_GetSimilar() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Tag.GetSimilar("rock")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	tag := response.SimilarTags[0]
	fmt.Printf("Number of matches : %v\n", len(response.SimilarTags))
	fmt.Printf("First match :\n")
	fmt.Printf("Tag : %v\n", tag.Name)
	fmt.Printf("URL : %v\n", tag.URL)
	// Output:
	// Number of matches : 50
	// First match :
	// Tag : classic rock
	// URL : http://www.last.fm/tag/classic%20rock
}

// Get Top Albums for Tag.
func ExampleTagClient_GetTopAlbums() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Tag.GetTopAlbums("dnb", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	album := response.TopAlbums[0]
	fmt.Printf("Top Album #1 : %v\n", album.Name)
	fmt.Printf("Artist : %v\n", album.Artist.Name)
	fmt.Printf("URL : %v\n", album.URL)
	fmt.Printf("MBID : %v\n", album.MBID)
	fmt.Printf("Image [%v] : %v\n", album.Image[0].Size, album.Image[0].URL)

	// Output:
	// Top Album #1 : Hold Your Colour
	// Artist : Pendulum
	// URL : http://www.last.fm/music/Pendulum/Hold+Your+Colour
	// MBID : 9d9b873c-fbd4-43df-9533-b401dd86081d
	// Image [small] : http://userserve-ak.last.fm/serve/34s/41919803.png
}

// Get Top Artists for Tag.
func ExampleTagClient_GetTopArtists() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Tag.GetTopArtists("dnb", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	artist := response.TopArtists[0]
	fmt.Printf("Artist #1 : %v\n", artist.Name)
	fmt.Printf("URL : %v\n", artist.URL)
	fmt.Printf("MBID : %v\n", artist.MBID)

	// Output:
	// Artist #1 : Pendulum
	// URL : http://www.last.fm/music/Pendulum
	// MBID : 2030e776-73b2-4cf8-8c15-813e801f8151
}

// Get Top Tags in Lastfm.
func ExampleTagClient_GetTopTags() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Tag.GetTopTags()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("Number of Top Tags : %v\n", len(response.TopTags))
	fmt.Printf("Tag #1 :\n")
	fmt.Printf("Name : %v\n", response.TopTags[0].Name)
	fmt.Printf("URL : %v\n", response.TopTags[0].URL)
	fmt.Printf("Count : %v\n", response.TopTags[0].Count)
	// Output:
	// Number of Top Tags : 250
	// Tag #1 :
	// Name : rock
	// URL : www.last.fm/tag/rock
	// Count : 3979953
}

// Get Top Tracks for Tag.
func ExampleTagClient_GetTopTracks() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Tag.GetTopTracks("rock", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	track := response.TopTracks[0]
	fmt.Printf("Top Track #1 : %v\n", track.Name)
	fmt.Printf("MBID : %v\n", track.MBID)
	fmt.Printf("Artist : %v\n", track.Artist.Name)
	fmt.Printf("URL : %v\n", track.URL)

	// Output:
	// Top Track #1 : Counting Stars
	// MBID : 5e8d518e-328a-43da-a9fa-cc27513a4c86
	// Artist : OneRepublic
	// URL : http://www.last.fm/music/OneRepublic/_/Counting+Stars
}

// Search for Tags.
func ExampleTagClient_Search() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Tag.Search("rock", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	var tag lastfm.TagResponse

	fmt.Printf("Search request for : %v\n", response.Query.SearchTerms)
	fmt.Printf("Total results : %v\n", response.TotalResults)
	if response.TotalResults > 0 {
		tag = response.TagMatches[0]
	}
	fmt.Printf("First result :\n")
	fmt.Printf("  Tag Name : %v\n", tag.Name)
	fmt.Printf("  Tag Count : %v\n", tag.Count)
	fmt.Printf("  Tag URL : %v\n", tag.URL)

	// Output:
	// Search request for : rock
	// Total results : 15206
	// First result :
	//   Tag Name : rock
	//   Tag Count : 3979953
	//   Tag URL : www.last.fm/tag/rock
}

// Compare Users
func ExampleTasteometerClient_Compare() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Tasteometer.Compare("user", "n3mo-", "user", "RJ", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	fmt.Printf("Compare between %v and %v :\n", response.Input.Users[0].Name, response.Input.Users[1].Name)
	fmt.Printf("Score : %v\n", response.Result.Score)
	fmt.Printf("First Artist: %v\n", response.Result.Artists[0].Name)
	// Output:
	// Compare between n3mo- and RJ :
	// Score : 0.34451797604561
	// First Artist: Iron Maiden
}

// Get full info for Track.
func ExampleTrackClient_GetInfo() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Track.GetInfo("Kids with guns", "Gorillaz", "", "", 1)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("Track : %v\n", response.Track.Name)
	fmt.Printf("MBID : %v\n", response.Track.MBID)
	fmt.Printf("URL : %v\n", response.Track.URL)
	fmt.Printf("Listeners : %v\n", response.Track.Listeners)
	fmt.Printf("Playcount : %v\n", response.Track.Playcount)
	fmt.Printf("Duration : %v\n", response.Track.Duration)
	// Output:
	// Track : Kids With Guns
	// MBID : 87fe260f-96c5-47bc-9d22-8a1c0f723475
	// URL : http://www.last.fm/music/Gorillaz/_/Kids+With+Guns
	// Listeners : 609621
	// Playcount : 3321554
	// Duration : 224000
}

// Get Similar Tracks to Track.
func ExampleTrackClient_GetSimilar() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Track.GetSimilar("Kids with guns", "Gorillaz", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	track := response.SimilarTracks[0]
	fmt.Printf("Similar Track #1 : %v\n", track.Name)
	fmt.Printf("MBID : %v\n", track.MBID)
	fmt.Printf("Artist : %v\n", track.Artist.Name)
	fmt.Printf("URL : %v\n", track.URL)

	// Output:
	// Similar Track #1 : Last Living Souls
	// MBID : 14be88ec-e5a6-4b41-b999-d4ca44ac0c52
	// Artist : Gorillaz
	// URL : http://www.last.fm/music/Gorillaz/_/Last+Living+Souls
}

// Get Tags for Track.
func ExampleTrackClient_GetTags() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Track.GetTags("Snow", "Red Hot Chili Peppers", "", "n3mo-", 1)
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
	// Name : snow
	// URL  : http://www.last.fm/tag/snow
}

// Get Top Fans for Track.
func ExampleTrackClient_GetTopFans() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Track.GetTopFans("", "", "87fe260f-96c5-47bc-9d22-8a1c0f723475", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	user := response.TopFans[0]
	fmt.Printf("Top Fan #1 : %v\n", user.Name)
	fmt.Printf("URL : %v\n", user.URL)
	fmt.Printf("Image [%v] : %v\n", user.Image[0].Size, user.Image[0].URL)

	// Output:
	// Top Fan #1 : Snerfmonster
	// URL : http://www.last.fm/user/Snerfmonster
	// Image [small] : http://userserve-ak.last.fm/serve/34/92047817.jpg
}

// Get Top Tags for Track.
func ExampleTrackClient_GetTopTags() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Track.GetTopTags("Clint Eastwood", "Gorillaz", "", 1)
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
	// Name : alternative
	// URL  : http://www.last.fm/tag/alternative
}
