package xmls

type SDNPlaceOfBirth struct {
	UID       int64  `xml:"uid"`
	Place     string `xml:"place"`
	MainEntry bool   `xml:"mainEntry"`
}
