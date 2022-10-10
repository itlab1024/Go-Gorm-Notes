package main

import "Go-Gorm-Notes/models"

func main() {
	// 使用AutoMigrate自动创建数据库表
	//models.Db.AutoMigrate(&models.Author{})
	//models.Db.Create(&models.Author{
	//	Name:          "张飞",
	//	Identify:      "001",
	//	ByteJson:      []byte("字节切片"),
	//	ConcatWayJSON: models.ConcatWay{Address: "https://itlab1024.com", Email: "itlab1024@163.com"},
	//	ConcatWayGob:  models.ConcatWay{Address: "https://itlab1024.com", Email: "itlab1024@163.com"},
	//	TimeUnixtime:  12,
	//})
	//关联
	//models.Db.AutoMigrate(&models.User{}, &models.Country{})
	models.Db.AutoMigrate(&models.User{}, &models.Language{})
}
