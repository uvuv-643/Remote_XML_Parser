package xmls

type SDNDateOfBirth struct {
	UID         int64   `xml:"uid"`
	DateOfBirth XMLDate `xml:"dateOfBirth"`
	MainEntry   bool    `xml:"mainEntry"`
}
