package main

import (
	"net/http"
	"testing"

	"github.com/WindomZ/testify/assert"
	"github.com/appleboy/gofight"
	"github.com/hawkingrei/redp/conf"
)

func TestSimpleApi(t *testing.T) {
	var conf conf.Configure
	conf.DbDriver = "mysql"
	conf.DbURL = "root:@/redp?charset=utf8&parseTime=True&loc=Local"
	conf.Debug = true
	conf.HBtimeout = 3
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

	r.POST("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"money":     "10",
			"num":       "10",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})

	r.GET("/api/hongbao/1").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)

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
