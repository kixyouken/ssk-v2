package services

import (
	"ssk-v2/jsons/forms"
	"ssk-v2/jsons/models"
	"ssk-v2/jsons/tables"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type sResultServices struct{}

var ResultServices = sResultServices{}

// HandleFormWiths 处理 form 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param form
func (s *sResultServices) HandleFormWiths(c *gin.Context, result map[string]interface{}, form forms.FormJson) {
	for _, v := range form.Withs {
		modelJson := ModelServices.GetModelFile(c, v.Model)
		if result[v.Key] != nil {
			columns := FormServices.GetFormWithsColumns(c, v)
			if v.Has == "hasOne" {
				withResult := map[string]interface{}{}
				DbService.HasOne(c, modelJson.Table, &withResult, columns, map[string]interface{}{v.Foreign: result[v.Key]})
				result["with_"+modelJson.Table] = withResult
			} else if v.Has == "hasMany" {
				withResult := []map[string]interface{}{}
				orders := FormServices.GetModelWithsOrders(c, v)
				DbService.HasMany(c, modelJson.Table, &withResult, columns, orders, map[string]interface{}{v.Foreign: result[v.Key]})
				result["with_"+modelJson.Table] = withResult
			}
		} else {
			if v.Has == "hasOne" {
				withResult := map[string]interface{}{}
				result["with_"+modelJson.Table] = withResult
			} else if v.Has == "hasMany" {
				withResult := []map[string]interface{}{}
				result["with_"+modelJson.Table] = withResult
			}
		}
	}
}

// HandleTableWiths 处理 table 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param table
func (s *sResultServices) HandleTableWiths(c *gin.Context, result map[string]interface{}, table tables.TableJson) {
	for _, v := range table.Withs {
		modelJson := ModelServices.GetModelFile(c, v.Model)
		if result[v.Key] != nil {
			columns := TableServices.GetTableWithsColumns(c, v)
			if v.Has == "hasOne" {
				withResult := map[string]interface{}{}
				DbService.HasOne(c, modelJson.Table, &withResult, columns, map[string]interface{}{v.Foreign: result[v.Key]})
				result["with_"+modelJson.Table] = withResult
			} else if v.Has == "hasMany" {
				withResult := []map[string]interface{}{}
				orders := TableServices.GetTableWithsOrders(c, v)
				DbService.HasMany(c, modelJson.Table, &withResult, columns, orders, map[string]interface{}{v.Foreign: result[v.Key]})
				result["with_"+modelJson.Table] = withResult
			}
		} else {
			if v.Has == "hasOne" {
				withResult := map[string]interface{}{}
				result["with_"+modelJson.Table] = withResult
			} else if v.Has == "hasMany" {
				withResult := []map[string]interface{}{}
				result["with_"+modelJson.Table] = withResult
			}
		}
	}
}

// HandleTableWithsList 处理 table 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param table
func (s *sResultServices) HandleTableWithsList(c *gin.Context, result []map[string]interface{}, table tables.TableJson) {
	for _, value := range result {
		s.HandleTableWiths(c, value, table)
	}
}

// HandleModelWiths 处理 model 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelWiths(c *gin.Context, result map[string]interface{}, model models.ModelJson) {
	for _, v := range model.Withs {
		if result[v.Key] != nil {
			columns := ModelServices.GetModelWithsColumns(c, v)
			if v.Has == "hasOne" {
				withResult := map[string]interface{}{}
				DbService.HasOne(c, v.Table, &withResult, columns, map[string]interface{}{v.Foreign: result[v.Key]})
				result["with_"+v.Table] = withResult
			} else if v.Has == "hasMany" {
				orders := ModelServices.GetModelWithsOrders(c, v)
				withResult := []map[string]interface{}{}
				DbService.HasMany(c, v.Table, &withResult, columns, orders, map[string]interface{}{v.Foreign: result[v.Key]})
				result["with_"+v.Table] = withResult
			}
		} else {
			if v.Has == "hasOne" {
				withResult := map[string]interface{}{}
				result["with_"+v.Table] = withResult
			} else if v.Has == "hasMany" {
				withResult := []map[string]interface{}{}
				result["with_"+v.Table] = withResult
			}
		}
	}
}

// HandleModelWiths 处理 model 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelWithsList(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	for _, value := range result {
		s.HandleModelWiths(c, value, model)
	}
}

// HandleModelFieldFormat 处理 field 格式化
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelFieldFormat(c *gin.Context, result map[string]interface{}, model models.ModelJson) {
	for _, v := range model.Columns {
		if v.Format != "" {
			v.Format = strings.ReplaceAll(v.Format, "Y", "2006")
			v.Format = strings.ReplaceAll(v.Format, "m", "01")
			v.Format = strings.ReplaceAll(v.Format, "d", "02")
			v.Format = strings.ReplaceAll(v.Format, "H", "15")
			v.Format = strings.ReplaceAll(v.Format, "i", "04")
			v.Format = strings.ReplaceAll(v.Format, "s", "05")

			date, _ := result[v.Field].(time.Time)
			result[v.Field] = date.Format(v.Format)
		}
	}
}

// HandleModelFieldFormat 处理 field 格式化
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelFieldFormatList(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	for _, value := range result {
		s.HandleModelFieldFormat(c, value, model)
	}
}
