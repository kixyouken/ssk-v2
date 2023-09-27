package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/tables"

	"github.com/gin-gonic/gin"
)

type sTableServices struct{}

var TableServices = sTableServices{}

func (s *sTableServices) GetTableFile(c *gin.Context, table string) *tables.TableJson {
	modelFile := "./json/table/" + table + ".json"
	body, err := os.ReadFile(modelFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read JSON file"})
		return nil
	}

	tableJson := tables.TableJson{}
	json.Unmarshal(body, &tableJson)

	return &tableJson
}
