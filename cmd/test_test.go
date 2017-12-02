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
	store_ := CreateStote(conf)
	handler := CreateHttpHandler(store_, &conf)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/version", nil)
	handler.ServeHTTP(w, r)
}
