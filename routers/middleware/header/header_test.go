package header

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/redp/routers/middleware"
)

func TestNoCache(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(NoCache) //only use the mw I want to test
	router.Use(middleware.Version)
	router.GET("/version", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, r)
}
