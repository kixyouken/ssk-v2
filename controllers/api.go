package controllers

import (
	"ssk-v2/services"
	"strconv"

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

	if modelJson.Joins != nil && len(modelJson.Joins) > 0 {
		modelJoinsColumns := services.ModelServices.GetModelJoinsColumns(c, *modelJson)
		columns = append(columns, modelJoinsColumns...)
	}

	if modelJson.JoinsCount != nil && len(modelJson.JoinsCount) > 0 {
		modelJoinsCountColumns := services.ModelServices.GetModelJoinsCountColumns(c, *modelJson)
		columns = append(columns, modelJoinsCountColumns...)
	}

	if tableJson.Joins != nil && len(tableJson.Joins) > 0 {
		tableJoins := services.TableServices.GetTableJoins(c, *tableJson)
		joins = append(joins, tableJoins...)

		tableJoinsColumns := services.TableServices.GetTableJoinsColumns(c, *tableJson)
		columns = append(columns, tableJoinsColumns...)
	}

	if modelJson.JoinsCount != nil && len(modelJson.JoinsCount) > 0 {
		modelJoinsCount := services.ModelServices.GetModelJoinsCount(c, *modelJson)
		joins = append(joins, modelJoinsCount...)
		modelJoinsCountOrders := services.ModelServices.GetModelJoinsCountOrders(c, *modelJson)
		if modelJoinsCountOrders != "" {
			orders = modelJoinsCountOrders + ", " + orders
		}
	}

	if tableJson.Orders != nil && len(tableJson.Orders) > 0 {
		orders += ", " + services.TableServices.GetTableOrders(c, *tableJson)
	}

	result := []map[string]interface{}{}
	count := services.DbService.Count(c, modelJson.Table, joins)
	if count > 0 {
		services.DbService.Page(c, modelJson.Table, &result, columns, orders, joins, groups)
	}

	if modelJson.Withs != nil && len(modelJson.Withs) > 0 {
		services.ResultServices.HandleModelWithsList(c, result, *modelJson)
	}

	if tableJson.Withs != nil && len(tableJson.Withs) > 0 {
		services.ResultServices.HandleTableWithsList(c, result, *tableJson)
	}

	if modelJson.WithsCount != nil && len(modelJson.WithsCount) > 0 {
		services.ResultServices.HandleModelWithsCountList(c, result, *modelJson)
	}

	if modelJson.WithsSum != nil && len(modelJson.WithsSum) > 0 {
		services.ResultServices.HandleModelWithsSumList(c, result, *modelJson)
	}

	if tableJson.WithsCount != nil && len(tableJson.WithsCount) > 0 {
		services.ResultServices.HandleTableWithsCountList(c, result, *tableJson)
	}

	if tableJson.WithsSum != nil && len(tableJson.WithsSum) > 0 {
		services.ResultServices.HandleTableWithsSumList(c, result, *tableJson)
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

	if modelJson.WithsCount != nil && len(modelJson.WithsCount) > 0 {
		services.ResultServices.HandleModelWithsCount(c, result, *modelJson)
	}

	if modelJson.WithsSum != nil && len(modelJson.WithsSum) > 0 {
		services.ResultServices.HandleModelWithsSum(c, result, *modelJson)
	}

	if formJson.Withs != nil && len(formJson.Withs) > 0 {
		services.ResultServices.HandleFormWiths(c, result, *formJson)
	}

	if formJson.WithsCount != nil && len(formJson.WithsCount) > 0 {
		services.ResultServices.HandleFormWithsCount(c, result, *formJson)
	}

	if formJson.WithsSum != nil && len(formJson.WithsSum) > 0 {
		services.ResultServices.HandleFormWithsSum(c, result, *formJson)
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
