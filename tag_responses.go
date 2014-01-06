package lastfm

type TagResponse struct {
	Name       string `xml:"name"`
	URL        string `xml:"url"`
	Reach      int64  `xml:"reach"`
	Taggings   int64  `xml:"taggings"`
	Streamable int    `xml:"streamable"`
}

type TagInfoResponse struct {
	LastfmStatusResponse
	TagResponse
	Wiki LastfmWikiResponse `xml:"wiki"`
}

type TagsResponse struct {
	LastfmStatusResponse
	Tags []TagResponse `xml:"tags>tag"`
}

type TopTagsResponse struct {
	LastfmStatusResponse
	TopTags []TagResponse `xml:"toptags>tag"`
}

type TagSimilarResponse struct {
	LastfmStatusResponse
	SimilarTags []TagResponse `xml:"similartags>tag"`
}

type TagSearchResponse struct {
	LastfmStatusResponse
	LastfmOpenSearchResponse
	TagMatches []TagResponse `xml:"tagmatches>tag"`
}
