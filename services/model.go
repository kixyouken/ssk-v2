package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/models"

	"github.com/gin-gonic/gin"
)

type sModelServices struct{}

var ModelServices = sModelServices{}

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
