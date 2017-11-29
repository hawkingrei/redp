package model

import (
	"github.com/gin-gonic/gin"
)

type Hongbao struct {
	Hbid  int64
	Uid   int64
	Money float32
	Num   int
	Type  int // 0: sended redp 1： got redp
}

func CreateHongbao(c *gin.Context) {

}

func GetAllHongbaoInfo(c *gin.Context) {

}

func ListHongbao(c *gin.Context) {

}

func GrabHongbao(c *gin.Context) {

}
