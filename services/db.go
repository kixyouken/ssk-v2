package services

import (
	"ssk-v2/databases"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sDbServices struct{}

var DbService = sDbServices{}

var db = databases.InitMysql()

func (s *sDbServices) Get(c *gin.Context, table string, out interface{}, column interface{}, order string, join []string, search interface{}) error {
	return db.Table(table).Where(search).
		Scopes(s.Order(order), s.Joins(join...)).
		Select(column).
		Limit(30).
		Find(out).Error
}

func (s *sDbServices) Page(c *gin.Context, table string, out interface{}, column interface{}, order string, join []string, search interface{}) error {
	return db.Table(table).Where(search).
		Scopes(s.Paginate(c), s.Order(order), s.Joins(join...)).
		Select(column).
		Find(out).Error
}

func (s *sDbServices) Count(c *gin.Context, table string, join []string, search interface{}) int64 {
	var count int64
	err := db.Table(table).Where(search).
		Scopes(s.Joins(join...)).
		Count(&count).Error

	if err != nil {
		return 0
	}

	return count
}

func (s *sDbServices) Read(c *gin.Context) {

}

func (s *sDbServices) Save(c *gin.Context) {

}

func (s *sDbServices) Update(c *gin.Context) {

}

func (s *sDbServices) Delete(c *gin.Context) {

}

func (s *sDbServices) Order(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func (s *sDbServices) Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page < 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(c.Query("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func (s *sDbServices) Joins(joins ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, v := range joins {
			db.Joins(v)
		}
		return db
	}
}
