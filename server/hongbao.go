package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/redp/store"
)

func CreateSendedHongbao(c *gin.Context) {
	username, _ := c.Get("user")
	strNum := c.Request.Header.Get("num")
	if strNum == "" {
		c.JSON(500, "HTTP HEADER num is empty")
		return
	}
	num, err := strconv.Atoi(strNum)
	if err != nil && num > 0 {
		c.JSON(500, "HTTP HEADER num is mistake")
	}
	strMoney := c.Request.Header.Get("money")
	if strNum == "" {
		c.JSON(500, "HTTP HEADER num is empty")
		return
	}
	money, err := strconv.ParseFloat(strMoney, 32)
	if err != nil && money > 0 {
		c.JSON(500, "HTTP HEADER money is mistake")
	}
	hb, err := store.CreateSendedHongbao(c, username.(string), float32(money), num)
	if err != nil {
		c.String(500, "Error send hongbao. %s", err)
		return
	}
	c.JSON(200, *hb)
}

func GetAllHongbaoInfo(c *gin.Context) {

}

func ListHongbao(c *gin.Context) {

}

func GrabHongbao(c *gin.Context) {

}
