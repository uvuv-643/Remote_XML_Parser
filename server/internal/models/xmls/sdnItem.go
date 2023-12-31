package xmls

type SDNItem struct {
	UID       int64  `xml:"uid"`
	FirstName string `xml:"firstName"`
	LastName  string `xml:"lastName"`
	Title     string `xml:"title"`
	SDNType   string `xml:"sdnType"`
	Remarks   string `xml:"remarks"`

	ProgramList struct {
		Program []SDNProgram `xml:"program"`
	} `xml:"programList"`

	AkaList struct {
		Aka []SDNAka `xml:"aka"`
	} `xml:"akaList"`

	IdList struct {
		ID []SDNId `xml:"id"`
	} `xml:"idList"`

	AddressList struct {
		Address []SDNAddress `xml:"address"`
	} `xml:"addressList"`

	NationalityList struct {
		Nationality []SDNNationality `xml:"nationality"`
	} `xml:"nationalityList"`

	DateOfBirthList struct {
		DateOfBirthItem []SDNDateOfBirth `xml:"dateOfBirthItem"`
	} `xml:"dateOfBirthList"`

	PlaceOfBirthList struct {
		PlaceOfBirthItem []SDNPlaceOfBirth `xml:"placeOfBirthItem"`
	} `xml:"placeOfBirthList"`

	CitizenshipList struct {
		Citizenship []SDNCitizenship `xml:"citizenship"`
	} `xml:"citizenshipList"`
}
