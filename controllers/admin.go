package controllers

import (
	"ssk-v2/services"

	"github.com/gin-gonic/gin"
)

func GetModel(c *gin.Context) {
	model := c.Param("model")
	modelJson := services.ModelServices.GetModelFile(c, model)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    modelJson,
	})
}
