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

	joinColumns := s.GetModelJoinsColumns(c, model)
	columns = append(columns, joinColumns...)
	return columns
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
			joinWhere := []string{}
			for _, v := range value.Wheres {
				joinWhere = append(joinWhere, value.Table+"."+s.HandleModelJoinsWheres(c, v))
			}
			if joinWhere != nil {
				joinTable += " AND ( " + strings.Join(joinWhere, " AND ") + " )"
			}
		}
		joins = append(joins, joinTable)
	}

	return joins
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

// HandleModelJoinsWheres 处理 model.json 文件 joins 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param where
//	@return string
func (s *sModelServices) HandleModelJoinsWheres(c *gin.Context, where models.Wheres) string {
	wheres := []string{}
	switch strings.ToUpper(where.Match) {
	case "=", "!=", "<>", ">", "<", ">=", "<=":
		wheres = append(wheres, where.Field+" "+where.Match+" '"+where.Value+"'")
	case "IN":
		wheres = append(wheres, where.Field+" IN ("+where.Value+")")
	case "LIKE":
		wheres = append(wheres, where.Field+" LIKE '%"+where.Value+"%'")
	case "LIKE.LEFT":
		wheres = append(wheres, where.Field+" LIKE '%"+where.Value)
	case "LIKE.RIGHT":
		wheres = append(wheres, where.Field+" LIKE '"+where.Value+"%'")
	case "BETWEEN":
		values := strings.Split(where.Value, ",")
		wheres = append(wheres, where.Field+" BETWEEN '"+values[0]+"' AND '"+values[1]+"'")
	case "IS":
		switch strings.ToUpper(where.Value) {
		case "NULL":
			wheres = append(wheres, where.Field+" IS NULL")
		case "NOTNULL":
			wheres = append(wheres, where.Field+" IS NOT NULL")
		}
	}

	return strings.Join(wheres, " AND ")
}

// HandleModelWithsWheres 处理 model.json 文件 withs 下 wheres 信息
//
//	@receiver s
//	@param c
//	@param where
//	@return string
func (s *sModelServices) HandleModelWithsWheres(c *gin.Context, where models.Wheres) string {
	wheres := []string{}
	switch strings.ToUpper(where.Match) {
	case "=", "!=", "<>", ">", "<", ">=", "<=":
		wheres = append(wheres, where.Field+" "+where.Match+" '"+where.Value+"'")
	case "IN":
		wheres = append(wheres, where.Field+" IN ("+where.Value+")")
	case "LIKE":
		wheres = append(wheres, where.Field+" LIKE '%"+where.Value+"%'")
	case "LIKE.LEFT":
		wheres = append(wheres, where.Field+" LIKE '%"+where.Value)
	case "LIKE.RIGHT":
		wheres = append(wheres, where.Field+" LIKE '"+where.Value+"%'")
	case "BETWEEN":
		values := strings.Split(where.Value, ",")
		wheres = append(wheres, where.Field+" BETWEEN '"+values[0]+"' AND '"+values[1]+"'")
	case "IS":
		switch strings.ToUpper(where.Value) {
		case "NULL":
			wheres = append(wheres, where.Field+" IS NULL")
		case "NOTNULL":
			wheres = append(wheres, where.Field+" IS NOT NULL")
		}
	}

	return strings.Join(wheres, " AND ")
}
