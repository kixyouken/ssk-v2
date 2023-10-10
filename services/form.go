package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/forms"
	"strings"
	"time"

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
func (s *sFormServices) GetFormFile(c *gin.Context, form string) *forms.FormJson {
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

// GetFormFileQueryAfter 获取 form.json 文件查询后信息处理
//
//	@receiver s
//	@param c
//	@param result
//	@param form
func (s *sFormServices) GetFormFileQueryAfter(c *gin.Context, result map[string]interface{}, form forms.FormJson) {
	if form.Withs != nil && len(form.Withs) > 0 {
		ResultServices.HandleFormWiths(c, result, form)

		for _, value := range form.Withs {
			modelJson := ModelServices.GetModelFile(c, value.Model)
			for _, val := range value.Columns {
				if val.Format != "" {
					val.Format = strings.ReplaceAll(val.Format, "Y", "2006")
					val.Format = strings.ReplaceAll(val.Format, "m", "01")
					val.Format = strings.ReplaceAll(val.Format, "d", "02")
					val.Format = strings.ReplaceAll(val.Format, "H", "15")
					val.Format = strings.ReplaceAll(val.Format, "i", "04")
					val.Format = strings.ReplaceAll(val.Format, "s", "05")

					for _, v := range result["withs_"+modelJson.Table].([]map[string]interface{}) {
						date, _ := v[val.Field].(time.Time)
						v[val.Field] = date.Format(val.Format)
					}
				}
			}
		}
	}
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

// HandleFormWithsWheres 处理 form.json 文件 withs 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param wheres
//	@return string
func (s *sFormServices) HandleFormWithsWheres(c *gin.Context, wheres []forms.Wheres) string {
	whereSlice := []string{}
	for _, v := range wheres {
		switch strings.ToUpper(v.Match) {
		case "=", "!=", "<>", ">", "<", ">=", "<=":
			whereSlice = append(whereSlice, v.Field+" "+v.Match+" '"+v.Value+"'")
		case "IN":
			whereSlice = append(whereSlice, v.Field+" IN ("+v.Value+")")
		case "LIKE":
			whereSlice = append(whereSlice, v.Field+" LIKE '%"+v.Value+"%'")
		case "LIKE.LEFT":
			whereSlice = append(whereSlice, v.Field+" LIKE '%"+v.Value)
		case "LIKE.RIGHT":
			whereSlice = append(whereSlice, v.Field+" LIKE '"+v.Value+"%'")
		case "BETWEEN":
			values := strings.Split(v.Value, "~")
			whereSlice = append(whereSlice, v.Field+" BETWEEN '"+values[0]+"' AND '"+values[1]+"'")
		case "IS":
			switch strings.ToUpper(v.Value) {
			case "NULL":
				whereSlice = append(whereSlice, v.Field+" IS NULL")
			case "NOTNULL":
				whereSlice = append(whereSlice, v.Field+" IS NOT NULL")
			}
		}
	}

	return strings.Join(whereSlice, " AND ")
}
