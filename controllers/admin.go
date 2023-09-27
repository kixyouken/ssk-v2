package controllers

import (
	"ssk-v2/services"

	"github.com/gin-gonic/gin"
)

func GetTable(c *gin.Context) {
	table := c.Param("table")
	tableJson := services.TableServices.GetTableFile(c, table)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    tableJson,
	})
}

func GetForm(c *gin.Context) {
	form := c.Param("form")
	formJson := services.FormServices.GetForm(c, form)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    formJson,
	})
}

func GetModel(c *gin.Context) {
	model := c.Param("model")
	modelJson := services.ModelServices.GetModelFile(c, model)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    modelJson,
	})
}
