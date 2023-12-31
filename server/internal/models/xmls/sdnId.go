package xmls

type SDNId struct {
	UID       int64   `xml:"uid"`
	Type      string  `xml:"idType"`
	Number    string  `xml:"idNumber"`
	Country   string  `xml:"idCountry"`
	IssueDate XMLDate `xml:"issueDate"`
}
