package controllers

import (
	"Remote_XML_Parser/internal/models"
	"Remote_XML_Parser/internal/models/dbs"
	"Remote_XML_Parser/internal/models/global"
	"Remote_XML_Parser/internal/models/xmls"
	"Remote_XML_Parser/internal/parser"
	"Remote_XML_Parser/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

func UpdateHandler(c *gin.Context, config *services.Config) {
	config.PGClient.Create(&global.UpdateStatus{Status: global.Updating})
	sdn, err := parser.ParseRemoteXML(config.XMLRemoteURL)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"code":    http.StatusServiceUnavailable,
			"info":    err.Error(),
		})
	} else {
		sdnItems := make([]xmls.SDNItem, 0)
		for _, sdnItem := range sdn.SDNEntry {
			if sdnItem.SDNType == "Individual" {
				sdnItems = append(sdnItems, sdnItem)
			}
		}

		uniquePrograms := make(map[xmls.SDNProgram]bool)
		uniqueAkas := make(map[xmls.SDNAka]bool)
		uniqueIds := make(map[xmls.SDNId]bool)
		uniqueAddresses := make(map[xmls.SDNAddress]bool)
		uniqueNationalities := make(map[xmls.SDNNationality]bool)
		uniqueDatesOfBirth := make(map[xmls.SDNDateOfBirth]bool)
		uniquePlacesOfBirth := make(map[xmls.SDNPlaceOfBirth]bool)
		uniqueCitizenships := make(map[xmls.SDNCitizenship]bool)

		// todo: redis
		for _, sdnItem := range sdnItems {
			for _, program := range sdnItem.ProgramList.Program {
				uniquePrograms[program] = true
			}
			for _, aka := range sdnItem.AkaList.Aka {
				uniqueAkas[aka] = true
			}
			for _, id := range sdnItem.IdList.ID {
				uniqueIds[id] = true
			}
			for _, address := range sdnItem.AddressList.Address {
				uniqueAddresses[address] = true
			}
			for _, nationality := range sdnItem.NationalityList.Nationality {
				uniqueNationalities[nationality] = true
			}
			for _, date := range sdnItem.DateOfBirthList.DateOfBirthItem {
				uniqueDatesOfBirth[date] = true
			}
			for _, place := range sdnItem.PlaceOfBirthList.PlaceOfBirthItem {
				uniquePlacesOfBirth[place] = true
			}
			for _, citizenship := range sdnItem.CitizenshipList.Citizenship {
				uniqueCitizenships[citizenship] = true
			}
		}

		queriedProgramList := make([]dbs.SDNProgram, 0)
		queriedAkaList := make([]dbs.SDNAka, 0)
		queriedIdList := make([]dbs.SDNId, 0)
		queriedAddressList := make([]dbs.SDNAddress, 0)
		queriedNationalityList := make([]dbs.SDNNationality, 0)
		queriedDateList := make([]dbs.SDNDateOfBirth, 0)
		queriedPlaceList := make([]dbs.SDNPlaceOfBirth, 0)
		queriedCitizenshipList := make([]dbs.SDNCitizenship, 0)

		for entity, state := range uniquePrograms {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNProgram)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(entityConverted).String() + "." + entityConverted.Program
					if !config.CacheHit(redisKey, *entityConverted) {
						queriedProgramList = append(queriedProgramList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueAkas {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNAka)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheHit(redisKey, *entityConverted) {
						queriedAkaList = append(queriedAkaList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueIds {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNId)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheHit(redisKey, *entityConverted) {
						queriedIdList = append(queriedIdList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueAddresses {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNAddress)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheHit(redisKey, *entityConverted) {
						queriedAddressList = append(queriedAddressList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueNationalities {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNNationality)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheHit(redisKey, *entityConverted) {
						queriedNationalityList = append(queriedNationalityList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueDatesOfBirth {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNDateOfBirth)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheHit(redisKey, *entityConverted) {
						queriedDateList = append(queriedDateList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniquePlacesOfBirth {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNPlaceOfBirth)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheHit(redisKey, *entityConverted) {
						queriedPlaceList = append(queriedPlaceList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueCitizenships {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNCitizenship)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheHit(redisKey, *entityConverted) {
						queriedCitizenshipList = append(queriedCitizenshipList, *entityConverted)
					}
				}
			}
		}

		// todo: database insert in batches

		config.PGClient.Create(&global.UpdateStatus{Status: global.Ok})
		c.JSON(http.StatusOK, gin.H{
			"result": true,
			"info":   "",
			"code":   200,
		})
	}
}
