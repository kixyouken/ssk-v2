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

// GetTableFileQueryBefore 获取 table.json 文件查询前信息处理
//
//	@receiver s
//	@param c
//	@param table
//	@return []string
//	@return []string
//	@return []string
//	@return string
func (s *sTableServices) GetTableFileQueryBefore(c *gin.Context, table tables.TableJson) ([]string, []string, []string, string) {
	columns := []string{}
	joins := []string{}
	groups := []string{}
	if table.Joins != nil && len(table.Joins) > 0 {
		tableJoinsColumns := s.GetTableJoinsColumns(c, table)
		columns = append(columns, tableJoinsColumns...)

		tableJoins := s.GetTableJoins(c, table)
		joins = append(joins, tableJoins...)
	}

	orders := ""
	if table.Orders != nil && len(table.Orders) > 0 {
		orders = s.GetTableOrders(c, table)
	}

	if table.JoinsCount != nil && len(table.JoinsCount) > 0 {
		tableJoinsCountColumns := s.GetTableJoinsCountColumns(c, table)
		columns = append(columns, tableJoinsCountColumns...)

		tableJoinsCount := s.GetTableJoinsCount(c, table)
		joins = append(joins, tableJoinsCount...)

		tableJoinsGroups := s.GetTableJoinsCountGroups(c, table)
		groups = append(groups, tableJoinsGroups...)

		tableJoinsCountOrders := s.GetTableJoinsCountOrders(c, table)
		if tableJoinsCountOrders != "" {
			orders = tableJoinsCountOrders + "," + orders
		}
	}

	return columns, joins, groups, orders
}

// GetTableFileQueryAfterList 获取 table.json 文件查询后信息处理
//
//	@receiver s
//	@param c
//	@param result
//	@param table
func (s *sTableServices) GetTableFileQueryAfterList(c *gin.Context, result []map[string]interface{}, table tables.TableJson) {
	if table.Withs != nil && len(table.Withs) > 0 {
		ResultServices.HandleTableWithsList(c, result, table)
	}

	if table.WithsCount != nil && len(table.WithsCount) > 0 {
		ResultServices.HandleTableWithsCountList(c, result, table)
	}

	if table.WithsSum != nil && len(table.WithsSum) > 0 {
		ResultServices.HandleTableWithsSumList(c, result, table)
	}
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

// GetTableJoinsCountGroups 获取 table.json 文件 joinsCount 时 group 信息
//
//	@receiver s
//	@param c
//	@param table
//	@return []string
func (s *sTableServices) GetTableJoinsCountGroups(c *gin.Context, table tables.TableJson) []string {
	groups := []string{}
	for _, value := range table.JoinsCount {
		modelJson := ModelServices.GetModelFile(c, value.Model)
		groups = append(groups, modelJson.Table+"."+value.Foreign)
	}

	return groups
}

// GetTableJoinsCountOrders 获取 table.json 文件 joinsCount 下 orders 信息
//
//	@receiver s
//	@param c
//	@param table
//	@return string
func (s *sTableServices) GetTableJoinsCountOrders(c *gin.Context, table tables.TableJson) string {
	orders := []string{}
	for _, value := range table.JoinsCount {
		for _, v := range value.Orders {
			modelJson := ModelServices.GetModelFile(c, value.Model)
			orders = append(orders, modelJson.Table+"_"+v.Field+"_count "+strings.ToUpper(v.Sort))
		}
	}

	return strings.Join(orders, ",")
}

// GetTableJoinsCountColumns 获取 table.json 文件 joinsCount 下 columns 信息
//
//	@receiver s
//	@param c
//	@param table
//	@return []string
func (s *sTableServices) GetTableJoinsCountColumns(c *gin.Context, table tables.TableJson) []string {
	columns := []string{}
	for _, value := range table.JoinsCount {
		modelJson := ModelServices.GetModelFile(c, value.Model)
		for _, v := range value.Columns {
			columns = append(columns, "COUNT( "+modelJson.Table+"."+v.Field+" ) AS "+modelJson.Table+"_"+v.Field+"_count")
		}
	}

	return columns
}

// GetTableJoinsCount 获取 table.json 文件 joinsCount 信息
//
//	@receiver s
//	@param c
//	@param table
//	@return []string
func (s *sTableServices) GetTableJoinsCount(c *gin.Context, table tables.TableJson) []string {
	modelJson := ModelServices.GetModelFile(c, table.Model)
	joins := []string{}
	for _, value := range table.JoinsCount {
		modelJoinJson := ModelServices.GetModelFile(c, value.Model)
		joinTable := strings.ToUpper(value.Join) + " JOIN " + modelJoinJson.Table + " ON " + modelJoinJson.Table + "." + value.Foreign + " = " + modelJson.Table + "." + value.Key
		if value.Wheres != nil && len(value.Wheres) > 0 {
			joinWhere := s.HandleTableJoinsCountWheres(c, value.Wheres, modelJoinJson.Table)
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

// HandleTableJoinsCountWheres 处理 table.json 文件 joinsCount 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param wheres
//	@param table
//	@return string
func (s *sTableServices) HandleTableJoinsCountWheres(c *gin.Context, wheres []tables.Wheres, table string) []string {
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
