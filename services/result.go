package services

import (
	"ssk-v2/jsons/models"

	"github.com/gin-gonic/gin"
)

type sResultServices struct{}

var ResultServices = sResultServices{}

func (s *sResultServices) GetWiths(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	for _, value := range result {
		for _, v := range model.Withs {
			columns := ModelServices.GetModelWithsColumns(c, model)
			if v.Has == "hasOne" {
				withResult := map[string]interface{}{}
				DbService.HasOne(c, v.Table, &withResult, columns, map[string]interface{}{v.Foreign: value[v.Key]})
				value["with_"+v.Table] = withResult
			} else {
				withResult := []map[string]interface{}{}
				DbService.HasMany(c, v.Table, &withResult, columns, "", map[string]interface{}{v.Foreign: value[v.Key]})
				value["with_"+v.Table] = withResult
			}
		}
	}
}
