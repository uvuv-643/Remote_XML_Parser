# Remote XML Parser

Доступные XML теги [исходной структуры](https://www.treasury.gov/ofac/downloads/sdn.xm):

* Publish_Date
* Record_Count
* address
* address1
* address2
* address3
* addressList
* aka
* akaList
* callSign ```Only for Vessel type``` 
* category
* citizenship
* citizenshipList
* city
* country
* dateOfBirth
* dateOfBirthItem
* dateOfBirthList
* expirationDate
* firstName
* grossRegisteredTonnage ```Only for Vessel type```
* id
* idCountry
* idList
* idNumber
* idType
* issueDate
* lastName
* mainEntry
* nationality
* nationalityList
* placeOfBirth
* placeOfBirthItem
* placeOfBirthList
* postalCode
* program
* programList
* publishInformation
* remarks
* sdnEntry
* sdnList
* sdnType
* stateOrProvince
* title
* tonnage ```Only for Vessel type```
* type
* uid
* vesselFlag ```Only for Vessel type```
* vesselInfo ```Only for Vessel type```
* vesselOwner ```Only for Vessel type```
* vesselType ```Only for Vessel type```

XML структура имеет следующий вид:
```go

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

		ProgramList xml.Name `xml:"programList"`
		Program     []string `xml:"program"`

		AkaList xml.Name `xml:"akaList"`
		Aka     []struct {
			UID      string   `xml:"uid"`
			Type     string   `xml:"type"`
			Category string   `xml:"category"`
			LastName xml.Name `xml:"lastName"`
		} `xml:"aka"`

		IdList xml.Name `xml:"idList"`
		ID     []struct {
			UID       string `xml:"uid"`
			Type      string `xml:"idType"`
			Number    string `xml:"idNumber"`
			Country   string `xml:"idCountry"`
			IssueDate string `xml:"issueDate"`
		} `xml:"id"`

		AddressList xml.Name `xml:"addressList"`
		Address     []struct {
			UID             string `xml:"uid"`
			City            string `xml:"city"`
			Address1        string `xml:"address1"`
			Address2        string `xml:"address2"`
			Address3        string `xml:"address3"`
			StateOrProvince string `xml:"stateOrProvince"`
			PostalCode      string `xml:"postalCode"`
			Country         string `xml:"country"`
		} `xml:"address"`

		NationalityList xml.Name `xml:"nationalityList"`
		Nationality     []struct {
			UID       string `xml:"uid"`
			Country   string `xml:"country"`
			MainEntry bool   `xml:"mainEntry"`
		} `xml:"nationality"`

		DateOfBirthList xml.Name `xml:"dateOfBirthList"`
		DateOfBirthItem []struct {
			UID         string `xml:"uid"`
			DateOfBirth string `xml:"dateOfBirth"`
			MainEntry   bool   `xml:"mainEntry"`
		} `xml:"dateOfBirthItem"`

		PlaceOfBirthList xml.Name `xml:"placeOfBirthList"`
		PlaceOfBirthItem []struct {
			UID       string `xml:"uid"`
			Place     string `xml:"place"`
			MainEntry bool   `xml:"mainEntry"`
		} `xml:"placeOfBirthItem"`

		CitizenshipList xml.Name `xml:"citizenshipList"`
		Citizenship     []struct {
			UID       string `xml:"uid"`
			Country   string `xml:"country"`
			MainEntry bool   `xml:"mainEntry"`
		} `xml:"citizenship"`
	} `xml:"sdnEntry"`
}


```