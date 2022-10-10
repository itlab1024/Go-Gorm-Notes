package models

import "gorm.io/gorm"

// User type User struct {
//	gorm.Model
//	Name      string
//	CountryId uint
//	// 这里不能使用匿名
//	Country Country
//}
//type Country struct {
//	*gorm.Model
//	Name string
//}
//
//// User 有一张 CreditCard，UserID 是外键
//type User struct {
//	gorm.Model
//	CreditCard CreditCard
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number string
//	UserID uint
//}

// User 有多张 CreditCard，UserID 是外键
//type User struct {
//	gorm.Model
//	CreditCards []CreditCard
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number string
//	UserID uint
//}

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
	UserRef   uint
}

type Language struct {
	gorm.Model
	Name string
}
