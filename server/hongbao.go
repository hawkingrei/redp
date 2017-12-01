package server

import (
	"fmt"
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

func ListGotHongbao(c *gin.Context) {
	username, _ := c.Get("user")
	result, err := store.ListGotHongbao(c, username.(string))
	if err != nil {
		c.String(500, "Error list got hongbao. %s", err)
	}
	c.JSON(200, result)
}

func GrabHongbao(c *gin.Context) {
	username, _ := c.Get("user")
	pid := c.Param("pid")
	hid, err := strconv.ParseInt(pid, 10, 64)
	fmt.Println(hid)
	if err != nil {
		c.JSON(500, "HTTP parm pid is mistake")
		return
	}
	password := c.Request.Header.Get("password")
	hb, err := store.GrabHongbao(c, hid, username.(string), password)
	if err != nil {
		c.String(500, "Error grab hongbao. %s", err)
		return
	}
	c.JSON(200, *hb)
}
