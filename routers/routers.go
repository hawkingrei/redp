package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/redp/routers/middleware/auth"
	"github.com/hawkingrei/redp/server"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	authMiddleware := &auth.GinAuthMiddleware{}
	e.Use(gin.Recovery())
	e.Use(middleware...)
	hongbao := e.Group("/api/hongbao")
	hongbao.Use(authMiddleware.MiddlewareFunc())
	{
		hongbao.GET("", server.GetAllHongbaoInfo)
		hongbao.POST("", server.CreateSendedHongbao)
		hongbao.GET("/:pid", server.GrabHongbao)
	}
	user := e.Group("/api/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.GET("", server.GetUser)
	}
	version := e.Group("/api/version")
	{
		version.GET("", server.GetVersion)
	}
	return e
}
