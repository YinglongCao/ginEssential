package common

import (
	"fmt"
	"ginEssential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	// 初始化连接池

	// args
	host := "localhost"
	port := "33066"
	user := "root"
	password := "root"
	database := "ginEssential"
	charset := "utf8mb4"

	// DSN
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, database, charset)

	// 连接数据库
	var err error
	// 赋值本包db变量
	db, err = gorm.Open(mysql.Open(args), &gorm.Config{})

	if err != nil {
		panic("database connect err:" + err.Error())
	}

	// 自动创建数据库字段
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil
	}

	return db

}

func GetDB() *gorm.DB {
	fmt.Println(db)
	return db
}
