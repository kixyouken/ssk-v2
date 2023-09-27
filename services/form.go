package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/forms"

	"github.com/gin-gonic/gin"
)

type sFormServices struct{}

var FormServices = sFormServices{}

// GetForm 获取 form.json 文件
//
//	@receiver s
//	@param c
//	@param form
//	@return *forms.FormJson
func (s *sFormServices) GetForm(c *gin.Context, form string) *forms.FormJson {
	formFile := "./json/form/" + form + ".json"
	body, err := os.ReadFile(formFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read JSON file"})
		return nil
	}

	formJson := forms.FormJson{}
	json.Unmarshal(body, &formJson)

	return &formJson
}
