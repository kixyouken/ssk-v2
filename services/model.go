package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/models"

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
	for _, v := range model.Columns {
		columns = append(columns, v.Field)
	}

	return columns
}
