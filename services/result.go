package services

import (
	"ssk-v2/jsons/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type sResultServices struct{}

var ResultServices = sResultServices{}

// HandleModelWiths 处理 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelWiths(c *gin.Context, result map[string]interface{}, model models.ModelJson) {
	for _, v := range model.Withs {
		if result[v.Key] != nil {
			columns := ModelServices.GetModelWithsColumns(c, model)
			orders := ModelServices.GetModelWithsOrders(c, model)
			if v.Has == "hasOne" {
				withResult := map[string]interface{}{}
				DbService.HasOne(c, v.Table, &withResult, columns, map[string]interface{}{v.Foreign: result[v.Key]})
				result["with_"+v.Table] = withResult
			} else if v.Has == "hasMany" {
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

// HandleModelWiths 处理 withs 关联信息
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
