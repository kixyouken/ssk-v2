package controllers

import (
	"ssk-v2/services"

	"github.com/gin-gonic/gin"
)

// GetTable 获取 table 配置
//
//	@param c
func GetTable(c *gin.Context) {
	table := c.Param("table")
	tableJson := services.TableServices.GetTableFile(c, table)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    tableJson,
	})
}

// GetForm 获取 form 配置
//
//	@param c
func GetForm(c *gin.Context) {
	form := c.Param("form")
	formJson := services.FormServices.GetFormFile(c, form)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    formJson,
	})
}

// GetModel 获取 model 配置
//
//	@param c
func GetModel(c *gin.Context) {
	model := c.Param("model")
	modelJson := services.ModelServices.GetModelFile(c, model)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    modelJson,
	})
}
