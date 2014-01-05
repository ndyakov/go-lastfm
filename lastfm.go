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
