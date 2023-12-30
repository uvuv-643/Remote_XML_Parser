package parser

import (
	"Remote_XML_Parser/internal/http"
	"os"
	"time"
)

func ParseRemoteXML(remoteUrl string) (string, error) {
	currentDirectory := "./server/xml/" + time.Now().Format("2006-01-02") + "/"
	os.Mkdir(currentDirectory, 0755)
	return http.Download(remoteUrl, currentDirectory, "xml")
}
