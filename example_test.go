package lastfm_test

import (
	"fmt"
	"github.com/ndyakov/go-lastfm"
	"reflect"
	"strings"
)

// Create new api client object
// with your api key as argument.
func ExampleLastFM_New() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	fmt.Printf("The type is %v\n", reflect.TypeOf(lfm))
	// Output:
	// The type is *lastfm.LastFM
}

// Get Top Tags for Artist.
func ExampleArtistClient_GetTopTags() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Artist.GetTopTags("Daft Punk", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Artist.GetTags("Red Hot Chili Peppers", "", "RJ", 1)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Artist.GetSimilar("Cher", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Artist.GetTopAlbums("Ogonek", "", 0, 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Artist.GetTopFans("Ogonek", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Artist.GetTopTracks("Ogonek", "", 0, 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Artist.Search("Ogo", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Artist.GetInfo("Gorillaz", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Album.GetInfo("Ogonek", "Drum And Bass Massacre", "", "", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Album.GetTags("Guano Apes", "Bel Air", "", "n3mo-", 1)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Album.GetTopTags("Guano Apes", "Bel Air", "", 1)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Album.Search("go", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Tag.GetInfo("rock")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Tag.GetSimilar("rock")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Tag.GetTopAlbums("dnb", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Tag.GetTopArtists("dnb", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Tag.GetTopTags()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Tag.GetTopTracks("rock", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Tag.Search("rock", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Tasteometer.Compare("user", "n3mo-", "user", "RJ", 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Track.GetInfo("Kids with guns", "Gorillaz", map[string]string{"autocorrect": "1"})
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Track.GetSimilar("Kids with guns", "Gorillaz", map[string]string{})
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Track.GetTags("Snow", "Red Hot Chili Peppers", map[string]string{"user": "n3mo-", "autocorrect": "1"})

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Track.GetTopFans("", "", map[string]string{"mbid": "87fe260f-96c5-47bc-9d22-8a1c0f723475"})
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Track.GetTopTags("Clint Eastwood", "Gorillaz", map[string]string{"autocorrect": "1"})
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
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

// Get Track correction.
func ExampleTrackClient_GetCorrection() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Track.GetCorrection("Clint eastwood", "Gorilas")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	fmt.Printf("Corrections : %v\n", len(response.TrackCorrections))

	if len(response.TrackCorrections) > 0 {
		var track lastfm.TrackResponse
		track = response.TrackCorrections[0].Track
		fmt.Printf("Correction #1 :\n")
		fmt.Printf("  Corrected: [%v] Artist , [%v] Track\n", response.TrackCorrections[0].ArtistCorrected, response.TrackCorrections[0].TrackCorrected)
		fmt.Printf("  Track : %v\n", track.Name)
		fmt.Printf("  Artist : %v\n", track.Artist.Name)
	}
	// Output:
	// Corrections : 1
	// Correction #1 :
	//   Corrected: [1] Artist , [0] Track
	//   Track : Clint Eastwood
	//   Artist : Gorillaz
}

// Search for Track.
func ExampleTrackClient_Search() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Track.Search("guns", map[string]string{"artist": "Gorillaz"})
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}

	fmt.Printf("Search request for : %v\n", response.Query.SearchTerms)
	fmt.Printf("Total results : %v\n", response.TotalResults)

	if response.TotalResults > 0 {
		var track lastfm.TrackResponseNoArtistStruct
		track = response.TrackMatches[0]
		fmt.Printf("First result :\n")
		fmt.Printf("  Track Name : %v\n", track.Name)
		fmt.Printf("  Artist Name : %v\n", track.Artist)
		fmt.Printf("  Track MBID : %v\n", track.MBID)
	}

	// Output:
	// Search request for : guns
	// Total results : 260
	// First result :
	//   Track Name : Kids With Guns
	//   Artist Name : Gorillaz
	//   Track MBID : 87fe260f-96c5-47bc-9d22-8a1c0f723475
}

func ExampleTrackClient_Scrobble() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	_, err := lfm.Auth.GetSession()
	response, err := lfm.Track.Scrobble("Balkansky", "8.9", "1400380586", map[string]string{})

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}

	fmt.Printf("Accepted: %d\n", response.Scrobbles.Accepted)
	fmt.Printf("Ignored: %d\n", response.Scrobbles.Ignored)
	fmt.Println("Scrobbles[0]:")
	fmt.Printf("Artist: %s\n", response.Scrobbles.Scrobble[0].Artist.Name)
	fmt.Printf("Track: %s\n", response.Scrobbles.Scrobble[0].Track.Name)
	//Output:
	//Accepted: 1
	//Ignored: 0
	//Scrobbles[0]:
	//Artist: Balkansky
	//Track: 8.9
}

// Get full info for User.
func ExampleUserClient_GetInfo() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.User.GetInfo("RJ")

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}

	fmt.Printf("User : %v\n", response.User.Name)
	fmt.Printf("ID : %v\n", response.User.ID)
	// In this example the user had extra space
	// in his "Real Name" field, we have to trim it
	// for the test to pass.
	fmt.Printf("Real Name : %s\n", strings.Trim(response.User.RealName, " "))
	fmt.Printf("URL : %v\n", response.User.URL)
	fmt.Printf("Country : %v\n", response.User.Country)
	fmt.Printf("Registered : %v\n", response.User.Registered.Date)
	// Output:
	// User : RJ
	// ID : 1000002
	// Real Name : Richard Jones
	// URL : http://www.last.fm/user/RJ
	// Country : UK
	// Registered : 2002-11-20 11:50
}

// Get Loved tracks from user.
func ExampleUserClient_GetLovedTracks() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.User.GetLovedTracks("RJ", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	track := response.LovedTracks[0]
	fmt.Printf("Loved Tracks by RJ\n")
	fmt.Printf("Track #1 : %v\n", track.Name)
	fmt.Printf("URL : %v\n", track.URL)
	fmt.Printf("Artist Name : %v\n", track.Artist.Name)

	// Output:
	// Loved Tracks by RJ
	// Track #1 : New Year
	// URL : http://www.last.fm/music/Beach+House/_/New+Year
	// Artist Name : Beach House
}

// Get neighbours to some user.
func ExampleUserClient_GetNeighbours() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.User.GetNeighbours("RJ", 2)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	user := response.Neighbours[0]
	fmt.Printf("Neighbour #1 : %v\n", user.Name)
	fmt.Printf("Match : %v\n", user.Match)
	fmt.Printf("URL : %v\n", user.URL)

	// Output:
	// Neighbour #1 : tonyetc
	// Match : 0.99757462739944
	// URL : http://www.last.fm/user/tonyetc
}

// Get Recent Tracks for User.
func ExampleUserClient_GetRecentTracks() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.User.GetRecentTracks("RJ", 0, 0, 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	track := response.RecentTracks[0]
	fmt.Printf("Recent Track #1 : %v\n", track.Name)
	if track.NowPlaying != "" {
		fmt.Printf("!! Now Playing !!\n")
	}
	fmt.Printf("Loved : %v\n", track.Loved)
	fmt.Printf("Artist : %v\n", track.Artist.Name)
	fmt.Printf("Album : %v\n", track.Album.Name)
	fmt.Printf("URL : %v\n", track.URL)

	// Output:
	// Recent Track #1 : Novacane
	// Loved : 1
	// Artist : Frank Ocean
	// Album : Novacane
	// URL : http://www.last.fm/music/Frank+Ocean/_/Novacane
}

// Get Top Albums for User.
func ExampleUserClient_GetTopAlbums() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.User.GetTopAlbums("RJ", "", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	album := response.TopAlbums[0]
	fmt.Printf("Top Album #1 : %v\n", album.Name)
	fmt.Printf("Artist Name : %v\n", album.Artist.Name)
	fmt.Printf("URL : %v\n", album.URL)
	fmt.Printf("Playcount: %v\n", album.Playcount)
	fmt.Printf("Image [%v] : %v\n", album.Image[0].Size, album.Image[0].URL)

	// Output:
	// Top Album #1 : Images and Words
	// Artist Name : Dream Theater
	// URL : http://www.last.fm/music/Dream+Theater/Images+and+Words
	// Playcount: 266
	// Image [small] : http://userserve-ak.last.fm/serve/34s/93189941.png
}

// Get Top Artists for User.
func ExampleUserClient_GetTopArtists() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.User.GetTopArtists("RJ", "", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	artist := response.TopArtists[0]
	fmt.Printf("Artist #1 : %v\n", artist.Name)
	fmt.Printf("URL : %v\n", artist.URL)
	fmt.Printf("MBID : %v\n", artist.MBID)

	// Output:
	// Artist #1 : Dream Theater
	// URL : http://www.last.fm/music/Dream+Theater
	// MBID : 28503ab7-8bf2-4666-a7bd-2644bfc7cb1d
}

// Get Top Tracks for User.
func ExampleUserClient_GetTopTracks() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.User.GetTopTracks("RJ", "", 0, 0)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	track := response.TopTracks[0]
	fmt.Printf("Top Track #1 : %v\n", track.Name)
	fmt.Printf("Artist Name : %v\n", track.Artist.Name)
	fmt.Printf("URL : %v\n", track.URL)
	fmt.Printf("Playcount: %v\n", track.Playcount)

	// Output:
	// Top Track #1 : Sultans of Swing
	// Artist Name : Dire Straits
	// URL : http://www.last.fm/music/Dire+Straits/_/Sultans+of+Swing
	// Playcount: 78
}

// Get Top Tags by User.
func ExampleUserClient_GetTopTags() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.User.GetTopTags("RJ", 2)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of Top Tags : %v\n", len(response.TopTags))
	fmt.Printf("Tag #1 :\n")
	fmt.Printf("Name : %v\n", response.TopTags[0].Name)
	fmt.Printf("Count : %v\n", response.TopTags[0].Count)
	fmt.Printf("URL  : %v\n", response.TopTags[0].URL)
	// Output:
	// Number of Top Tags : 2
	// Tag #1 :
	// Name : rock
	// Count : 19
	// URL  : www.last.fm/tag/rock
}

// Get Token
func ExampleAuthClient_GetToken() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Auth.GetToken()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	fmt.Printf("Token : %s", response.Token)
	// Output:
	// Token : fba51a644f342186ba4d579d5675a167
}

// Get Session
func ExampleAuthClient_GetSession() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Auth.GetSession()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	fmt.Printf("Name : %s\n", response.Session.Name)
	fmt.Printf("SessionKey : %s\n", response.Session.Key)
	fmt.Printf("Subscriber : %d\n", response.Session.Subscriber)

	// Output:
	// Name : MyLastFMUsername
	// SessionKey : d580d57f32848f5dcf574d1ce18d78b2
	// Subscriber : 0
}

// Get Session
func ExampleAuthClient_GetMobileSession() {
	lfm := lastfm.New("api_key_for_testing", "api_secret_for_testing")
	response, err := lfm.Auth.GetMobileSession("testPassword", "testUserName")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	fmt.Printf("Name : %s\n", response.Session.Name)
	fmt.Printf("SessionKey : %s\n", response.Session.Key)
	fmt.Printf("Subscriber : %d\n", response.Session.Subscriber)

	// Output:
	// Name : MyLastFMUsername
	// SessionKey : d580d57f32848f5dcf574d1ce18d78b2
	// Subscriber : 0
}
