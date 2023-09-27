package main

import (
	"log"
	"os"
	"ssk-v2/databases"
	"ssk-v2/routers"

	"github.com/joho/godotenv"
)

func main() {
	r := routers.GetRouter()
	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	port := os.Getenv("APP_PORT")
	// 监听并在 0.0.0.0:9000 上启动服务
	r.Run(":" + port)

	databases.CloseMysql()
}
