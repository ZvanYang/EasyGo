package main

import (
	"EasyGo/controller"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func main() {

	//http.Handle("/", middleware.TimeMiddleware(middleware.HandlerFunc(HelloV2)))
	
	router := httprouter.New()

	// 需要默认使用controller下面的函数，类似echo
	router.GET("/", controller.Hello)
	router.POST("/", controller.UpdateUsername)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("error message:", err)
	}
}


