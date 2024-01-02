package xmls

type SDNAka struct {
	UID       int64  `xml:"uid"`
	Type      string `xml:"type"`
	Category  string `xml:"category"`
	FirstName string `xml:"firstName"`
	LastName  string `xml:"lastName"`
}
