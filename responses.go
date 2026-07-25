package lastfm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// Integer accepts both JSON numbers and the quoted numbers returned by Last.fm.
type Integer int64

func (i *Integer) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, `"`)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		*i = 0
		return nil
	}
	value, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return fmt.Errorf("lastfm: decode integer %q: %w", data, err)
	}
	*i = Integer(value)
	return nil
}

// Float accepts both JSON numbers and quoted decimal values.
type Float float64

func (f *Float) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, `"`)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		*f = 0
		return nil
	}
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return fmt.Errorf("lastfm: decode float %q: %w", data, err)
	}
	*f = Float(value)
	return nil
}

// List accepts both an array and Last.fm's single-object representation.
type List[T any] []T

func (list *List[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		*list = nil
		return nil
	}
	var values []T
	if err := json.Unmarshal(data, &values); err == nil {
		*list = values
		return nil
	}
	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*list = []T{value}
	return nil
}

type Image struct {
	Size string `json:"size"`
	URL  string `json:"#text"`
}

type Date struct {
	Text string  `json:"#text"`
	UTS  Integer `json:"uts"`
}

type PaginationResponse struct {
	Page       Integer `json:"page"`
	PerPage    Integer `json:"perPage"`
	TotalPages Integer `json:"totalPages"`
	Total      Integer `json:"total"`
}

type Wiki struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

type Stats struct {
	Listeners Integer `json:"listeners"`
	Playcount Integer `json:"playcount"`
}

type Artist struct {
	Name       string      `json:"name"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Images     List[Image] `json:"image"`
	Streamable Integer     `json:"streamable"`
	Match      Float       `json:"match"`
	Playcount  Integer     `json:"playcount"`
	Listeners  Integer     `json:"listeners"`
	TagCount   Integer     `json:"tagcount"`
	Stats      Stats       `json:"stats"`
	Bio        Wiki        `json:"bio"`
	Similar    struct {
		Artists List[Artist] `json:"artist"`
	} `json:"similar"`
	Tags struct {
		Tags List[Tag] `json:"tag"`
	} `json:"tags"`
	Attr struct {
		Rank Integer `json:"rank"`
	} `json:"@attr"`
}

type Album struct {
	Name        string      `json:"name"`
	Artist      any         `json:"artist"`
	MBID        string      `json:"mbid"`
	URL         string      `json:"url"`
	Images      List[Image] `json:"image"`
	Listeners   Integer     `json:"listeners"`
	Playcount   Integer     `json:"playcount"`
	ReleaseDate string      `json:"releasedate"`
	Tracks      struct {
		Tracks List[Track] `json:"track"`
	} `json:"tracks"`
	Tags struct {
		Tags List[Tag] `json:"tag"`
	} `json:"tags"`
	Wiki Wiki `json:"wiki"`
}

type Tag struct {
	Name       string  `json:"name"`
	URL        string  `json:"url"`
	Reach      Integer `json:"reach"`
	Taggings   Integer `json:"taggings"`
	Streamable Integer `json:"streamable"`
	Count      Integer `json:"count"`
	Wiki       Wiki    `json:"wiki"`
}

type Track struct {
	Name       string      `json:"name"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Duration   Integer     `json:"duration"`
	Listeners  Integer     `json:"listeners"`
	Playcount  Integer     `json:"playcount"`
	UserLoved  Integer     `json:"userloved"`
	Artist     Artist      `json:"artist"`
	Album      Album       `json:"album"`
	Images     List[Image] `json:"image"`
	Wiki       Wiki        `json:"wiki"`
	Streamable struct {
		FullTrack Integer `json:"fulltrack"`
		Text      Integer `json:"#text"`
	} `json:"streamable"`
	Attr struct {
		Rank       Integer `json:"rank"`
		NowPlaying string  `json:"nowplaying"`
	} `json:"@attr"`
}

