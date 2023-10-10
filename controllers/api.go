package controllers

import (
	"ssk-v2/services"
	"strconv"
	"strings"
	"time"

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

	groups := []string{}
	modelJoinsGroupsJoins, modelJoinsGroupsColumns, modelJoinsGroupsGroups, modelJoinsGroupsOrders := services.ModelServices.GetModelJoinsGroups(c, *modelJson)
	joins = append(joins, modelJoinsGroupsJoins...)
	columns = append(columns, modelJoinsGroupsColumns...)
	groups = append(groups, modelJoinsGroupsGroups...)
	if modelJoinsGroupsOrders != "" {
		orders = strings.Trim(modelJoinsGroupsOrders+","+orders, ",")
	}

	if modelJson.Groups != nil {
		columns, groups = services.ModelServices.GetModelGroups(c, *modelJson)
	}
	count := services.DbService.Count(c, modelJson.Table, joins)
	result := []map[string]interface{}{}
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
	formJson := services.FormServices.GetFormFile(c, form)
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

// Delete 删除
//
//	@param c
func Delete(c *gin.Context) {
	table := c.Param("table")
	tableJson := services.TableServices.GetTableFile(c, table)
	modelJson := services.ModelServices.GetModelFile(c, tableJson.Model)
	columns, deleted := services.ModelServices.GetModelDeleteds(c, *modelJson)
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	services.DbService.Delete(c, modelJson.Table, idInt, columns, deleted)
	c.JSON(200, gin.H{
		"message": "Delete",
	})
}

func Test(c *gin.Context) {
	result := []map[string]interface{}{}
	services.DbService.Get(c, "logins", &result, "*", "", []string{}, "")

	loginStr := "2023-01-01 00:00:00"
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, loginStr)
	timestamp := t.Unix()
	for k, v := range result {
		loginTime := (k+1)*3600 + int(timestamp)
		t = time.Unix(int64(loginTime), 0)
		loginDate := t.Format(layout)
		services.DbService.Update(c, "logins", int(v["id"].(uint32)), map[string]interface{}{"created_at": loginDate, "updated_at": loginDate})
	}

	c.JSON(200, gin.H{
		"message": "Update",
		"data":    result,
	})
}
