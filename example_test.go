package lastfm_test

import (
	"fmt"
	"github.com/ndyakov/go-lastfm"
)

func ExampleLastFM_Artist_GetTopTags() {
	lfm := lastfm.New("api_key_for_testing")
	response, err := lfm.Artist.GetTopTags("Daft Punk", "", 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of Top Tags %v\n", len(response.TopTags))
	// Output:
	// Number of Top Tags 100
}