type User struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	RealName   string      `json:"realname"`
	URL        string      `json:"url"`
	Images     List[Image] `json:"image"`
	Country    string      `json:"country"`
	Age        Integer     `json:"age"`
	Gender     string      `json:"gender"`
	Subscriber Integer     `json:"subscriber"`
	Playcount  Integer     `json:"playcount"`
	Playlists  Integer     `json:"playlists"`
	Bootstrap  Integer     `json:"bootstrap"`
	Registered Date        `json:"registered"`
}

type AlbumInfoResponse struct {
	Album Album `json:"album"`
}

type ArtistInfoResponse struct {
	Artist Artist `json:"artist"`
}

type TrackInfoResponse struct {
	Track Track `json:"track"`
}

type TagInfoResponse struct {
	Tag Tag `json:"tag"`
}

type UserInfoResponse struct {
	User User `json:"user"`
}

type TagsResponse struct {
	Tags struct {
		Tags List[Tag] `json:"tag"`
	} `json:"tags"`
}

type TopTagsResponse struct {
	TopTags struct {
		Tags List[Tag]          `json:"tag"`
		Attr PaginationResponse `json:"@attr"`
	} `json:"toptags"`
}

type TopAlbumsResponse struct {
	TopAlbums struct {
		Albums List[Album]        `json:"album"`
		Attr   PaginationResponse `json:"@attr"`
	} `json:"topalbums"`
}

type TopArtistsResponse struct {
	TopArtists struct {
		Artists List[Artist]       `json:"artist"`
		Attr    PaginationResponse `json:"@attr"`
	} `json:"topartists"`
}

type TopTracksResponse struct {
	TopTracks struct {
		Tracks List[Track]        `json:"track"`
		Attr   PaginationResponse `json:"@attr"`
	} `json:"toptracks"`
}

type AlbumSearchResponse struct {
	Results struct {
		Matches struct {
			Albums List[Album] `json:"album"`
		} `json:"albummatches"`
	} `json:"results"`
}

type ArtistSearchResponse struct {
	Results struct {
		Matches struct {
			Artists List[Artist] `json:"artist"`
		} `json:"artistmatches"`
	} `json:"results"`
}

type TrackSearchResponse struct {
	Results struct {
		Matches struct {
			Tracks List[Track] `json:"track"`
		} `json:"trackmatches"`
	} `json:"results"`
}

type ArtistCorrectionResponse struct {
	Corrections struct {
		Corrections List[struct {
			Artist Artist `json:"artist"`
		}] `json:"correction"`
	} `json:"corrections"`
}

type TrackCorrectionResponse struct {
	Corrections struct {
		Corrections List[struct {
			Track Track `json:"track"`
		}] `json:"correction"`
	} `json:"corrections"`
}

type ArtistSimilarResponse struct {
	Similar struct {
		Artists List[Artist]       `json:"artist"`
		Attr    PaginationResponse `json:"@attr"`
	} `json:"similarartists"`
}

type SimilarTracksResponse struct {
	Similar struct {
		Tracks List[Track]        `json:"track"`
		Attr   PaginationResponse `json:"@attr"`
	} `json:"similartracks"`
}

type TagSimilarResponse struct {
	Similar struct {
		Tags List[Tag] `json:"tag"`
	} `json:"similartags"`
}

type AuthTokenResponse struct {
	Token string `json:"token"`
}

type Session struct {
	Name       string  `json:"name"`
	Key        string  `json:"key"`
	Subscriber Integer `json:"subscriber"`
}

type AuthSessionResponse struct {
	Session Session `json:"session"`
}

type StatusResponse struct{}

