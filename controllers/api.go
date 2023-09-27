package controllers

import (
	"ssk-v2/services"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	table := c.Param("model")
	tableJson := services.TableServices.GetTableFile(c, table)
	result := []map[string]interface{}{}
	services.DbService.Get(c, tableJson.Model, &result, "*", "", []string{}, "")
	c.JSON(200, gin.H{
		"message": "Get",
		"data":    result,
		"table":   tableJson,
	})
}

func Page(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Page",
	})
}

func Read(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Read",
	})
}

func Save(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Save",
	})
}

func Update(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Update",
	})
}

func Delete(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Delete",
	})
}
