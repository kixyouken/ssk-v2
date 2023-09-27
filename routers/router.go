package routers

import (
	"ssk-v2/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/json", "./json")

	admin := r.Group("admin")
	{
		admin.GET("/model/:model", controllers.GetModel)
		admin.GET("/table/:table", controllers.GetTable)
		admin.GET("form/:form", controllers.GetForm)
	}

	api := r.Group("api")
	{
		// 所有
		api.GET("/table/:table", controllers.Get)
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
