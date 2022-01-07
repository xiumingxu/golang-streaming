package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router { //都是goroutine来处理的
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login) //这里是create user 而不是login
	return router
}
func main() {
	r := RegisterHandlers()
	//注册的函数
	http.ListenAndServe(":8000", r)
}
