package controllers

import (
	"ssk-v2/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Get 获取所有数据
//
//	@param c
func Get(c *gin.Context) {
	model := c.Param("model")
	result := []map[string]interface{}{}
	services.DbService.Get(c, model, &result, "*", "", []string{}, "")

	c.JSON(200, gin.H{
		"message": "Get",
		"data":    result,
	})
}

// Page 获取列表数据
//
//	@param c
func Page(c *gin.Context) {
	table := c.Param("table")
	tableJson := services.TableServices.GetTableFile(c, table)
	modelJson := services.ModelServices.GetModelFile(c, tableJson.Model)
	columns := services.ModelServices.GetModelColumns(c, *modelJson)
	orders := services.ModelServices.GetModelOrders(c, *modelJson)
	joins := services.ModelServices.GetModelJoins(c, *modelJson)
	groups := services.ModelServices.GetModelJoinsCountGroup(c, *modelJson)

	modelBeforeColumns, modelBeforeJoins, modelBeforeOrders := services.ModelServices.GetModelFileQueryBefore(c, *modelJson)
	columns = append(columns, modelBeforeColumns...)
	joins = append(joins, modelBeforeJoins...)
	if modelBeforeOrders != "" {
		orders = strings.TrimRight(modelBeforeOrders+","+orders, ",")
	}

	tableBeforeColumns, tableBeforeJoins, tableBeforeOrders := services.TableServices.GetTableFileQueryBefore(c, *tableJson)
	columns = append(columns, tableBeforeColumns...)
	joins = append(joins, tableBeforeJoins...)
	if tableBeforeOrders != "" {
		orders += "," + tableBeforeOrders
	}

	result := []map[string]interface{}{}
	count := services.DbService.Count(c, modelJson.Table, joins)
	if count > 0 {
		services.DbService.Page(c, modelJson.Table, &result, columns, orders, joins, groups)
	}

	services.TableServices.GetTableFileQueryAfterList(c, result, *tableJson)
	services.ModelServices.GetModelFileQueryAfterList(c, result, *modelJson)

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

	services.FormServices.GetFormFileQueryAfter(c, result, *formJson)
	services.ModelServices.GetModelFileQueryAfter(c, result, *modelJson)

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
