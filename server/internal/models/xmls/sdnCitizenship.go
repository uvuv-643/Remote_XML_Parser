package xmls

type SDNCitizenship struct {
	UID       int64  `xml:"uid"`
	Country   string `xml:"country"`
	MainEntry bool   `xml:"mainEntry"`
}
