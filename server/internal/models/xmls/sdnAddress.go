package xmls

type SDNAddress struct {
	UID             int64  `xml:"uid"`
	City            string `xml:"city"`
	Address1        string `xml:"address1"`
	Address2        string `xml:"address2"`
	Address3        string `xml:"address3"`
	StateOrProvince string `xml:"stateOrProvince"`
	PostalCode      string `xml:"postalCode"`
	Country         string `xml:"country"`
}
