package routes

import (
	"dynamicWeb/config"
	"dynamicWeb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupConfigRoutes(router *gin.Engine) {
	configGroup := router.Group("/api/configuration")
	{
		configGroup.POST("/", func(c *gin.Context) {
			var newConfig models.Configuration
			if err := c.ShouldBindJSON(&newConfig); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			newConfig.ID = "" // Generate new UUID
			savedConfig := models.NewConfiguration(newConfig.Actions)
			if err := config.SaveConfiguration(savedConfig); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, savedConfig)
		})

		configGroup.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			configData, err := config.GetConfiguration(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Configuration not found"})
				return
			}

			c.JSON(http.StatusOK, configData)
		})

		configGroup.GET("/all", func(c *gin.Context) {
			configs, err := config.GetAllConfigurations()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, configs)
		})

		configGroup.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			var updatedConfig models.Configuration
			if err := c.ShouldBindJSON(&updatedConfig); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := config.UpdateConfiguration(id, &updatedConfig); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Configuration not found"})
				return
			}

			c.JSON(http.StatusOK, updatedConfig)
		})

		configGroup.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if err := config.DeleteConfiguration(id); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Configuration not found"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Configuration deleted"})
		})
	}
}
