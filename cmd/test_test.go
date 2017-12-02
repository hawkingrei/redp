package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/WindomZ/testify/assert"
	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/hawkingrei/redp/conf"
)

func TestSimpleApi(t *testing.T) {
	var conf conf.Configure
	conf.DbDriver = "mysql"
	conf.DbURL = "root:@/redp?charset=utf8&parseTime=True&loc=Local"
	conf.Debug = true
	conf.HBtimeout = 9
	store_ := CreateStote(&conf)
	handler := CreateHttpHandler(store_, &conf)

	r := gofight.New()
	r.GET("/api/version").
		SetDebug(true).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})

	r.GET("/api/user").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})

	r.GET("/api/user").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241aea",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	var password string
	var hbid string
	r.POST("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"money":     "10",
			"num":       "10",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			fmt.Println(r.Body.String())
			data := []byte(r.Body.String())
			password, _ = jsonparser.GetString(data, "Password")
			hbid, _ = jsonparser.GetString(data, "Hbid")

		})

	r.GET("/api/hongbao/"+hbid).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	r.GET("/api/hongbao/"+hbid).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"Password":  password,
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code, r.Body.String())
		})
	r.GET("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
