package xml

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func getRandomXMLFile(directory string) (string, error) {
	var result string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) == ".xml" {
			result = fmt.Sprintf("%s/%s", filepath.Base(filepath.Dir(path)), info.Name())
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if result == "" {
		return "", fmt.Errorf("no XML files found in the directory or its subdirectories")
	}
	return result, nil
}
func getXMLUniqueTags(filePath string) []string {
	uniqueTags := make(map[string]bool)
	file, err := os.Open(filePath)
	if err != nil {
		return make([]string, 0)
	}
	defer file.Close()
	response := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		lineText = strings.Trim(lineText, " ")
		if lineText[0] == '<' && lineText[1] != '/' {
			re := regexp.MustCompile("[ >]")
			lineTextSplitted := re.Split(lineText, -1)
			uniqueTags[lineTextSplitted[0][1:]] = true
		}
	}
	for key, value := range uniqueTags {
		if value {
			response = append(response, key)
		}
	}
	if err := scanner.Err(); err != nil {
		return make([]string, 0)
	}
	return response
}

func PrintTags() {
	directory := "/xml/"
	randomXMLFile, err := getRandomXMLFile(directory)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Random XML File: %s\n", randomXMLFile)
	uniqueTags := getXMLUniqueTags(directory + randomXMLFile)
	sort.Strings(uniqueTags)
	for _, tag := range uniqueTags {
		fmt.Println(tag)
	}
}
