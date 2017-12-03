package main

import (
	"net/http"
	"strconv"
	"testing"
	"time"

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
	var password string
	var hbid int64
	r.POST("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"money":     "10",
			"num":       "1",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			data := []byte(r.Body.String())
			password, _ = jsonparser.GetString(data, "Password")
			hbid, _ = jsonparser.GetInt(data, "Hbid")

		})
	r.GET("/api/hongbao/"+strconv.FormatInt(hbid, 10)).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wwz:e235ac07af7a969a52bec0985f6a9f85",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	r.GET("/api/hongbao/"+strconv.FormatInt(hbid, 10)).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	r.GET("/api/hongbao/"+strconv.FormatInt(hbid, 10)).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"Password":  password,
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code, r.Body.String())
		})
	r.GET("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
	r = gofight.New()
	r.GET("/api/hongbao").
		SetDebug(true).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	time.Sleep(10)
	for i := 0; i <= 10; i = i + 1 {
		store_.Background(conf.HBtimeout)
	}
}
