package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

/*
@Author: 71made
@Date: 2023/02/19 13:29
@ProductName: init.go
@Description: 建立数据库连接
*/

var Db *gorm.DB

// InitDB 初始化DB
func InitDB() {
	// 构建 MySQL 数据库连接
	var err error
	Db, err = gorm.Open(
		mysql.Open("root:123456@tcp(localhost:3306)/qingxun?charset=utf8&parseTime=True&loc=Local&clientFoundRows=true"),
		&gorm.Config{
			PrepareStmt: true,
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags|log.Lmicroseconds), logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
				LogLevel:                  logger.Info,
			}),
		},
	)
	if err != nil {
		panic(err)
	}

	// 使用 tracing
	//if err := db.Use(tracing.NewPlugin()); err != nil {
	//	panic(err)
	//}
}
func GetInstance() *gorm.DB {
	if Db == nil {
		panic("date sources is missing")
	}

	return Db
}
