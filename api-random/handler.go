package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type HelloHandler struct {
}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		num := handler.randomNum()
		w.Write([]byte(fmt.Sprintf("%d", num)))
	}
}

func (handler *HelloHandler) randomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1

}
