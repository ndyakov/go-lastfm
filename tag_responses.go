package lastfm

type TagResponse struct {
	Name       string             `xml:"name"`
	URL        string             `xml:"url"`
	Reach      int64              `xml:"reach"`
	Taggings   int64              `xml:"taggings"`
	Streamable int                `xml:"streamable"`
	Wiki       LastfmWikiResponse `xml:"wiki"`
}
