package routes

import (
	"Remote_XML_Parser/internal/controllers"
	"Remote_XML_Parser/internal/services"
	"Remote_XML_Parser/xml"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupDefaultEndpoints(r *gin.Engine, conf *services.Config) {
	r.GET("/tags", func(c *gin.Context) {
		xml.PrintTags()
		c.JSON(200, gin.H{
			"success": true,
		})
	})
}

func AddRoutes(r *gin.Engine, config *services.Config) {
	r.POST("/update", func(c *gin.Context) {
		signal := make(chan struct{}, 1)
		controllers.UpdateHandler(c, config, signal)
		select {
		case <-signal:
			fmt.Println("Request SIGEND")
			close(signal)
		case <-c.Request.Context().Done():
			config.RedisClient.FlushAll()
			return
		}
	})
	r.GET("/state", func(c *gin.Context) {
		controllers.GetStatus(c, config)
	})
	r.GET("/get_names", func(c *gin.Context) {
		controllers.GetNames(c, config)
	})
}
