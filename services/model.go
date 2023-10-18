package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/models"
	"strings"
	"time"

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

// GetModelFileQueryAfter 获取 model.json 文件查询后信息处理
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sModelServices) GetModelFileQueryAfter(c *gin.Context, result map[string]interface{}, model models.ModelJson) {
	if model.Withs != nil && len(model.Withs) > 0 {
		ResultServices.HandleModelWiths(c, result, model)
	}

	if model.WithsGroups != nil && len(model.WithsGroups) > 0 {
		ResultServices.HandleModelWithsGroups(c, result, model)
	}

	if model.Columns != nil && len(model.Columns) > 0 {
		ResultServices.HandleModelFieldFormat(c, result, model)
	}
}

// GetModelFileQueryAfterList 获取 model.json 文件查询后信息处理
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sModelServices) GetModelFileQueryAfterList(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	if model.Withs != nil && len(model.Withs) > 0 {
		ResultServices.HandleModelWithsList(c, result, model)

		for _, value := range model.Withs {
			for _, val := range value.Columns {
				if val.Format != "" {
					val.Format = strings.ReplaceAll(val.Format, "Y", "2006")
					val.Format = strings.ReplaceAll(val.Format, "m", "01")
					val.Format = strings.ReplaceAll(val.Format, "d", "02")
					val.Format = strings.ReplaceAll(val.Format, "H", "15")
					val.Format = strings.ReplaceAll(val.Format, "i", "04")
					val.Format = strings.ReplaceAll(val.Format, "s", "05")

					for _, v := range result {
						for _, v := range v["withs_"+value.Table].([]map[string]interface{}) {
							date, _ := v[val.Field].(time.Time)
							v[val.Field] = date.Format(val.Format)
						}
					}
				}
			}
		}
	}

	if model.WithsGroups != nil && len(model.WithsGroups) > 0 {
		ResultServices.HandleModelWithsGroupsList(c, result, model)
	}

	if model.Columns != nil && len(model.Columns) > 0 {
		ResultServices.HandleModelFieldFormatList(c, result, model)
	}
}

// GetModelColumns 获取 model.json 文件 columns 信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
func (s *sModelServices) GetModelColumns(c *gin.Context, model models.ModelJson) []string {
	columns := []string{}
	if model.Columns == nil || len(model.Columns) == 0 {
		columns = append(columns, model.Table+".*")
	} else {
		for _, v := range model.Columns {
			columns = append(columns, model.Table+"."+v.Field)
		}
	}

	return columns
}

// GetModelOrders 获取 model.json 文件 orders 信息
//
//	@receiver s
//	@param c
//	@param model
//	@return string
func (s *sModelServices) GetModelOrders(c *gin.Context, model models.ModelJson) string {
	orders := []string{}
	for _, v := range model.Orders {
		orders = append(orders, model.Table+"."+v.Field+" "+strings.ToUpper(v.Sort))
	}
	return strings.Join(orders, ",")
}

// GetModelPrimary 获取 model.json 文件主键信息
//
//	@receiver s
//	@param c
//	@param model
//	@return string
func (s *sModelServices) GetModelPrimary(c *gin.Context, model models.ModelJson) string {
	primary := "id"
	for _, v := range model.Columns {
		if v.Primary {
			primary = v.Field
			break
		}
	}

	return primary
}

