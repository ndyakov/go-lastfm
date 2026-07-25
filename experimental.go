package lastfm

import (
	"context"
	"encoding/json"
	"net/url"
)

// ExperimentalClient contains legacy methods that are absent from Last.fm's
// current public method index. Last.fm may disable or change them without notice.
type ExperimentalClient struct {
	client *Client
}

// ExperimentalResult contains both the typed result and the original JSON.
type ExperimentalResult[T any] struct {
	Data T
	Raw  json.RawMessage
}

func (result *ExperimentalResult[T]) UnmarshalJSON(data []byte) error {
	result.Raw = append(result.Raw[:0], data...)
	return json.Unmarshal(data, &result.Data)
}

type TagSearchParams struct {
	Tag string
	Pagination
}

type TagSearchResponse struct {
	Results struct {
		Matches struct {
			Tags List[Tag] `json:"tag"`
		} `json:"tagmatches"`
		Attr PaginationResponse `json:"@attr"`
	} `json:"results"`
}

func (c ExperimentalClient) SearchTags(ctx context.Context, params TagSearchParams) (*ExperimentalResult[TagSearchResponse], error) {
	values := url.Values{}
	set(values, "tag", params.Tag)
	params.Pagination.values(values)
	result := new(ExperimentalResult[TagSearchResponse])
	return result, c.client.call(ctx, "tag.search", values, result)
}

type Metro struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type MetrosResponse struct {
	Metros struct {
		Metros List[Metro] `json:"metro"`
	} `json:"metros"`
}

func (c ExperimentalClient) GetMetros(ctx context.Context, country string) (*ExperimentalResult[MetrosResponse], error) {
	values := url.Values{}
	set(values, "country", country)
	result := new(ExperimentalResult[MetrosResponse])
	return result, c.client.call(ctx, "geo.getMetros", values, result)
}

type NeighboursParams struct {
	User  string
	Limit int
}

type NeighboursResponse struct {
	Neighbours struct {
		Users List[User] `json:"user"`
	} `json:"neighbours"`
}

func (c ExperimentalClient) GetNeighbours(ctx context.Context, params NeighboursParams) (*ExperimentalResult[NeighboursResponse], error) {
	values := url.Values{}
	set(values, "user", params.User)
	setInt(values, "limit", params.Limit)
	result := new(ExperimentalResult[NeighboursResponse])
	return result, c.client.call(ctx, "user.getNeighbours", values, result)
}

type TopFansParams struct {
	Artist string
	Track  string
	MBID   string
}

type Fan struct {
	User
	Weight Integer `json:"weight"`
}

type TopFansResponse struct {
	TopFans struct {
		Fans List[Fan] `json:"user"`
	} `json:"topfans"`
}

func (c ExperimentalClient) GetArtistTopFans(ctx context.Context, params TopFansParams) (*ExperimentalResult[TopFansResponse], error) {
	values := url.Values{}
	set(values, "artist", params.Artist)
	set(values, "mbid", params.MBID)
	result := new(ExperimentalResult[TopFansResponse])
	return result, c.client.call(ctx, "artist.getTopFans", values, result)
}

func (c ExperimentalClient) GetTrackTopFans(ctx context.Context, params TopFansParams) (*ExperimentalResult[TopFansResponse], error) {
	values := url.Values{}
	set(values, "artist", params.Artist)
	set(values, "track", params.Track)
	set(values, "mbid", params.MBID)
	result := new(ExperimentalResult[TopFansResponse])
	return result, c.client.call(ctx, "track.getTopFans", values, result)
}

type TasteometerParams struct {
	Type1  string
	Value1 string
	Type2  string
	Value2 string
	Limit  int
}

type TasteometerResponse struct {
	Comparison struct {
		Result struct {
			Score   Float `json:"score"`
			Artists struct {
				Artists List[Artist] `json:"artist"`
			} `json:"artists"`
		} `json:"result"`
	} `json:"comparison"`
}

func (c ExperimentalClient) CompareTaste(ctx context.Context, params TasteometerParams) (*ExperimentalResult[TasteometerResponse], error) {
	values := url.Values{}
	set(values, "type1", params.Type1)
	set(values, "value1", params.Value1)
	set(values, "type2", params.Type2)
	set(values, "value2", params.Value2)
	setInt(values, "limit", params.Limit)
	result := new(ExperimentalResult[TasteometerResponse])
	return result, c.client.call(ctx, "tasteometer.compare", values, result)
}
