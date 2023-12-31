package xmls

import (
	"encoding/xml"
)

type SDN struct {
	SDNList            xml.Name `xml:"sdnList"`
	PublishInformation struct {
		PublishDate string `xml:"Publish_Date"`
		RecordCount int    `xml:"Record_Count"`
	} `xml:"publshInformation"`
	SDNEntry []SDNItem `xml:"sdnEntry"`
}
