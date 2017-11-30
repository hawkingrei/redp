package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/redp/store"
)

func GetUser(c *gin.Context) {
	username, _ := c.Get("user")
	user, err := store.GetUser(c, username.(string))
	if err != nil {
		c.String(500, "Error fetching user. %s", err)
		return
	}
	c.JSON(200, user)
}