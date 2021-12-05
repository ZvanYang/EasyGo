package main

import (
	"EasyGo/controller"
	"EasyGo/middleware"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/http/pprof"
)

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %s", ps.ByName("name"))
}

// middleware is used to intercept incoming HTTP calls and apply general functions upon them.
func middleware1(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)

		// call registered handler
		n(w, r, ps)
	}
}

func main() {

	all := middleware.NestedMiddleware(middleware.WithLogging, middleware.WithTracing)

	router := httprouter.New()

	// 需要默认使用controller下面的函数，类似echo
	//router.GET("/", )

	//r := middleware.NewRouter()
	//r.Use(middleware.TimeMiddlewareV2)
	//r.Use(middleware.BasicAuth)
	//r.Add("/hello/:name", controller.HelloV1Handler)

	//router.GET("/hellome", middleware.TimeMiddleware(http.HandlerFunc(controller.HelloV2)))
	router.GET("/hello/:name", middleware.TimeMiddleware(http.HandlerFunc(controller.HelloV2)))


	router.GET("/welcome/:name", middleware.TimeMiddleware(all(controller.HelloV1)))


	//router.GET("/hello/:name", middleware1(handler))

	router.GET("/debug/pprof/goroutine", middleware.TimeMiddleware(pprof.Handler("goroutine")))
	//router.Handler(http.MethodPost, "/UpdateUsername", middleware.TimeMiddleware(middleware.HandlerFunc(controller.UpdateUsername)))
	//http.Handle("/helloV2", middleware.TimeMiddleware(middleware.HandlerFunc(controller.HelloV2)))
	//http.Handle("/helloV1", middleware.TimeMiddleware(middleware.HandlerFunc(controller.HelloV1)))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}
}
