package controllers

import (
	"ssk-v2/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get 获取列表数据
//
//	@param c
func Get(c *gin.Context) {
	table := c.Param("table")
	tableJson := services.TableServices.GetTableFile(c, table)
	modelJson := services.ModelServices.GetModelFile(c, tableJson.Model)
	columns := services.ModelServices.GetModelColumns(c, *modelJson)
	orders := services.TableServices.GetTableOrders(c, *tableJson)
	joins := services.ModelServices.GetModelJoins(c, *modelJson)

	var count int64
	var result []map[string]interface{}
	if tableJson.Paginate == "true" {
		count = services.DbService.Count(c, modelJson.Table, joins)
		if count > 0 {
			services.DbService.Page(c, modelJson.Table, &result, columns, orders, joins)
		}
	} else {
		services.DbService.Get(c, modelJson.Table, &result, columns, orders, joins, "")
		count = int64(len(result))
	}

	if modelJson.Withs != nil && len(modelJson.Withs) > 0 {
		services.ResultServices.HandleModelWithsList(c, result, *modelJson)
	}

	if modelJson.Columns != nil && len(modelJson.Columns) > 0 {
		services.ResultServices.HandleModelFieldFormatList(c, result, *modelJson)
	}

	c.JSON(200, gin.H{
		"message": "Get",
		"count":   count,
		"data":    result,
	})
}

// Read 获取详情
//
//	@param c
func Read(c *gin.Context) {
	form := c.Param("form")
	formJson := services.FormServices.GetForm(c, form)
	modelJson := services.ModelServices.GetModelFile(c, formJson.Model)
	columns := services.ModelServices.GetModelColumns(c, *modelJson)

	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	result := map[string]interface{}{}
	services.DbService.Read(c, modelJson.Table, idInt, &result, columns)

	if modelJson.Withs != nil && len(modelJson.Withs) > 0 {
		services.ResultServices.HandleModelWiths(c, result, *modelJson)
	}

	if formJson.Withs != nil && len(formJson.Withs) > 0 {
		services.ResultServices.HandleFormWiths(c, result, *formJson)
	}

	if modelJson.Columns != nil && len(modelJson.Columns) > 0 {
		services.ResultServices.HandleModelFieldFormat(c, result, *modelJson)
	}
	c.JSON(200, gin.H{
		"message": "Read",
		"data":    result,
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
