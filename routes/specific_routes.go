package routes

import (
	"dynamicWeb/models"
	config "dynamicWeb/specific"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupSpecificRoutes(router *gin.Engine) {
	configGroup := router.Group("/api/specific")
	{
		configGroup.POST("/", func(c *gin.Context) {
			var newConfig models.Specific
			if err := c.ShouldBindJSON(&newConfig); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			newConfig.ID = "" // Generate new UUID
			savedConfig := models.NewSpecific(newConfig.Datasource)
			if err := config.SaveSpecific(savedConfig); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, savedConfig)
		})

		configGroup.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")

			host := c.Query("host")
			url := c.Query("url")
			page := c.Query("page")

			configData, err := config.GetSpecific(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Specific Configuration not found"})
				return
			}

			var matchedConfigs []string
			if val, ok := configData.Datasource.Pages[page]; ok {
				matchedConfigs = append(matchedConfigs, val...)
			}
			if val, ok := configData.Datasource.Urls[url]; ok {
				matchedConfigs = append(matchedConfigs, val...)
			}
			if val, ok := configData.Datasource.Hosts[host]; ok {
				matchedConfigs = append(matchedConfigs, val...)
			}

			if matchedConfigs != nil {
				c.JSON(http.StatusOK, gin.H{"matchedConfigs": matchedConfigs})
				return
			}

			c.JSON(http.StatusOK, configData)

		})

		configGroup.GET("/all", func(c *gin.Context) {
			configs, err := config.GetAllSpecifics()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, configs)
		})

		configGroup.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			var updatedConfig models.Specific
			if err := c.ShouldBindJSON(&updatedConfig); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := config.UpdateSpecific(id, &updatedConfig); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Specific Configuration not found"})
				return
			}

			c.JSON(http.StatusOK, updatedConfig)
		})

		configGroup.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if err := config.DeleteSpecific(id); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Specific Configuration not found"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Specific Configuration deleted"})
		})
	}
}
