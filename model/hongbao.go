package model

import (
	"time"
)

type Hongbao struct {
	Hbid       int64 `gorm:"AUTO_INCREMENT;primary_key"`
	Uid        int64
	Money      float32
	Num        int
	Type       int // 0: sended redp 1： got redp
	CreateTime time.Time
}
