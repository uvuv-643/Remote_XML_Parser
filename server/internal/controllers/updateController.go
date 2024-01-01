package controllers

import (
	"Remote_XML_Parser/internal/models"
	"Remote_XML_Parser/internal/models/dbs"
	"Remote_XML_Parser/internal/models/global"
	"Remote_XML_Parser/internal/models/xmls"
	"Remote_XML_Parser/internal/parser"
	"Remote_XML_Parser/internal/services"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

func serverUnavailableUpdate(c *gin.Context, err error, config *services.Config) {
	// database error -> data is not consistent
	// todo: make optimization only for broken data
	config.RedisClient.FlushAll()
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"success": false,
		"code":    http.StatusServiceUnavailable,
		"info":    err.Error(),
	})
}

func UpdateHandler(c *gin.Context, config *services.Config) {
	config.DBClient.Create(&global.UpdateStatus{Status: global.Updating})
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

		// delete removed from remote xml
		sdnItemsIdsSet := make(map[int64]bool)
		for _, sdnRemoteItem := range sdnItems {
			sdnItemsIdsSet[sdnRemoteItem.UID] = true
		}

		redisItemsIdsKey := reflect.TypeOf(sdnItems[0]).String()
		redisResponse := config.RedisClient.Get(redisItemsIdsKey)
		var itemIdsInDatabase []int64
		if result, err := redisResponse.Result(); !errors.Is(err, redis.Nil) {
			err = json.Unmarshal([]byte(result), &itemIdsInDatabase)
			if err != nil {
				config.RedisClient.Del(redisItemsIdsKey)
			}
		} else {
			config.RedisClient.Del(redisItemsIdsKey)
		}

		if len(itemIdsInDatabase) == 0 {
			if err := config.DBClient.Model(&dbs.SDNItem{}).Pluck("uid", &itemIdsInDatabase).Error; err != nil {
				serverUnavailableUpdate(c, errors.New("cannot parse removed elements"), config)
			}
		}
		for _, itemId := range itemIdsInDatabase {
			if !sdnItemsIdsSet[itemId] {
				// already in database but not in remote XML
				if err := config.DBClient.Delete(&dbs.SDNItem{}, itemId).Error; err != nil {
					serverUnavailableUpdate(c, errors.New("cannot delete element from database"), config)
				}
			}
		}
		itemIdsActualInDatabase := make([]int64, 0)
		for key, state := range sdnItemsIdsSet {
			if state {
				itemIdsActualInDatabase = append(itemIdsActualInDatabase, key)
			}
		}
		jsonEncodedDatabaseElements, err := json.Marshal(itemIdsActualInDatabase)
		if err != nil {
			serverUnavailableUpdate(c, errors.New(""), config)
		}
		config.RedisClient.Set(redisItemsIdsKey, string(jsonEncodedDatabaseElements), time.Duration(int(time.Second)*config.RedisTTL))

		uniquePrograms := make(map[xmls.SDNProgram]bool)
		uniqueAkas := make(map[xmls.SDNAka]bool)
		uniqueIds := make(map[xmls.SDNId]bool)
		uniqueAddresses := make(map[xmls.SDNAddress]bool)
		uniqueNationalities := make(map[xmls.SDNNationality]bool)
		uniqueDatesOfBirth := make(map[xmls.SDNDateOfBirth]bool)
		uniquePlacesOfBirth := make(map[xmls.SDNPlaceOfBirth]bool)
		uniqueCitizenships := make(map[xmls.SDNCitizenship]bool)

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
		queriedItemsList := make([]dbs.SDNItem, 0)

		for entity, state := range uniquePrograms {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNProgram)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(*entityConverted).String() + "." + entityConverted.Program
					if !config.CacheGetOrSet(redisKey, *entityConverted) {
						queriedProgramList = append(queriedProgramList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueAkas {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNAka)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(*entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheGetOrSet(redisKey, *entityConverted) {
						queriedAkaList = append(queriedAkaList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueIds {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNId)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(*entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheGetOrSet(redisKey, *entityConverted) {
						queriedIdList = append(queriedIdList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueAddresses {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNAddress)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(*entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheGetOrSet(redisKey, *entityConverted) {
						queriedAddressList = append(queriedAddressList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueNationalities {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNNationality)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(*entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheGetOrSet(redisKey, *entityConverted) {
						queriedNationalityList = append(queriedNationalityList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniqueDatesOfBirth {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNDateOfBirth)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(*entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheGetOrSet(redisKey, *entityConverted) {
						queriedDateList = append(queriedDateList, *entityConverted)
					}
				}
			}
		}

		for entity, state := range uniquePlacesOfBirth {
			if state {
				entityConverted := models.ConvertXmlToDb(entity).(*dbs.SDNPlaceOfBirth)
				if entityConverted != nil {
					redisKey := reflect.TypeOf(*entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
					if !config.CacheGetOrSet(redisKey, *entityConverted) {
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
					if !config.CacheGetOrSet(redisKey, *entityConverted) {
						queriedCitizenshipList = append(queriedCitizenshipList, *entityConverted)
					}
				}
			}
		}

		for _, sdnItem := range sdnItems {
			entityConverted := models.ConvertXmlToDb(sdnItem).(*dbs.SDNItem)
			if entityConverted != nil {
				redisKey := reflect.TypeOf(*entityConverted).String() + "." + strconv.Itoa(int(entityConverted.UID))
				if !config.CacheGetOrSet(redisKey, *entityConverted) {
					queriedItemsList = append(queriedItemsList, *entityConverted)
				}
			}
		}

		for _, program := range queriedProgramList {
			if err := config.DBClient.First(&dbs.SDNProgram{}, "program = ?", program.Program).Updates(program).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&program).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}
		for _, aka := range queriedAkaList {
			if err := config.DBClient.First(&dbs.SDNAka{}, aka.UID).Updates(aka).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&aka).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}
		for _, id := range queriedIdList {
			if err := config.DBClient.First(&dbs.SDNId{}, id.UID).Updates(id).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&id).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}
		for _, address := range queriedAddressList {
			if err := config.DBClient.First(&dbs.SDNAddress{}, address.UID).Updates(address).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&address).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}
		for _, nationality := range queriedNationalityList {
			if err := config.DBClient.First(&dbs.SDNNationality{}, nationality.UID).Updates(nationality).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&nationality).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}
		for _, date := range queriedDateList {
			if err := config.DBClient.First(&dbs.SDNDateOfBirth{}, date.UID).Updates(date).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&date).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}
		for _, place := range queriedPlaceList {
			if err := config.DBClient.First(&dbs.SDNPlaceOfBirth{}, place.UID).Updates(place).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&place).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}
		for _, citizenship := range queriedCitizenshipList {
			if err := config.DBClient.First(&dbs.SDNCitizenship{}, citizenship.UID).Updates(citizenship).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&citizenship).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}

		fmt.Println("cache missed items when updating: ", len(queriedItemsList))
		for _, item := range queriedItemsList {
			if err := config.DBClient.First(&dbs.SDNItem{}, item.UID).Updates(item).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = config.DBClient.Create(&item).Error; err != nil {
						serverUnavailableUpdate(c, err, config)
					}
				}
			}
		}

		config.DBClient.Create(&global.UpdateStatus{Status: global.Ok})

		c.JSON(http.StatusOK, gin.H{
			"result": true,
			"info":   "",
			"code":   200,
		})
	}
}
