package controllers

import (
	"Remote_XML_Parser/internal/dto"
	"Remote_XML_Parser/internal/repository"
	"Remote_XML_Parser/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

const WEAK = "weak"
const STRONG = "strong"

func GetNames(c *gin.Context, config *services.Config) {
	queryParams := c.Request.URL.Query()
	if len(queryParams["name"]) > 0 {
		actualName := queryParams["name"][0]
		nameParts := strings.Split(actualName, " ")
		if len(nameParts) > 2 {
			log.Println("GET /get_name: strange name", nameParts)
			nameParts = nameParts[0:1]
		}
		nameRepository := repository.NewNameRepository()
		if len(queryParams["type"]) == 1 {
			targetType := strings.ToLower(queryParams["type"][0])
			if targetType == WEAK {

				result, err := nameRepository.Weak(config.DBClient, nameParts...)
				if err != nil {
					c.JSON(503, result)
					return
				}
				if len(result) == 0 {
					c.JSON(200, []string{})
					return
				}
				c.JSON(200, result)
				return

			} else if targetType == STRONG {

				result, err := nameRepository.Strong(config.DBClient, nameParts...)
				if err != nil {
					c.JSON(503, result)
					return
				}
				if len(result) == 0 {
					c.JSON(200, []string{})
					return
				}
				c.JSON(200, result)
				return

			}
		}

		resultWeak, errWeak := nameRepository.Weak(config.DBClient, nameParts...)
		resultStrong, errStrong := nameRepository.Strong(config.DBClient, nameParts...)
		if errWeak != nil || errStrong != nil {
			c.JSON(503, nil)
			return
		}

		// to remove duplicates by dto.UserName.RecordKey
		resultByKeyValue := make(map[string]dto.UserName)
		for _, weakElem := range resultWeak {
			resultByKeyValue[weakElem.RecordKey] = weakElem
		}
		for _, strongElem := range resultStrong {
			resultByKeyValue[strongElem.RecordKey] = strongElem
		}
		result := make([]dto.UserName, 0)
		for _, resultElem := range resultByKeyValue {
			result = append(result, resultElem)
		}

		if len(result) == 0 {
			c.JSON(200, []string{})
			return
		}
		c.JSON(200, result)

	} else {
		c.JSON(400, gin.H{})
	}
}
