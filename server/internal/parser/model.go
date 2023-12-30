package parser

import "encoding/xml"

type SDN struct {
	SDNList            xml.Name `xml:"sdnList"`
	PublishInformation struct {
		PublishDate string `xml:"Publish_Date"`
		RecordCount int    `xml:"Record_Count"`
	} `xml:"publshInformation"`

	SDNEntry []struct {
		UID       string `xml:"uid"`
		FirstName string `xml:"firstName"`
		LastName  string `xml:"lastName"`
		Title     string `xml:"title"`
		SDNType   string `xml:"sdnType"`
		Remarks   string `xml:"remarks"`

		ProgramList struct {
			Program []string `xml:"program"`
		} `xml:"programList"`

		AkaList struct {
			Aka []struct {
				UID      string   `xml:"uid"`
				Type     string   `xml:"type"`
				Category string   `xml:"category"`
				LastName xml.Name `xml:"lastName"`
			} `xml:"aka"`
		} `xml:"akaList"`

		IdList struct {
			ID []struct {
				UID       string `xml:"uid"`
				Type      string `xml:"idType"`
				Number    string `xml:"idNumber"`
				Country   string `xml:"idCountry"`
				IssueDate string `xml:"issueDate"`
			} `xml:"id"`
		} `xml:"idList"`

		AddressList struct {
			Address []struct {
				UID             string `xml:"uid"`
				City            string `xml:"city"`
				Address1        string `xml:"address1"`
				Address2        string `xml:"address2"`
				Address3        string `xml:"address3"`
				StateOrProvince string `xml:"stateOrProvince"`
				PostalCode      string `xml:"postalCode"`
				Country         string `xml:"country"`
			} `xml:"address"`
		} `xml:"addressList"`

		NationalityList struct {
			Nationality []struct {
				UID       string `xml:"uid"`
				Country   string `xml:"country"`
				MainEntry bool   `xml:"mainEntry"`
			} `xml:"nationality"`
		} `xml:"nationalityList"`

		DateOfBirthList struct {
			DateOfBirthItem []struct {
				UID         string `xml:"uid"`
				DateOfBirth string `xml:"dateOfBirth"`
				MainEntry   bool   `xml:"mainEntry"`
			} `xml:"dateOfBirthItem"`
		} `xml:"dateOfBirthList"`

		PlaceOfBirthList struct {
			PlaceOfBirthItem []struct {
				UID       string `xml:"uid"`
				Place     string `xml:"place"`
				MainEntry bool   `xml:"mainEntry"`
			} `xml:"placeOfBirthItem"`
		} `xml:"placeOfBirthList"`

		CitizenshipList struct {
			Citizenship []struct {
				UID       string `xml:"uid"`
				Country   string `xml:"country"`
				MainEntry bool   `xml:"mainEntry"`
			} `xml:"citizenship"`
		} `xml:"citizenshipList"`
	} `xml:"sdnEntry"`
}
