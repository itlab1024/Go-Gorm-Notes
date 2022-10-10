package models

import "time"

type Author struct {
	ID   uint   `gorm:"primarykey;comment:主键ID"`
	Name string `gorm:"column:name;type:varchar(200) not null default '';index"`
	Sex  string `gorm:"size:10;default:'男'"`
	// 身份信息唯一
	Identify      string    `gorm:"size:100;uniqueIndex;not null;"`
	CtTimeNano    time.Time `gorm:"autoCreateTime:nano"` //nano/milli
	CtTimeMilli   time.Time `gorm:"autoCreateTime:milli"`
	CtNano        int       `gorm:"autoCreateTime:nano"`
	CtMilli       int       `gorm:"autoCreateTime:milli"`
	ConcatWay     `gorm:"embedded;embeddedPrefix:cw_"`
	ByteJson      []byte    `gorm:"serializer:json"`
	ConcatWayJSON ConcatWay `gorm:"serializer:json"`
	ConcatWayGob  ConcatWay `gorm:"type:bytes;serializer:gob"`
	TimeUnixtime  int64     `gorm:"serializer:unixtime;type:time"` //将int64的内容转化为ddatetime存储
}

// ConcatWay 联系方式
type ConcatWay struct {
	Address string
	Phone   string
	Email   string
}

// TableName 自定义表名
func (Author) TableName() string {
	return "author"
}
