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
	conf.HBtimeout = 5
	store_ := CreateStote(&conf)
	handler := CreateHttpHandler(store_, &conf)

	r := gofight.New()
	r.GET("/api/version").
		SetDebug(true).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
	r = gofight.New()
	r.GET("/api/user").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
	r = gofight.New()
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
	r = gofight.New()
	r.POST("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"money":     "1",
			"num":       "1000",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	r = gofight.New()
	r.POST("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"money":     "1",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	r = gofight.New()
	r.POST("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"money":     "0",
			"num":       "1",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	r = gofight.New()
	r.POST("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"money":     "1",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	r = gofight.New()
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
	r = gofight.New()
	r.GET("/api/hongbao/"+strconv.FormatInt(hbid, 10)).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wwz:e235ac07af7a969a52bec0985f6a9f85",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code, r.Body.String())
		})
	r = gofight.New()
	r.GET("/api/hongbao/"+strconv.FormatInt(hbid, 10)).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"Password":  password,
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code, r.Body.String())
		})
	r = gofight.New()
	r.GET("/api/hongbao/1000").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"Password":  password,
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})
	r = gofight.New()
	r.GET("/api/hongbao/"+strconv.FormatInt(hbid, 10)).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
			"Password":  password,
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code)
		})

	r = gofight.New()
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
	time.Sleep(3 * time.Second)
	for i := 0; i <= 10; i = i + 1 {
		store_.Background(1)
	}
	store_.Close()
}

func TestDatabase(t *testing.T) {
	var conf conf.Configure
	conf.DbDriver = "mysql"
	conf.DbURL = "root:123123@/redp?charset=utf8&parseTime=True&loc=Local"
	conf.Debug = true
	conf.HBtimeout = 5
	_, err := setupStore(&conf)
	assert.NotEqual(t, nil, err)
}

func TestGrabHongbao1(t *testing.T) {
	var conf conf.Configure
	conf.DbDriver = "mysql"
	conf.DbURL = "root:@/redp?charset=utf8&parseTime=True&loc=Local"
	conf.Debug = true
	conf.HBtimeout = 5
	store_ := CreateStote(&conf)
	handler := CreateHttpHandler(store_, &conf)
	var gotmoney float64
	gotmoney = 0
	var password string
	var hbid int64
	r := gofight.New()

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
			hbid, _ = jsonparser.GetInt(data, "Money")
		})
	r = gofight.New()
	r.GET("/api/hongbao/"+strconv.FormatInt(hbid, 10)).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wwz:e235ac07af7a969a52bec0985f6a9f85",
			"Password":  password,
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code, r.Body.String())
			data := []byte(r.Body.String())
			gotmoney, _ = jsonparser.GetFloat(data, "Hbid")
		})
	time.Sleep(5 * time.Second)
	store_.Background(1)
	r = gofight.New()
	r.GET("/api/hongbao/"+strconv.FormatInt(hbid, 10)).
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "a:0cc175b9c0f1b6a831c399e269772661",
			"Password":  password,
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.NotEqual(t, http.StatusOK, r.Code, r.Body.String())
		})
	r = gofight.New()
	r.GET("/api/hongbao").
		SetDebug(true).
		SetHeader(gofight.H{
			"Signature": "wz:d0965c07d1a00fcc85d28b8a241ae35a",
		}).
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			data := []byte(r.Body.String())
			hasmoney, _ := jsonparser.GetFloat(data, "Money")
			assert.Equal(t, float64(10), hasmoney+gotmoney)
		})
	store_.Close()
}
