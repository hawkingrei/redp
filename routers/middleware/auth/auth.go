package auth

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/redp/store"
	"github.com/sirupsen/logrus"
)

type GinAuthMiddleware struct {
	// Callback function that should perform the authorization of the authenticated user. Called
	// only after an authentication success. Must return true on success, false on failure.
	// Optional, default to success.
	Authorizator func(c *gin.Context) bool
	// User can define own Unauthorized func.
	Unauthorized func(*gin.Context, int, string)
}

func (mw *GinAuthMiddleware) MiddlewareInit() error {
	if mw.Authorizator == nil {
		mw.Authorizator = func(c *gin.Context) bool {

			takon, ok := c.Get("takon")
			if !ok {
				return false
			}
			user, ok := c.Get("user")
			if !ok {
				return false
			}
			md5user := md5.New()
			io.WriteString(md5user, user.(string))
			digest := fmt.Sprintf("%x", md5user.Sum(nil))
			if digest == takon.(string) {
				return true
			}
			return false
		}
	}

	if mw.Unauthorized == nil {
		mw.Unauthorized = func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		}
	}
	return nil
}

func (mw *GinAuthMiddleware) MiddlewareFunc() gin.HandlerFunc {
	if err := mw.MiddlewareInit(); err != nil {
		return func(c *gin.Context) {
			mw.unauthorized(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	return func(c *gin.Context) {
		mw.middlewareImpl(c)
		return
	}
}

func (mw *GinAuthMiddleware) middlewareImpl(c *gin.Context) {
	user, takon, err := mw.jwtFromHeader(c, "Signature")
	if err != nil {
		mw.unauthorized(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("takon", takon)
	c.Set("user", user)

	if !mw.Authorizator(c) {
		mw.unauthorized(c, http.StatusForbidden, "You don't have permission to access.")
		return
	}
	if !store.HasUser(c, user) {
		store.CreateUser(c, user)
		logrus.Info("create user ", user)
	}
	c.Next()
}

func (mw *GinAuthMiddleware) jwtFromHeader(c *gin.Context, key string) (string, string, error) {
	authHeader := c.Request.Header.Get(key)
	if authHeader == "" {
		return "", "", errors.New("auth header empty")
	}
	UserAndTakon := strings.Split(authHeader, ":")
	return UserAndTakon[0], UserAndTakon[1], nil
}

func (mw *GinAuthMiddleware) unauthorized(c *gin.Context, code int, message string) {
	c.Abort()
	mw.Unauthorized(c, code, message)
	return
}
