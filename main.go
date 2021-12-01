package main

import (
	"EasyGo/controller"
	"EasyGo/middleware"
	"fmt"
	"net/http"
)

func main() {


	//router := httprouter.New()

	// 需要默认使用controller下面的函数，类似echo
	//router.GET("/", )
	//router.Handler(http.MethodGet, "/hello", middleware.TimeMiddleware(middleware.HandlerFunc(controller.HelloV2)))

	//router.Handler(http.MethodPost, "/UpdateUsername", middleware.TimeMiddleware(middleware.HandlerFunc(controller.UpdateUsername)))
	http.Handle("/helloV2", middleware.TimeMiddleware(middleware.HandlerFunc(controller.HelloV2)))
	http.Handle("/helloV1", middleware.TimeMiddleware(middleware.HandlerFunc(controller.HelloV1)))


	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		fmt.Println("error message:", err)
	}
}
