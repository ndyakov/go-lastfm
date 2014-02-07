package lastfm

// TagResponse, used where <tag> tag is present.
type TagResponse struct {
	Name       string             `xml:"name"`
	URL        string             `xml:"url"`
	Reach      int64              `xml:"reach"`
	Taggings   int64              `xml:"taggings"`
	Streamable int                `xml:"streamable"`
	Count      int64              `xml:"count"`
	Wiki       LastfmWikiResponse `xml:"wiki"`
}

// TagInfoResponse, used for tag.getInfo's response.
type TagInfoResponse struct {
	LastfmStatusResponse
	Tag TagResponse `xml:"tag"`
}

// TagsResponse, used where there are multiple tags.
type TagsResponse struct {
	LastfmStatusResponse
	Tags []TagResponse `xml:"tags>tag"`
}

// TopTagResponse, used where <toptags> is present.
type TopTagsResponse struct {
	LastfmStatusResponse
	TopTags []TagResponse `xml:"toptags>tag"`
}

// TagSimilarResponse, used where <similartags> is present.
// For example in tag.getSimilar's response.
type TagSimilarResponse struct {
	LastfmStatusResponse
	SimilarTags []TagResponse `xml:"similartags>tag"`
}

// TagSearchResponse, used for tag.search's response.
type TagSearchResponse struct {
	LastfmStatusResponse
	LastfmOpenSearchResponse
	TagMatches []TagResponse `xml:"results>tagmatches>tag"`
}