type ChartArtistsResponse struct {
	Artists struct {
		Artists List[Artist]       `json:"artist"`
		Attr    PaginationResponse `json:"@attr"`
	} `json:"artists"`
}
type ChartTagsResponse struct {
	Tags struct {
		Tags List[Tag]          `json:"tag"`
		Attr PaginationResponse `json:"@attr"`
	} `json:"tags"`
}
type ChartTracksResponse struct {
	Tracks struct {
		Tracks List[Track]        `json:"track"`
		Attr   PaginationResponse `json:"@attr"`
	} `json:"tracks"`
}
type LibraryArtistsResponse struct {
	Artists struct {
		Artists List[Artist]       `json:"artist"`
		Attr    PaginationResponse `json:"@attr"`
	} `json:"artists"`
}

type WeeklyChart struct {
	From Integer `json:"from"`
	To   Integer `json:"to"`
}

type WeeklyChartAttributes struct {
	User string `json:"user"`
	Tag  string `json:"tag"`
}

type WeeklyChartListResponse struct {
	Charts struct {
		Charts List[WeeklyChart]     `json:"chart"`
		Attr   WeeklyChartAttributes `json:"@attr"`
	} `json:"weeklychartlist"`
}

type TrackScrobbleField struct {
	Text      string  `json:"#text"`
	Corrected Integer `json:"corrected"`
}

type ScrobbleResult struct {
	Track          TrackScrobbleField `json:"track"`
	Artist         TrackScrobbleField `json:"artist"`
	Album          TrackScrobbleField `json:"album"`
	AlbumArtist    TrackScrobbleField `json:"albumArtist"`
	Timestamp      Integer            `json:"timestamp"`
	IgnoredMessage struct {
		Code Integer `json:"code"`
		Text string  `json:"#text"`
	} `json:"ignoredMessage"`
}

type ScrobbleAttributes struct {
	Accepted Integer `json:"accepted"`
	Ignored  Integer `json:"ignored"`
}

type TrackScrobbleResponse struct {
	Scrobbles struct {
		Scrobbles List[ScrobbleResult] `json:"scrobble"`
		Attr      ScrobbleAttributes   `json:"@attr"`
	} `json:"scrobbles"`
}

type TrackNowPlayingResponse struct {
	NowPlaying ScrobbleResult `json:"nowplaying"`
}

type LovedTracksResponse struct {
	LovedTracks struct {
		Tracks List[Track]        `json:"track"`
		Attr   PaginationResponse `json:"@attr"`
	} `json:"lovedtracks"`
}
type RecentTracksResponse struct {
	RecentTracks struct {
		Tracks List[Track]        `json:"track"`
		Attr   PaginationResponse `json:"@attr"`
	} `json:"recenttracks"`
}
type FriendsResponse struct {
	Friends struct {
		Users List[User]         `json:"user"`
		Attr  PaginationResponse `json:"@attr"`
	} `json:"friends"`
}
type PersonalTagsResponse struct {
	Taggings struct {
		Artists struct {
			Artists List[Artist] `json:"artist"`
		} `json:"artists"`
		Albums struct {
			Albums List[Album] `json:"album"`
		} `json:"albums"`
		Tracks struct {
			Tracks List[Track] `json:"track"`
		} `json:"tracks"`
		Attr PaginationResponse `json:"@attr"`
	} `json:"taggings"`
}

type WeeklyPeriodAttributes struct {
	From Integer `json:"from"`
	To   Integer `json:"to"`
}

type WeeklyAlbumChartResponse struct {
	Chart struct {
		Albums List[Album]            `json:"album"`
		Attr   WeeklyPeriodAttributes `json:"@attr"`
	} `json:"weeklyalbumchart"`
}
type WeeklyArtistChartResponse struct {
	Chart struct {
		Artists List[Artist]           `json:"artist"`
		Attr    WeeklyPeriodAttributes `json:"@attr"`
	} `json:"weeklyartistchart"`
}
type WeeklyTrackChartResponse struct {
	Chart struct {
		Tracks List[Track]            `json:"track"`
		Attr   WeeklyPeriodAttributes `json:"@attr"`
	} `json:"weeklytrackchart"`
}
