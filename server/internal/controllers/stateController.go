package controllers

import (
	"Remote_XML_Parser/internal/models/global"
	"Remote_XML_Parser/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatus(c *gin.Context, config *services.Config) {
	status := global.Empty
	var lastStatus global.UpdateStatus
	if err := config.DBClient.Order("created_at desc").Limit(1).Find(&lastStatus).Error; err == nil && lastStatus.Status != "" {
		status = lastStatus.Status
	}
	result := false
	if status == global.Ok {
		result = true
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"info":   status,
	})
}
