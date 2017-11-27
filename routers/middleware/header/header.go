package header

import (
	"net/http"
	"time"
	"github.com/hawkingrei/redp/internal/version"

	"github.com/gin-gonic/gin"
)

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Version is a middleware function that appends the Redp
// version information to the HTTP response. This is intended
// for debugging and troubleshooting.
func Version(c *gin.Context) {
	c.Header("X-REDP-VERSION", version.GitCommit)
	c.Next()
}
