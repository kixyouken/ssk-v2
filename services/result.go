package services

import (
	"ssk-v2/jsons/forms"
	"ssk-v2/jsons/models"
	"ssk-v2/jsons/tables"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type sResultServices struct{}

var ResultServices = sResultServices{}

// HandleFormWiths 处理 form 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param form
func (s *sResultServices) HandleFormWiths(c *gin.Context, result map[string]interface{}, form forms.FormJson) {
	for _, value := range form.Withs {
		modelJson := ModelServices.GetModelFile(c, value.Model)
		if result[value.Key] != nil {
			columns := FormServices.GetFormWithsColumns(c, value)
			if value.Has == "hasOne" {
				withResult := map[string]interface{}{}
				wheres := ""
				if value.Wheres != nil && len(value.Wheres) > 0 {
					wheres = FormServices.HandleFormWithsWheres(c, value.Wheres)
				}
				DbService.HasOne(c, modelJson.Table, &withResult, columns, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
				result["withs_"+modelJson.Table] = withResult
			} else if value.Has == "hasMany" {
				withResult := []map[string]interface{}{}
				orders := FormServices.GetModelWithsOrders(c, value)
				wheres := ""
				if value.Wheres != nil && len(value.Wheres) > 0 {
					wheres = FormServices.HandleFormWithsWheres(c, value.Wheres)
				}
				DbService.HasMany(c, modelJson.Table, &withResult, columns, orders, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
				result["withs_"+modelJson.Table] = withResult
			}
		} else {
			result["withs_"+modelJson.Table] = nil
		}
	}
}

// HandleTableWiths 处理 table 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param table
func (s *sResultServices) HandleTableWiths(c *gin.Context, result map[string]interface{}, table tables.TableJson) {
	for _, value := range table.Withs {
		modelJson := ModelServices.GetModelFile(c, value.Model)
		if result[value.Key] != nil {
			columns := TableServices.GetTableWithsColumns(c, value)
			if value.Has == "hasOne" {
				withResult := map[string]interface{}{}
				wheres := ""
				if value.Wheres != nil && len(value.Wheres) > 0 {
					wheres = TableServices.HandleTableWithsWheres(c, value.Wheres)
				}
				DbService.HasOne(c, modelJson.Table, &withResult, columns, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
				result["withs_"+modelJson.Table] = withResult
			} else if value.Has == "hasMany" {
				withResult := []map[string]interface{}{}
				orders := TableServices.GetTableWithsOrders(c, value)
				wheres := ""
				if value.Wheres != nil && len(value.Wheres) > 0 {
					wheres = TableServices.HandleTableWithsWheres(c, value.Wheres)
				}
				DbService.HasMany(c, modelJson.Table, &withResult, columns, orders, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
				result["withs_"+modelJson.Table] = withResult
			}
		} else {
			result["withs_"+modelJson.Table] = nil
		}
	}
}

// HandleTableWithsList 处理 table 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param table
func (s *sResultServices) HandleTableWithsList(c *gin.Context, result []map[string]interface{}, table tables.TableJson) {
	for _, value := range result {
		s.HandleTableWiths(c, value, table)
	}
}

// HandleModelWiths 处理 model 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelWiths(c *gin.Context, result map[string]interface{}, model models.ModelJson) {
	for _, value := range model.Withs {
		if result[value.Key] != nil {
			columns := ModelServices.GetModelWithsColumns(c, value)
			if value.Has == "hasOne" {
				withResult := map[string]interface{}{}
				wheres := ""
				if value.Wheres != nil && len(value.Wheres) > 0 {
					wheres = ModelServices.HandleModelWithsWheres(c, value.Wheres)
				}
				DbService.HasOne(c, value.Table, &withResult, columns, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
				result["withs_"+value.Table] = withResult
			} else if value.Has == "hasMany" {
				orders := ModelServices.GetModelWithsOrders(c, value)
				withResult := []map[string]interface{}{}
				wheres := ""
				if value.Wheres != nil && len(value.Wheres) > 0 {
					wheres = ModelServices.HandleModelWithsWheres(c, value.Wheres)
				}
				DbService.HasMany(c, value.Table, &withResult, columns, orders, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
				result["withs_"+value.Table] = withResult
			}
		} else {
			result["withs_"+value.Table] = nil
		}
	}
}

// HandleModelWiths 处理 model 下 withs 关联信息
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelWithsList(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	for _, value := range result {
		s.HandleModelWiths(c, value, model)
	}
}

// HandleModelFieldFormat 处理 field 格式化
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelFieldFormat(c *gin.Context, result map[string]interface{}, model models.ModelJson) {
	for _, v := range model.Columns {
		if v.Format != "" {
			v.Format = strings.ReplaceAll(v.Format, "Y", "2006")
			v.Format = strings.ReplaceAll(v.Format, "m", "01")
			v.Format = strings.ReplaceAll(v.Format, "d", "02")
			v.Format = strings.ReplaceAll(v.Format, "H", "15")
			v.Format = strings.ReplaceAll(v.Format, "i", "04")
			v.Format = strings.ReplaceAll(v.Format, "s", "05")
			switch v.Type {
			case "int":
				timestamp := result[v.Field].(uint32)
				if timestamp > 0 {
					t := time.Unix(int64(timestamp), 0)
					result[v.Field] = t.Format(v.Format)
				}
			case "date", "datetime", "timestamp":
				if result[v.Field] != nil {
					date, _ := result[v.Field].(time.Time)
					result[v.Field] = date.Format(v.Format)
				}
			}
		}
	}

	for _, value := range model.Joins {
		for _, v := range value.Columns {
			if v.Format != "" {
				v.Format = strings.ReplaceAll(v.Format, "Y", "2006")
				v.Format = strings.ReplaceAll(v.Format, "m", "01")
				v.Format = strings.ReplaceAll(v.Format, "d", "02")
				v.Format = strings.ReplaceAll(v.Format, "H", "15")
				v.Format = strings.ReplaceAll(v.Format, "i", "04")
				v.Format = strings.ReplaceAll(v.Format, "s", "05")

				date, _ := result["joins_"+value.Table+"_"+v.Field].(time.Time)
				result["joins_"+value.Table+"_"+v.Field] = date.Format(v.Format)
			}
		}
	}
}

// HandleModelFieldFormat 处理 field 格式化
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelFieldFormatList(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	for _, value := range result {
		s.HandleModelFieldFormat(c, value, model)
	}
}

// HandleModelWithsGroups 处理 model 下 withs_groups 统计信息
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelWithsGroups(c *gin.Context, result map[string]interface{}, model models.ModelJson) {
	for _, value := range model.WithsGroups {
		withsWheres := []string{}
		if value.Wheres != nil && len(value.Wheres) > 0 {
			for _, v := range value.Wheres {
				switch strings.ToUpper(v.Match) {
				case "=", "!=", "<>", ">", "<", ">=", "<=":
					withsWheres = append(withsWheres, v.Field+" "+v.Match+" '"+v.Value+"'")
				case "IN":
					withsWheres = append(withsWheres, v.Field+" IN ("+v.Value+")")
				case "LIKE":
					withsWheres = append(withsWheres, v.Field+" LIKE '%"+v.Value+"%'")
				case "LIKE.LEFT":
					withsWheres = append(withsWheres, v.Field+" LIKE '%"+v.Value)
				case "LIKE.RIGHT":
					withsWheres = append(withsWheres, v.Field+" LIKE '"+v.Value+"%'")
				case "BETWEEN":
					values := strings.Split(v.Value, "~")
					withsWheres = append(withsWheres, v.Field+" BETWEEN '"+values[0]+"' AND '"+values[1]+"'")
				case "IS":
					switch strings.ToUpper(v.Value) {
					case "NULL":
						withsWheres = append(withsWheres, v.Field+" IS NULL")
					case "NOTNULL":
						withsWheres = append(withsWheres, v.Field+" IS NOT NULL")
					}
				}
			}
		}
		wheres := strings.Join(withsWheres, " AND ")

		switch strings.ToUpper(value.Type) {
		case "SUM":
			withsSum := []string{}
			if value.Columns != nil && len(value.Columns) > 0 {
				for _, v := range value.Columns {
					withsSum = append(withsSum, "SUM("+v.Field+") AS "+v.Field+"_sum")
				}
			}
			sumResult := map[string]interface{}{}
			DbService.WithsSum(c, value.Table, &sumResult, withsSum, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
			result["withs_"+value.Table+"_sum"] = sumResult

		case "COUNT":
			result["withs_"+value.Table+"_count"] = DbService.WithsCount(c, value.Table, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)

		case "MAX":
			withsMax := []string{}
			if value.Columns != nil && len(value.Columns) > 0 {
				for _, v := range value.Columns {
					withsMax = append(withsMax, "MAX("+v.Field+") AS "+v.Field+"_max")
				}
			}
			maxResult := map[string]interface{}{}
			DbService.WithsMax(c, value.Table, &maxResult, withsMax, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
			result["withs_"+value.Table+"_max"] = maxResult
		case "MIN":
			withsMin := []string{}
			if value.Columns != nil && len(value.Columns) > 0 {
				for _, v := range value.Columns {
					withsMin = append(withsMin, "MIN("+v.Field+") AS "+v.Field+"_min")
				}
			}
			minResult := map[string]interface{}{}
			DbService.WithsMin(c, value.Table, &minResult, withsMin, map[string]interface{}{value.Foreign: result[value.Key]}, wheres)
			result["withs_"+value.Table+"_min"] = minResult
		}
	}
}

// HandleModelWithsGroupsList 处理 model 下 withs_groups 统计信息
//
//	@receiver s
//	@param c
//	@param result
//	@param model
func (s *sResultServices) HandleModelWithsGroupsList(c *gin.Context, result []map[string]interface{}, model models.ModelJson) {
	for _, v := range result {
		s.HandleModelWithsGroups(c, v, model)
	}
}
