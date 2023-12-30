package http

import (
	"Remote_XML_Parser/internal/services"
	"bytes"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

// Download returns filepath where file was loaded
func Download(remoteUrl string, prefixPath string, extension string) ([]byte, error) {

	var emptyResponse = []byte("")
	fileNameId := uuid.New().String()
	filePath := prefixPath + fileNameId + "." + extension

	// create new .xml file
	out, err := os.Create(filePath)
	if err != nil {
		return emptyResponse, services.ServerUnavailable
	}
	defer func() {
		out.Close()
	}()

	// get remote .xml file
	resp, err := http.Get(remoteUrl)
	if err != nil {
		return emptyResponse, services.ServerUnavailable
	}
	defer func(Body io.ReadCloser) {
		Body.Close()
	}(resp.Body)

	// read remote .xml file content to memory
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyResponse, services.ServerUnavailable
	}

	// write from memory to drive
	_, err = io.Copy(out, bytes.NewBuffer(content))
	if err != nil {
		return emptyResponse, services.ServerUnavailable
	}

	// return content
	return content, nil
	
}
