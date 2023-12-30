package parser

import (
	"Remote_XML_Parser/internal/http"
	"Remote_XML_Parser/internal/services"
	"encoding/xml"
	"os"
	"time"
)

func ParseRemoteXML(remoteUrl string) (*SDN, error) {
	currentDirectory := "./server/xml/" + time.Now().Format("2006-01-02") + "/"
	os.Mkdir(currentDirectory, 0755)
	remoteXMLBytes, err := http.Download(remoteUrl, currentDirectory, "xml")
	if err != nil {
		return nil, err
	}
	sdnData := new(SDN)
	err = xml.Unmarshal(remoteXMLBytes, sdnData)
	if err != nil {
		return nil, services.ServerUnavailable
	}
	return sdnData, nil
}
