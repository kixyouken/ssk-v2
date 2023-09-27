package controllers

import (
	"ssk-v2/services"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	table := c.Param("model")
	tableJson := services.TableServices.GetTableFile(c, table)
	modelJson := services.ModelServices.GetModelFile(c, tableJson.Model)
	columns := services.ModelServices.GetModelColumns(c, *modelJson)
	orders := services.TableServices.GetTableOrders(c, *tableJson)

	var count int64
	var result []map[string]interface{}
	if tableJson.Page == "true" {
		count = services.DbService.Count(c, modelJson.Table, []string{}, "")
		if count > 0 {
			services.DbService.Page(c, modelJson.Table, &result, columns, orders, []string{}, "")
		}
	} else {
		services.DbService.Get(c, modelJson.Table, &result, columns, orders, []string{}, "")
	}

	c.JSON(200, gin.H{
		"message": "Get",
		"data":    result,
		"count":   count,
		"table":   tableJson,
		"orders":  orders,
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
