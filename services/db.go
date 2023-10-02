package services

import (
	"ssk-v2/databases"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sDbServices struct{}

var DbService = sDbServices{}

var db = databases.InitMysql()

// Get 获取所有数据
//
//	@receiver s
//	@param c
//	@param table
//	@param out
//	@param column
//	@param order
//	@param joins
//	@param search
//	@return error
func (s *sDbServices) Get(c *gin.Context, table string, out interface{}, column interface{}, order string, joins []string, search interface{}) error {
	return db.Table(table).Where(search).
		Scopes(s.Order(order), s.Joins(joins...)).
		Select(column).
		Find(out).Error
}

// Page 分页获取数据
//
//	@receiver s
//	@param c
//	@param table
//	@param out
//	@param column
//	@param order
//	@param joins
//	@return error
func (s *sDbServices) Page(c *gin.Context, table string, out interface{}, column interface{}, order string, joins []string) error {
	return db.Table(table).
		Scopes(s.Paginate(c), s.Order(order), s.Joins(joins...), s.Wheres(c), s.Search(c)).
		Select(column).
		Find(out).Error
}

// Count 获取数量
//
//	@receiver s
//	@param c
//	@param table
//	@param joins
//	@return int64
func (s *sDbServices) Count(c *gin.Context, table string, joins []string) int64 {
	var count int64
	err := db.Table(table).
		Scopes(s.Joins(joins...), s.Wheres(c), s.Search(c)).
		Count(&count).Error

	if err != nil {
		return 0
	}

	return count
}

// HasOne with 关联单条
//
//	@receiver s
//	@param c
//	@param table
//	@param out
//	@param column
//	@param wheres
//	@param search
//	@return error
func (s *sDbServices) HasOne(c *gin.Context, table string, out interface{}, column interface{}, wheres interface{}, search interface{}) error {
	return db.Table(table).Where(wheres).Where(search).
		Select(column).
		Limit(1).
		Find(out).Error
}

// HasMany with 关联所有
//
//	@receiver s
//	@param c
//	@param table
//	@param out
//	@param column
//	@param order
//	@param wheres
//	@param search
//	@return error
func (s *sDbServices) HasMany(c *gin.Context, table string, out interface{}, column interface{}, order string, wheres interface{}, search interface{}) error {
	return db.Table(table).Where(wheres).Where(search).
		Scopes(s.Order(order)).
		Select(column).
		Find(out).Error
}

// WithsCount withsCount 统计数量
//
//	@receiver s
//	@param c
//	@param table
//	@param wheres
//	@param search
//	@return int64
func (s *sDbServices) WithsCount(c *gin.Context, table string, wheres interface{}, search interface{}) int64 {
	var count int64
	err := db.Table(table).Where(wheres).Where(search).
		Count(&count).Error

	if err != nil {
		return 0
	}

	return count
}

// WithsSum Sum统计
//
//	@receiver s
//	@param c
//	@param table
//	@param out
//	@param column
//	@param wheres
//	@param search
//	@return error
func (s *sDbServices) WithsSum(c *gin.Context, table string, out interface{}, column interface{}, wheres interface{}, search interface{}) error {
	return db.Table(table).Where(wheres).Where(search).
		Select(column).
		Limit(1).
		Scan(out).Error
}

// Read 获取详情
//
//	@receiver s
//	@param c
//	@param table
//	@param id
//	@param out
//	@param column
//	@return error
func (s *sDbServices) Read(c *gin.Context, table string, id int, out interface{}, column interface{}) error {
	return db.Table(table).Where(table+".id = ?", id).
		Limit(1).
		Select(column).
		Find(out).Error
}

func (s *sDbServices) Save(c *gin.Context) {

}

func (s *sDbServices) Update(c *gin.Context) {

}

func (s *sDbServices) Delete(c *gin.Context) {

}

// Order 排序
//
//	@receiver s
//	@param order
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sDbServices) Order(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// Paginate 分页
//
//	@receiver s
//	@param c
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sDbServices) Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		table := c.Param("table")
		tableJson := TableServices.GetTableFile(c, table)

		page := tableJson.Page
		if page <= 0 {
			page = 1
		}
		urlPage, _ := strconv.Atoi(c.Query("page"))
		if urlPage > 0 {
			page = urlPage
		}
		limit := tableJson.Limit
		if limit <= 0 {
			limit = 10
		}
		urlLimit, _ := strconv.Atoi(c.Query("limit"))
		if urlLimit > 0 {
			limit = urlLimit
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

// Joins join 关联
//
//	@receiver s
//	@param joins
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sDbServices) Joins(joins ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, v := range joins {
			db.Joins(v)
		}
		return db
	}
}

