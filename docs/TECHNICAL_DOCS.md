# Remote XML Parser

Импорт / обновление необходимых данных из https://www.treasury.gov/ofac/downloads/sdn.xml в
локальную базу PostgreSQL 14. Получение текущего состояния данных.
Получение списка всех возможных имён человека из локальной базы данных с указанием
основного uid в виде JSON.

[Об оптимизации скорости выполнения запроса ```POST /update```](./redis/README.md)

### Deployment

```bash 
    cp ./database/.env.example ./database/.env
    cp ./redis/.env.example ./redis/.env
    cp ./server/.env.example ./server/.env
    docker compose build && docker compose up -d
```

### Общая XML структура имеет следующий вид:
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
				FirstName xml.Name `xmlmodel:"firstName"`
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