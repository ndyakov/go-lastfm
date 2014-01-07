package lastfm

type TasteometerClient struct {
	Client
}

func (c *TasteometerClient) Compare(type1, value1, type2, value2 string, limit int) (response TasteometerCompareResponse, err error) {
	return
}
