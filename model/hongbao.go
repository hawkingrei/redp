package model

import (
	"time"
)

type SendedHongbao struct {
	Hbid       int64 `gorm:"AUTO_INCREMENT;primary_key"`
	Username   string
	Money      float32
	Num        int
	Password   string
	CreateTime time.Time
}

type GotHongbao struct {
	Gothbid  int64 `gorm:"AUTO_INCREMENT;primary_key"`
	Hbid     int64
	Username string
	Money    float32
}

type AllHongbao struct {
	GotHongbaos   []GotHongbao
	SendedHongbao []SendedHongbao
}