// Wheres 默认搜索条件处理
//
//	@receiver s
//	@param c
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sDbServices) Wheres(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		table := c.Param("table")
		tableJson := TableServices.GetTableFile(c, table)
		modelJson := ModelServices.GetModelFile(c, tableJson.Model)

		if tableJson.Wheres != nil && len(tableJson.Wheres) > 0 {
			for _, v := range tableJson.Wheres {
				if !strings.Contains(v.Field, ".") {
					v.Field = modelJson.Table + "." + v.Field
				}
				switch strings.ToUpper(v.Match) {
				case "=", "!=", "<>", ">", "<", ">=", "<=":
					db.Where(v.Field+" "+v.Match+" ?", v.Value)
				case "IN":
					db.Where(v.Field+" IN ?", strings.Split(v.Value, ","))
				case "LIKE":
					db.Where(v.Field+" LIKE ?", "%"+v.Value+"%")
				case "LIKE.LEFT":
					db.Where(v.Field+" LIKE ?", "%"+v.Value)
				case "LIKE.RIGHT":
					db.Where(v.Field+" LIKE ?", v.Value+"%")
				case "BETWEEN":
					values := strings.Split(v.Value, "~")
					db.Where(v.Field+" BETWEEN ? AND ?", values[0], values[1])
				case "IS":
					switch strings.ToUpper(v.Value) {
					case "NULL":
						db.Where(v.Field + " IS NULL")
					case "NOTNULL":
						db.Where(v.Field + " IS NOT NULL")
					}
				}
			}
		}

		if modelJson.Wheres != nil && len(modelJson.Wheres) > 0 {
			for _, v := range modelJson.Wheres {
				if !strings.Contains(v.Field, ".") {
					v.Field = modelJson.Table + "." + v.Field
				}
				switch strings.ToUpper(v.Match) {
				case "=", "!=", "<>", ">", "<", ">=", "<=":
					db.Where(v.Field+" "+v.Match+" ?", v.Value)
				case "IN":
					db.Where(v.Field+" IN ?", strings.Split(v.Value, ","))
				case "LIKE":
					db.Where(v.Field+" LIKE ?", "%"+v.Value+"%")
				case "LIKE.LEFT":
					db.Where(v.Field+" LIKE ?", "%"+v.Value)
				case "LIKE.RIGHT":
					db.Where(v.Field+" LIKE ?", v.Value+"%")
				case "BETWEEN":
					values := strings.Split(v.Value, "~")
					db.Where(v.Field+" BETWEEN ? AND ?", values[0], values[1])
				case "IS":
					switch strings.ToUpper(v.Value) {
					case "NULL":
						db.Where(v.Field + " IS NULL")
					case "NOTNULL":
						db.Where(v.Field + " IS NOT NULL")
					}
				}
			}
		}
		return db
	}
}

// Search 搜索条件
// 规则1: 表名.字段@条件=值 (如果未 join 表名可省略)
// 规则2: 表名.字段@条件=null
// 规则3: 表名.字段@条件=notnull
// id@==100000000&users.name@like.left=test&users.deleted_at@is=notnull
//
//	@receiver s
//	@param c
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sDbServices) Search(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		search := c.Request.URL.Query()
		for k, v := range search {
			keys := strings.Split(k, "@")
			if strings.Contains(v[0], "=") {
				keys[1] += "="
				v[0] = strings.TrimLeft(v[0], "=")
			}
			switch strings.ToUpper(keys[1]) {
			case "=", "!=", "<>", ">", "<", ">=", "<=":
				db.Where(keys[0]+" "+keys[1]+" ?", v[0])
			case "IN":
				db.Where(keys[0]+" IN ?", strings.Split(v[0], ","))
			case "LIKE":
				db.Where(keys[0]+" LIKE ?", "%"+v[0]+"%")
			case "LIKE.LEFT":
				db.Where(keys[0]+" LIKE ?", "%"+v[0])
			case "LIKE.RIGHT":
				db.Where(keys[0]+" LIKE ?", v[0]+"%")
			case "BETWEEN":
				values := strings.Split(v[0], "~")
				db.Where(keys[0]+" BETWEEN ? AND ?", values[0], values[1])
			case "IS":
				switch strings.ToUpper(v[0]) {
				case "NULL":
					db.Where(keys[0] + " IS NULL")
				case "NOTNULL":
					db.Where(keys[0] + " IS NOT NULL")
				}
			}
		}
		return db
	}
}
