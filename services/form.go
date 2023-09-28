package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/forms"
	"strings"

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

// GetFormWithsColumns 获取 form.json 文件 withs 下 columns 信息
//
//	@receiver s
//	@param c
//	@param withs
//	@return []string
func (s *sFormServices) GetFormWithsColumns(c *gin.Context, withs forms.Withs) []string {
	columns := []string{}
	for _, v := range withs.Columns {
		columns = append(columns, v.Field)
	}

	return columns
}

// GetModelWithsOrders 获取 form.json 文件 withs 下 orders 信息
//
//	@receiver s
//	@param c
//	@param withs
//	@return string
func (s *sFormServices) GetModelWithsOrders(c *gin.Context, withs forms.Withs) string {
	orders := []string{}
	for _, v := range withs.Orders {
		orders = append(orders, v.Field+" "+strings.ToUpper(v.Sort))
	}

	return strings.Join(orders, ",")
}