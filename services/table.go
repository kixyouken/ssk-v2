package services

import (
	"encoding/json"
	"os"
	"ssk-v2/jsons/tables"
	"strings"

	"github.com/gin-gonic/gin"
)

type sTableServices struct{}

var TableServices = sTableServices{}

// GetTableFile 获取 table.json 文件
//
//	@receiver s
//	@param c
//	@param table
//	@return *tables.TableJson
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

// GetTableOrders 获取 table.json 文件 orders 信息
//
//	@receiver s
//	@param c
//	@param table
//	@return string
func (s *sTableServices) GetTableOrders(c *gin.Context, table tables.TableJson) string {
	model := ModelServices.GetModelFile(c, table.Model)
	orders := []string{}
	for _, v := range table.Orders {
		orders = append(orders, model.Table+"."+v.Field+" "+strings.ToUpper(v.Sort))
	}

	return strings.Join(orders, ",")
}

// GetTableWithsColumns 获取 table.json 文件 withs 下 orders 信息
//
//	@receiver s
//	@param c
//	@param withs
//	@return []string
func (s *sTableServices) GetTableWithsColumns(c *gin.Context, withs tables.Withs) []string {
	columns := []string{}
	for _, v := range withs.Columns {
		columns = append(columns, v.Field)
	}

	return columns
}

// GetTableWithsOrders 获取 table.json 文件 withs 下 orders 信息
//
//	@receiver s
//	@param c
//	@param withs
//	@return string
func (s *sTableServices) GetTableWithsOrders(c *gin.Context, withs tables.Withs) string {
	orders := []string{}
	for _, v := range withs.Orders {
		orders = append(orders, v.Field+" "+strings.ToUpper(v.Sort))
	}

	return strings.Join(orders, ",")
}

// GetTableJoins 获取 table.json 文件 joins 信息
//
//	@receiver s
//	@param c
//	@param table
//	@return []string
func (s *sTableServices) GetTableJoins(c *gin.Context, table tables.TableJson) []string {
	modelJson := ModelServices.GetModelFile(c, table.Model)
	joins := []string{}
	for _, value := range table.Joins {
		modelJoinJson := ModelServices.GetModelFile(c, value.Model)
		joinTable := strings.ToUpper(value.Join) + " JOIN " + modelJoinJson.Table + " ON " + modelJoinJson.Table + "." + value.Foreign + " = " + modelJson.Table + "." + value.Key
		if value.Wheres != nil && len(value.Wheres) > 0 {
			joinWhere := s.HandleTableJoinsWheres(c, value.Wheres, modelJoinJson.Table)
			if joinWhere != nil {
				joinTable += " AND ( " + strings.Join(joinWhere, " AND ") + " )"
			}
		}
		joins = append(joins, joinTable)
	}

	return joins
}

// GetTableJoinsColumns 获取 table.json 文件 joins 下 columns 信息
//
//	@receiver s
//	@param c
//	@param table
//	@return []string
func (s *sTableServices) GetTableJoinsColumns(c *gin.Context, table tables.TableJson) []string {
	type JoinColumns struct {
		Field string `json:"field"`
	}
	columns := []string{}
	for _, value := range table.Joins {
		modelJson := ModelServices.GetModelFile(c, value.Model)
		if value.Columns == nil || len(value.Columns) == 0 {
			joinColumns := []JoinColumns{}
			db.Raw("SHOW COLUMNS FROM `" + modelJson.Table + "`").Scan(&joinColumns)
			for _, v := range joinColumns {
				columns = append(columns, modelJson.Table+"."+v.Field+" AS joins_"+modelJson.Table+"_"+v.Field)
			}
		} else {
			for _, v := range value.Columns {
				columns = append(columns, modelJson.Table+"."+v.Field+" AS joins_"+modelJson.Table+"_"+v.Field)
			}
		}
	}

	return columns
}

// HandleTableJoinsWheres 处理 table.json 文件 joins 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param wheres
//	@param table
//	@return []string
func (s *sTableServices) HandleTableJoinsWheres(c *gin.Context, wheres []tables.Wheres, table string) []string {
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

// HandleTableWithsWheres 处理 table.json 文件 withs 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param wheres
//	@return string
func (s *sTableServices) HandleTableWithsWheres(c *gin.Context, wheres []tables.Wheres) string {
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
			values := strings.Split(v.Value, ",")
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
