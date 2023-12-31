package parser

import (
	"Remote_XML_Parser/internal/http"
	"Remote_XML_Parser/internal/models/xmls"
	"Remote_XML_Parser/internal/services"
	"encoding/xml"
	"log"
	"os"
	"time"
)

func ParseRemoteXML(remoteUrl string) (*xmls.SDN, error) {
	currentDirectory := "/xml/" + time.Now().Format("2006-01-02") + "/"
	os.Mkdir(currentDirectory, 0755)
	remoteXMLBytes, err := http.Download(remoteUrl, currentDirectory, "xml")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	sdnData := new(xmls.SDN)
	err = xml.Unmarshal(remoteXMLBytes, sdnData)
	if err != nil {
		log.Println(err.Error())
		return nil, services.ServerUnavailable
	}
	return sdnData, nil
}
