package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB 初始化连接池
func InitDB() *gorm.DB {

	// args
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	user := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")

	// DSN
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, database, charset)

	// 连接数据库
	var err error
	// 赋值本包db变量
	db, err = gorm.Open(mysql.Open(args), &gorm.Config{})

	if err != nil {
		panic("InitDB() 数据库连接错误" + err.Error())
	}

	// 自动创建数据表
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil
	}

	return db

}

func GetDB() *gorm.DB {
	if db == nil {
		panic("GetDB() db is nil")
	}
	return db
}
