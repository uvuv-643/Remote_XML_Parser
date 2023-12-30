package controllers

import (
	"Remote_XML_Parser/internal/parser"
	"Remote_XML_Parser/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateHandler(c *gin.Context, config *services.Config) {
	code, err := parser.ParseRemoteXML(config.XMLRemoteURL)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"code":    http.StatusServiceUnavailable,
			"info":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
		})
	}
}
