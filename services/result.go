package services

import (
	"ssk-v2/jsons/models"

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
		}
	}
}
