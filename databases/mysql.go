package databases

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	gormDB *gorm.DB
	err    error
)

// InitMysql 连接数据库
//
//	@return *gorm.DB
func InitMysql() *gorm.DB {
	if gormDB == nil {
		// 加载.env文件
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return nil
		}
		// 使用环境变量中的值
		db := os.Getenv("DB_DATABASE")
		port := os.Getenv("DB_PORT")
		host := os.Getenv("DB_HOST")
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
		logfile, _ := os.OpenFile("gorm.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.New(
				log.New(logfile, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Info,
					Colorful:      false,
				},
			),
		})

		if err != nil {
			fmt.Println("数据库连接失败", err)
			return nil
		}
	}

	return gormDB
}

// CloseMysql 关闭数据库
func CloseMysql() {
	sqlDB, err := gormDB.DB()
	if err != nil {
		fmt.Println("获取数据库失败", err)
		return
	}
	sqlDB.Close()
}