// GetModelJoinsGroups 获取 model.json 文件 joinGroups 信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
//	@return []string
//	@return []string
//	@return string
func (s *sModelServices) GetModelJoinsGroups(c *gin.Context, model models.ModelJson) ([]string, []string, []string, string) {
	joins := []string{}
	columns := []string{}
	groups := []string{}
	orders := []string{}

	for _, value := range model.JoinsGroups {
		joinsGroupsTable := strings.ToUpper(value.Join) + " JOIN " + value.Table + " ON " + value.Table + "." + value.Foreign + " = " + model.Table + "." + value.Key
		joinsGroupsWhere := s.HandleModelJoinsGroupsWheres(c, value.Wheres, value.Table)
		if len(joinsGroupsWhere) > 0 {
			joinsGroupsTable += " AND ( " + strings.Join(joinsGroupsWhere, " AND ") + " )"
		}
		joins = append(joins, joinsGroupsTable)

		for _, v := range value.Columns {
			switch strings.ToUpper(value.Type) {
			case "SUM":
				columns = append(columns, "SUM( "+value.Table+"."+v.Field+" ) AS "+value.Table+"_"+v.Field+"_sum")

			case "COUNT":
				columns = append(columns, "COUNT( "+value.Table+"."+v.Field+" ) AS "+value.Table+"_"+v.Field+"_count")

			case "MAX":
				columns = append(columns, "MAX( "+value.Table+"."+v.Field+" ) AS "+value.Table+"_"+v.Field+"_max")

			case "MIN":
				columns = append(columns, "MIN( "+value.Table+"."+v.Field+" ) AS "+value.Table+"_"+v.Field+"_min")
			}
		}

		groups = append(groups, value.Table+"."+value.Foreign)

		for _, v := range value.Orders {
			switch strings.ToUpper(value.Type) {
			case "SUM":
				orders = append(orders, value.Table+"_"+v.Field+"_sum "+strings.ToUpper(v.Sort))

			case "COUNT":
				orders = append(orders, value.Table+"_"+v.Field+"_count "+strings.ToUpper(v.Sort))

			case "MAX":
				orders = append(orders, value.Table+"_"+v.Field+"_max "+strings.ToUpper(v.Sort))

			case "MIN":
				orders = append(orders, value.Table+"_"+v.Field+"_min "+strings.ToUpper(v.Sort))
			}
		}
	}

	uniqueMap := make(map[string]bool)
	uniqueJoins := []string{}
	for _, v := range joins {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			uniqueJoins = append(uniqueJoins, v)
		}
	}

	return uniqueJoins, columns, groups, strings.Join(orders, ",")
}

// HandleModelJoinsGroupsWheres HandleModelJoinsGroupsWheres 处理 model.json 文件 joinGroups 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param wheres
//	@param table
//	@return []string
func (s *sModelServices) HandleModelJoinsGroupsWheres(c *gin.Context, wheres []models.Wheres, table string) []string {
	whereSlice := []string{}
	for _, v := range wheres {
		switch strings.ToUpper(v.Match) {
		case "=", "!=", "<>", ">", "<", ">=", "<=":
			whereSlice = append(whereSlice, table+"."+v.Field+" "+v.Match+" '"+v.Value+"'")
		case "IN":
			whereSlice = append(whereSlice, table+"."+v.Field+" IN ("+v.Value+")")
		case "LIKE":
			whereSlice = append(whereSlice, table+"."+v.Field+" LIKE '%"+v.Value+"%'")
		case "LIKE.LEFT":
			whereSlice = append(whereSlice, table+"."+v.Field+" LIKE '%"+v.Value)
		case "LIKE.RIGHT":
			whereSlice = append(whereSlice, table+"."+v.Field+" LIKE '"+v.Value+"%'")
		case "BETWEEN":
			values := strings.Split(v.Value, "~")
			whereSlice = append(whereSlice, table+"."+v.Field+" BETWEEN '"+values[0]+"' AND '"+values[1]+"'")
		case "IS":
			switch strings.ToUpper(v.Value) {
			case "NULL":
				whereSlice = append(whereSlice, table+"."+v.Field+" IS NULL")
			case "NOTNULL":
				whereSlice = append(whereSlice, table+"."+v.Field+" IS NOT NULL")
			}
		}
	}

	return whereSlice
}

// GetModelJoins 获取 model.json 文件 joins 信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
func (s *sModelServices) GetModelJoins(c *gin.Context, model models.ModelJson) []string {
	joins := []string{}
	for _, value := range model.Joins {
		joinTable := strings.ToUpper(value.Join) + " JOIN " + value.Table + " ON " + value.Table + "." + value.Foreign + " = " + model.Table + "." + value.Key
		if value.Wheres != nil && len(value.Wheres) > 0 {
			joinWhere := s.HandleModelJoinsWheres(c, value.Wheres, value.Table)
			if joinWhere != nil {
				joinTable += " AND ( " + strings.Join(joinWhere, " AND ") + " )"
			}
		}
		joins = append(joins, joinTable)
	}

	return joins
}

