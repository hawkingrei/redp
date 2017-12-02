package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/version", nil)
	handler.ServeHTTP(w, r)
	if r.Response.StatusCode != 200 {
		t.Error("GET /api/version not 200")
	}

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/api/user", nil)
	r.Header.Add("Signature", "wz:d0965c07d1a00fcc85d28b8a241ae35a")
	handler.ServeHTTP(w, r)
	if r.Response.StatusCode != 200 {
		t.Error("GET /api/user not 200")
	}

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/api/user", nil)
	r.Header.Add("Signature", "wz:d0965c07d1a00fcc85d28b8a241aa")
	handler.ServeHTTP(w, r)
	if r.Response.StatusCode != 200 {
		t.Error("GET /api/user not 200")
	}
	w = httptest.NewRecorder()
	r, _ = http.NewRequest("POST", "/api/hongbao", nil)
	r.Header.Add("Signature", "wz:d0965c07d1a00fcc85d28b8a241aa")
	r.Header.Add("money", "10")
	r.Header.Add("num", "1")
	handler.ServeHTTP(w, r)
	if r.Response.StatusCode != 200 {
		t.Error("POST /api/hongbao not 200")
	}
	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/api/hongbao/1", nil)
	r.Header.Add("Signature", "wz:d0965c07d1a00fcc85d28b8a241aa")
	handler.ServeHTTP(w, r)
	if r.Response.StatusCode != 200 {
		t.Error("GET /api/hongbao not 200")
	}

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/api/hongbao", nil)
	r.Header.Add("Signature", "wz:d0965c07d1a00fcc85d28b8a241aa")
	handler.ServeHTTP(w, r)
	if r.Response.StatusCode != 200 {
		t.Error("GET /api/hongbao not 200")
	}
}
