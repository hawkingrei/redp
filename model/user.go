package model

import "github.com/gin-gonic/gin"

type User struct {
	Uid      int    `gorm:"AUTO_INCREMENT;primary_key"`
	Username string `gorm:"not null;unique"`
	Memory   float32
}

func GetUser(c *gin.Context) {

}
