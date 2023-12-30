package http

import (
	"Remote_XML_Parser/internal/services"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

// Download returns filepath where file was loaded
func Download(remoteUrl string, prefixPath string, extension string) (string, error) {
	fileNameId := uuid.New().String()
	filePath := prefixPath + fileNameId + "." + extension
	out, err := os.Create(filePath)
	if err != nil {
		return "", services.ServerUnavailable
	}
	defer func() {
		err := out.Close()
		if err != nil {
			err = services.ServerUnavailable
		}
	}()
	if err != nil {
		return "", err
	}
	resp, err := http.Get(remoteUrl)
	if err != nil {
		return "", services.ServerUnavailable
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			err = services.ServerUnavailable
		}
	}(resp.Body)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", services.ServerUnavailable
	}
	return filePath, nil
}
