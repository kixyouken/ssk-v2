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
func (s *sResultServices) HandleModelWiths(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	for _, value := range result {
		for _, v := range model.Withs {
			if value[v.Key] != nil {
				columns := ModelServices.GetModelWithsColumns(c, model)
				if v.Has == "hasOne" {
					withResult := map[string]interface{}{}
					DbService.HasOne(c, v.Table, &withResult, columns, map[string]interface{}{v.Foreign: value[v.Key]})
					value["with_"+v.Table] = withResult
				} else if v.Has == "hasMany" {
					withResult := []map[string]interface{}{}
					DbService.HasMany(c, v.Table, &withResult, columns, "", map[string]interface{}{v.Foreign: value[v.Key]})
					value["with_"+v.Table] = withResult
				}
			} else {
				if v.Has == "hasOne" {
					withResult := map[string]interface{}{}
					value["with_"+v.Table] = withResult
				} else if v.Has == "hasMany" {
					withResult := []map[string]interface{}{}
					value["with_"+v.Table] = withResult
				}
			}
		}
	}
}

// HandleModelFieldFormat 处理 field 格式化
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelFieldFormat(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	for _, v := range model.Columns {
		if v.Format != "" {
			for _, value := range result {
				v.Format = strings.ReplaceAll(v.Format, "Y", "2006")
				v.Format = strings.ReplaceAll(v.Format, "m", "01")
				v.Format = strings.ReplaceAll(v.Format, "d", "02")
				v.Format = strings.ReplaceAll(v.Format, "H", "15")
				v.Format = strings.ReplaceAll(v.Format, "i", "04")
				v.Format = strings.ReplaceAll(v.Format, "s", "05")

				date, _ := value[v.Field].(time.Time)
				value[v.Field] = date.Format(v.Format)
			}
		}
	}
}
