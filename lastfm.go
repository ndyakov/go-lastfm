package lastfm

type LastfmStatus struct {
	Status string      `xml:"status,attr"`
	Error  LastfmError `xml:"error"`
}

type LastfmError struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}

func (e *LastfmError) Error() string {
	return e.Message
}

type LastfmDate struct {
	Date string `xml:",chardata"`
	UTS  int64  `xml:"uts,attr"`
}

type LastfmImage struct {
	Size string `xml:"size,attr"`
	URL  string `xml:",chardata"`
}

type LastfmStats struct {
	Listeners int64 `xml:"listeners"`
	Playes    int64 `xml:"plays"`
}

type LastfmWiki struct {
	Published string `xml:"published"`
	Summary   string `xml:"summary"`
	Content   string `xml:"content"`
}
