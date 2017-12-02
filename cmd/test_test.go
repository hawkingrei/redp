package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hawkingrei/redp/conf"
)

func TestAllHeader(t *testing.T) {
	var conf conf.Configure
	conf.DbDriver = "mysql"
	conf.DbURL = "root:@/redp?charset=utf8&parseTime=True&loc=Local"
	store_ := CreateStote(&conf)
	handler := CreateHttpHandler(store_, &conf)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/version", nil)
	handler.ServeHTTP(w, r)

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/api/user", nil)
	r.Header.Add("Signature", "wz:d0965c07d1a00fcc85d28b8a241ae35a")
	handler.ServeHTTP(w, r)

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/api/user", nil)
	r.Header.Add("Signature", "wz:d0965c07d1a00fcc85d28b8a241aa")
	handler.ServeHTTP(w, r)

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("POST", "/api/user", nil)
	r.Header.Add("Signature", "wz:d0965c07d1a00fcc85d28b8a241aa")
	r.Header.Add("money", "10")
	r.Header.Add("num", "1")
	handler.ServeHTTP(w, r)

}
