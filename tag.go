package lastfm

type Tag struct {
	Name      string     `xml:"name"`
	URL       string     `xml:"url"`
	Reach     int64      `xml:"reach"`
	Taggings  int64      `xml:"taggings"`
	Stramable int        `xml:"streamable"`
	Wiki      LastfmWiki `xml:"wiki"`
}
