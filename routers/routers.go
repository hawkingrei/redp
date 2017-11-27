package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/redp/model"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(middleware...)
	hongbao := e.Group("/api/hongbao")
	{
		hongbao.GET("", model.GetHongbao)
		hongbao.POST("", model.CreateHongbao)
		hongbao.GET("/:pid")
		hongbao.POST("/:pid")
	}
}
