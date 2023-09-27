package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type sModelServices struct{}

var ModelServices = sModelServices{}

// GetModelFile 获取 model.json 文件
//
//	@receiver s
//	@param c
//	@param model
//	@return *models.ModelJson
func (s *sModelServices) GetModelFile(c *gin.Context, model string) *models.ModelJson {
	modelFile := "./json/model/" + model + ".json"
	body, err := os.ReadFile(modelFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read JSON file"})
		return nil
	}

	modelJson := models.ModelJson{}
	json.Unmarshal(body, &modelJson)

	return &modelJson
}

func (s *sModelServices) GetModelColumns(c *gin.Context, model models.ModelJson) []string {
	columns := []string{}
	if model.Columns == nil || len(model.Columns) == 0 {
		columns = append(columns, model.Table+".*")
	} else {
		for _, v := range model.Columns {
			columns = append(columns, model.Table+"."+v.Field)
		}
	}

	joinColumns := s.GetModelJoinsColumns(c, model)
	columns = append(columns, joinColumns...)
	return columns
}

func (s *sModelServices) GetModelJoins(c *gin.Context, model models.ModelJson) []string {
	joins := []string{}
	for _, value := range model.Joins {
		joinTable := strings.ToUpper(value.Join) + " JOIN " + value.Table + " ON " + value.Table + "." + value.Foreign + " = " + model.Table + "." + value.Key
		joins = append(joins, joinTable)
	}

	return joins
}

func (s *sModelServices) GetModelJoinsColumns(c *gin.Context, model models.ModelJson) []string {
	type JoinColumns struct {
		Field string `json:"field"`
	}
	columns := []string{}
	for _, value := range model.Joins {
		if value.Columns == nil || len(value.Columns) == 0 {
			joinColumns := []JoinColumns{}
			db.Raw("SHOW COLUMNS FROM `" + value.Table + "`").Scan(&joinColumns)
			for _, v := range joinColumns {
				columns = append(columns, value.Table+"."+v.Field+" AS join_"+value.Table+"_"+v.Field)
			}
		} else {
			for _, v := range value.Columns {
				columns = append(columns, value.Table+"."+v.Field+" AS join_"+value.Table+"_"+v.Field)
			}
		}
	}

	return columns
}
