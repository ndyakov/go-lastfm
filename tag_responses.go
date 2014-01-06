package lastfm

type TagResponse struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}
type TagInfoResponse struct {
	LastfmStatusResponse
	Name       string             `xml:"name"`
	URL        string             `xml:"url"`
	Reach      int64              `xml:"reach"`
	Taggings   int64              `xml:"taggings"`
	Streamable int                `xml:"streamable"`
	Wiki       LastfmWikiResponse `xml:"wiki"`
}

type TagsResponse struct {
	LastfmStatusResponse
	Tags []TagResponse `xml:"tags>tag"`
}
type TopTagsResponse struct {
	LastfmStatusResponse
	TopTags []TagResponse `xml:"toptags>tag"`
}
