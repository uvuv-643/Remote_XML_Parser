package controllers

import (
	"Remote_XML_Parser/internal/models/xmls"
	"Remote_XML_Parser/internal/parser"
	"Remote_XML_Parser/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateHandler(c *gin.Context, config *services.Config) {
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

		c.JSON(http.StatusOK, gin.H{
			"result": true,
			"info":   "",
			"code":   200,
		})
	}
}