// GetModelGroups 获取 model.json 文件 groups 信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
//	@return []string
func (s *sModelServices) GetModelGroups(c *gin.Context, model models.ModelJson) ([]string, []string) {
	columns := []string{}
	groups := []string{}

	for _, value := range model.Groups {
		groupType := strings.ToUpper(value.Type)
		if groupType == "COUNT" {
			for _, v := range value.Columns {
				columns = append(columns, model.Table+"."+value.Group.Field+", COUNT( "+model.Table+"."+v.Field+" ) AS "+v.Field+"_count")
				groups = append(groups, value.Group.Field)
			}
		} else if groupType == "SUM" {
			for _, v := range value.Columns {
				columns = append(columns, model.Table+"."+value.Group.Field+", SUM( "+model.Table+"."+v.Field+" ) AS "+v.Field+"_sum")
				groups = append(groups, value.Group.Field)
			}
		}
	}

	return columns, groups
}

// GetModelJoinsColumns 获取 model.json 文件 joins 下 columns 信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
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
				columns = append(columns, value.Table+"."+v.Field+" AS joins_"+value.Table+"_"+v.Field)
			}
		} else {
			for _, v := range value.Columns {
				columns = append(columns, value.Table+"."+v.Field+" AS joins_"+value.Table+"_"+v.Field)
			}
		}
	}

	return columns
}

// GetModelWithsColumns 获取 model.json 文件 withs 下 columns 信息
//
//	@receiver s
//	@param c
//	@param withs
//	@return []string
func (s *sModelServices) GetModelWithsColumns(c *gin.Context, withs models.Withs) []string {
	columns := []string{}
	for _, v := range withs.Columns {
		columns = append(columns, v.Field)
	}

	return columns
}

// GetModelWithsOrders 获取 model.json 文件 withs 下 orders 信息
//
//	@receiver s
//	@param c
//	@param withs
//	@return string
func (s *sModelServices) GetModelWithsOrders(c *gin.Context, withs models.Withs) string {
	orders := []string{}
	for _, v := range withs.Orders {
		orders = append(orders, v.Field+" "+strings.ToUpper(v.Sort))
	}

	return strings.Join(orders, ",")
}

// GetModelDeleteds 获取 model.json 文件 deleteds 信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
//	@return map[string]interface{}
func (s *sModelServices) GetModelDeleteds(c *gin.Context, model models.ModelJson) ([]string, map[string]interface{}) {
	columns := []string{}
	deleted := map[string]interface{}{}
	t := time.Now()
	for _, v := range model.Deleteds {
		columns = append(columns, v.Field)
		switch v.Value {
		case "date":
			deleted[v.Field] = t.Format("2006-01-02")
		case "datetime":
			deleted[v.Field] = t.Format("2006-01-02 15:04:05")
		case "timestamp":
			deleted[v.Field] = t.Unix()
		default:
			deleted[v.Field] = v.Value
		}
	}

	return columns, deleted
}

// HandleModelJoinsWheres 处理 model.json 文件 joins 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param wheres
//	@param table
//	@return []string
func (s *sModelServices) HandleModelJoinsWheres(c *gin.Context, wheres []models.Wheres, table string) []string {
	whereSlice := []string{}
	for _, v := range wheres {
		switch strings.ToUpper(v.Match) {
		case "=", "!=", "<>", ">", "<", ">=", "<=":
			whereSlice = append(whereSlice, table+"."+v.Field+" "+v.Match+" '"+v.Value+"'")
		case "IN":
			whereSlice = append(whereSlice, table+"."+v.Field+" IN ("+v.Value+")")
		case "LIKE":
			whereSlice = append(whereSlice, table+"."+v.Field+" LIKE '%"+v.Value+"%'")
		case "LIKE.LEFT":
			whereSlice = append(whereSlice, table+"."+v.Field+" LIKE '%"+v.Value)
		case "LIKE.RIGHT":
			whereSlice = append(whereSlice, table+"."+v.Field+" LIKE '"+v.Value+"%'")
		case "BETWEEN":
			values := strings.Split(v.Value, "~")
			whereSlice = append(whereSlice, table+"."+v.Field+" BETWEEN '"+values[0]+"' AND '"+values[1]+"'")
		case "IS":
			switch strings.ToUpper(v.Value) {
			case "NULL":
				whereSlice = append(whereSlice, table+"."+v.Field+" IS NULL")
			case "NOTNULL":
				whereSlice = append(whereSlice, table+"."+v.Field+" IS NOT NULL")
			}
		}
	}
	return whereSlice
}

// HandleModelWithsWheres 处理 model.json 文件 withs 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param wheres
//	@return string
func (s *sModelServices) HandleModelWithsWheres(c *gin.Context, wheres []models.Wheres) string {
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
