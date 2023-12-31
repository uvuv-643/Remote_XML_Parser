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
	SDNList            xml.Name `xmlmodel:"sdnList"`
	PublishInformation struct {
		PublishDate string `xmlmodel:"Publish_Date"`
		RecordCount int    `xmlmodel:"Record_Count"`
	} `xmlmodel:"publshInformation"`

	SDNEntry []struct {
		UID       string `xmlmodel:"uid"`
		FirstName string `xmlmodel:"firstName"`
		LastName  string `xmlmodel:"lastName"`
		Title     string `xmlmodel:"title"`
		SDNType   string `xmlmodel:"sdnType"`
		Remarks   string `xmlmodel:"remarks"`

		ProgramList struct {
			Program []string `xmlmodel:"program"`
		} `xmlmodel:"programList"`

		AkaList struct {
			Aka []struct {
				UID      string   `xmlmodel:"uid"`
				Type     string   `xmlmodel:"type"`
				Category string   `xmlmodel:"category"`
				LastName xml.Name `xmlmodel:"lastName"`
			} `xmlmodel:"aka"`
		} `xmlmodel:"akaList"`

		IdList struct {
			ID []struct {
				UID       string `xmlmodel:"uid"`
				Type      string `xmlmodel:"idType"`
				Number    string `xmlmodel:"idNumber"`
				Country   string `xmlmodel:"idCountry"`
				IssueDate string `xmlmodel:"issueDate"`
			} `xmlmodel:"id"`
		} `xmlmodel:"idList"`

		AddressList struct {
			Address []struct {
				UID             string `xmlmodel:"uid"`
				City            string `xmlmodel:"city"`
				Address1        string `xmlmodel:"address1"`
				Address2        string `xmlmodel:"address2"`
				Address3        string `xmlmodel:"address3"`
				StateOrProvince string `xmlmodel:"stateOrProvince"`
				PostalCode      string `xmlmodel:"postalCode"`
				Country         string `xmlmodel:"country"`
			} `xmlmodel:"address"`
		} `xmlmodel:"addressList"`

		NationalityList struct {
			Nationality []struct {
				UID       string `xmlmodel:"uid"`
				Country   string `xmlmodel:"country"`
				MainEntry bool   `xmlmodel:"mainEntry"`
			} `xmlmodel:"nationality"`
		} `xmlmodel:"nationalityList"`

		DateOfBirthList struct {
			DateOfBirthItem []struct {
				UID         string `xmlmodel:"uid"`
				DateOfBirth string `xmlmodel:"dateOfBirth"`
				MainEntry   bool   `xmlmodel:"mainEntry"`
			} `xmlmodel:"dateOfBirthItem"`
		} `xmlmodel:"dateOfBirthList"`

		PlaceOfBirthList struct {
			PlaceOfBirthItem []struct {
				UID       string `xmlmodel:"uid"`
				Place     string `xmlmodel:"place"`
				MainEntry bool   `xmlmodel:"mainEntry"`
			} `xmlmodel:"placeOfBirthItem"`
		} `xmlmodel:"placeOfBirthList"`

		CitizenshipList struct {
			Citizenship []struct {
				UID       string `xmlmodel:"uid"`
				Country   string `xmlmodel:"country"`
				MainEntry bool   `xmlmodel:"mainEntry"`
			} `xmlmodel:"citizenship"`
		} `xmlmodel:"citizenshipList"`
	} `xmlmodel:"sdnEntry"`
}


```