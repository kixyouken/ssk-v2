package routers

import (
	"ssk-v2/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("api")
	{
		// 所有
		api.GET("/table/:model/get", controllers.Get)
		// 分页
		api.GET("/table/:model/page", controllers.Page)
		// 详情
		api.GET("/form/:model/:id", controllers.Read)
		// 新增
		api.POST("/form/:model", controllers.Save)
		// 修改
		api.PUT("/form/:model/:id", controllers.Update)
		// 删除
		api.DELETE("/table/:model/:id", controllers.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong~pong~",
		})
	})

	return r
}
