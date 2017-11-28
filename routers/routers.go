package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/redp/model"
	"github.com/hawkingrei/redp/routers/middleware/auth"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	authMiddleware := &auth.GinAuthMiddleware{}
	e.Use(gin.Recovery())
	e.Use(middleware...)
	hongbao := e.Group("/api/hongbao")
	hongbao.Use(authMiddleware.MiddlewareFunc())
	{
		hongbao.GET("", model.GetAllHongbaoInfo)
		hongbao.POST("", model.CreateHongbao)
		hongbao.GET("/:pid", model.GrabHongbao)
	}
	user := e.Group("/api/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.GET("/:uid")
	}
	version := e.Group("/api/version")
	{
		version.GET("", model.GetVersion)
	}
	return e
}
