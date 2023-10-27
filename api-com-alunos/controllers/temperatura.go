package controllers

import (
	"api/database"
	"api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTemperatura(c *gin.Context) {
	var temperatura models.Temperatura
	if err := c.ShouldBindJSON(&temperatura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	temperatura.Data = time.Now()

	database.DB.Create(&temperatura)
	c.JSON(http.StatusCreated, temperatura)
}

func GetTemperatura(c *gin.Context) {
	var temperaturas []models.Temperatura
	database.DB.Find(&temperaturas)

	var respose []map[string]interface{}
	for _, temperatura := range temperaturas {
		dataFormatada := temperatura.Data.Format("2005-01-02")

		temperaturaMap := map[string]interface{}{
			"temperatura": temperatura.Temperatura,
			"data":        dataFormatada,
		}

		respose = append(respose, temperaturaMap)

	}

	c.JSON(http.StatusOK, respose)
}
