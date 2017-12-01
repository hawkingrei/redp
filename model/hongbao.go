package model

import (
	"time"
)

type SendedHongbao struct {
	Hbid       int64 `gorm:"AUTO_INCREMENT;primary_key"`
	Username   string
	Money      float32 `sql:"type:decimal(10,2);"`
	Num        int
	Password   string
	Closed     int "0: open 1: closed"
	CreateTime time.Time
}

type GotHongbao struct {
	Gothbid  int64 `gorm:"AUTO_INCREMENT;primary_key"`
	Hbid     int64
	Username string
	Money    float32 `sql:"type:decimal(10,2);"`
}
