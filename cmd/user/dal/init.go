package dal

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
 @Author: 71made
 @Date: 2023/02/19 13:29
 @ProductName: init.go
 @Description: 建立数据库连接
*/

var DB *gorm.DB

// 创建数据库连接池
func InitDB(v *viper.Viper) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		v.GetString("mysql.user"),
		v.GetString("mysql.password"),
		v.GetString("mysql.host"),
		v.GetInt("mysql.port"),
		v.GetString("mysql.dbname"),
		v.GetString("mysql.charset"),
		v.GetBool("mysql.parseTime"),
		v.GetString("mysql.loc"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return DB
}

// 初始化config文件
func InitConfig() *viper.Viper {
	workDir, _ := os.Getwd()
	var v = viper.New()
	// 读取的文件名
	v.SetConfigName("dbConfig.yml")
	// 读取的文件类型
	v.SetConfigType("yml")
	// 读取的路径
	v.AddConfigPath(workDir + "\\config")

	err := v.ReadInConfig()
	if err != nil {
		// fmt.Println(workDir)
		fmt.Println(err)
		panic("Read dbConfig faild!")
	}
	return v
}

// 获取DB的示例
func GetDB() *gorm.DB {
	return DB
}

func Init() {
	var v = InitConfig()
	// 尝试连接数据库
	db := InitDB(v)
	sqlDB, err := db.DB()
	if err != nil {
		panic("sqlDB error:" + err.Error())
	}

	if err := sqlDB.Ping(); err != nil {
		panic("sqlDB ping error:" + err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
