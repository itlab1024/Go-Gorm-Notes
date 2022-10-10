package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db = initDb()

//初始化数据库连接
func initDb() *gorm.DB {
	// Mysql的链接字符串,我电脑上用户名是root，密码是qwe!@#123
	dsn := "root:qwe!@#123@tcp(127.0.0.1:3306)/go-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// Mysql驱动有自己的配置选项，可以通过 mysql.New(mysql.Config{})配置。具体可看mysql.Config。
	Db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return Db
}
