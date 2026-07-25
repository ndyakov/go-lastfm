package lastfm

import (
	"context"
	"net/url"
)

type UserClient struct{ client *Client }

type UserPagedParams struct {
	User                string
	IncludeRecentTracks bool
	Pagination
}

func (c UserClient) GetFriends(ctx context.Context, p UserPagedParams) (*FriendsResponse, error) {
	v := url.Values{}
	set(v, "user", p.User)
	setBool(v, "recenttracks", p.IncludeRecentTracks)
	p.Pagination.values(v)
	out := new(FriendsResponse)
	return out, c.client.call(ctx, "user.getFriends", v, out)
}

func (c UserClient) GetInfo(ctx context.Context, user string) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	return out, c.client.call(ctx, "user.getInfo", url.Values{"user": {user}}, out)
}

func (c UserClient) GetLovedTracks(ctx context.Context, p UserPagedParams) (*LovedTracksResponse, error) {
	v := url.Values{}
	set(v, "user", p.User)
	p.Pagination.values(v)
	out := new(LovedTracksResponse)
	return out, c.client.call(ctx, "user.getLovedTracks", v, out)
}

type PersonalTagsParams struct {
	User, Tag, TaggingType string
	Pagination
}

func (c UserClient) GetPersonalTags(ctx context.Context, p PersonalTagsParams) (*PersonalTagsResponse, error) {
	v := url.Values{}
	set(v, "user", p.User)
	set(v, "tag", p.Tag)
	set(v, "taggingtype", p.TaggingType)
	p.Pagination.values(v)
	out := new(PersonalTagsResponse)
	return out, c.client.call(ctx, "user.getPersonalTags", v, out)
}

type RecentTracksParams struct {
	User string
	Pagination
	From, To int64
	Extended bool
}

func (c UserClient) GetRecentTracks(ctx context.Context, p RecentTracksParams) (*RecentTracksResponse, error) {
	v := url.Values{}
	set(v, "user", p.User)
	p.Pagination.values(v)
	setInt64(v, "from", p.From)
	setInt64(v, "to", p.To)
	setBool(v, "extended", p.Extended)
	out := new(RecentTracksResponse)
	return out, c.client.call(ctx, "user.getRecentTracks", v, out)
}

type Period string

const (
	PeriodOverall Period = "overall"
	Period7Day    Period = "7day"
	Period1Month  Period = "1month"
	Period3Month  Period = "3month"
	Period6Month  Period = "6month"
	Period12Month Period = "12month"
)

type UserTopParams struct {
	User   string
	Period Period
	Pagination
}

func userTopValues(p UserTopParams) url.Values {
	v := url.Values{}
	set(v, "user", p.User)
	set(v, "period", string(p.Period))
	p.Pagination.values(v)
	return v
}

func (c UserClient) GetTopAlbums(ctx context.Context, p UserTopParams) (*TopAlbumsResponse, error) {
	out := new(TopAlbumsResponse)
	return out, c.client.call(ctx, "user.getTopAlbums", userTopValues(p), out)
}

func (c UserClient) GetTopArtists(ctx context.Context, p UserTopParams) (*TopArtistsResponse, error) {
	out := new(TopArtistsResponse)
	return out, c.client.call(ctx, "user.getTopArtists", userTopValues(p), out)
}

func (c UserClient) GetTopTracks(ctx context.Context, p UserTopParams) (*TopTracksResponse, error) {
	out := new(TopTracksResponse)
	return out, c.client.call(ctx, "user.getTopTracks", userTopValues(p), out)
}

type UserTopTagsParams struct {
	User  string
	Limit int
}

func (c UserClient) GetTopTags(ctx context.Context, p UserTopTagsParams) (*TopTagsResponse, error) {
	v := url.Values{}
	set(v, "user", p.User)
	setInt(v, "limit", p.Limit)
	out := new(TopTagsResponse)
	return out, c.client.call(ctx, "user.getTopTags", v, out)
}

type WeeklyChartParams struct {
	User     string
	From, To int64
}

func weeklyValues(p WeeklyChartParams) url.Values {
	v := url.Values{}
	set(v, "user", p.User)
	setInt64(v, "from", p.From)
	setInt64(v, "to", p.To)
	return v
}

func (c UserClient) GetWeeklyAlbumChart(ctx context.Context, p WeeklyChartParams) (*WeeklyAlbumChartResponse, error) {
	out := new(WeeklyAlbumChartResponse)
	return out, c.client.call(ctx, "user.getWeeklyAlbumChart", weeklyValues(p), out)
}

func (c UserClient) GetWeeklyArtistChart(ctx context.Context, p WeeklyChartParams) (*WeeklyArtistChartResponse, error) {
	out := new(WeeklyArtistChartResponse)
	return out, c.client.call(ctx, "user.getWeeklyArtistChart", weeklyValues(p), out)
}

func (c UserClient) GetWeeklyTrackChart(ctx context.Context, p WeeklyChartParams) (*WeeklyTrackChartResponse, error) {
	out := new(WeeklyTrackChartResponse)
	return out, c.client.call(ctx, "user.getWeeklyTrackChart", weeklyValues(p), out)
}

func (c UserClient) GetWeeklyChartList(ctx context.Context, user string) (*WeeklyChartListResponse, error) {
	out := new(WeeklyChartListResponse)
	return out, c.client.call(ctx, "user.getWeeklyChartList", url.Values{"user": {user}}, out)
}
