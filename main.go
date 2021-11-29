package main

import (
	"EasyGo/middleware"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)



func main() {
	http.Handle("/", middleware.TimeMiddleware(middleware.HandlerFunc(hello)))
	http.Handle("/", middleware.TimeMiddleware(middleware.HandlerFunc(HelloV2)))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("error message:", err)
	}
}

func hello(wr http.ResponseWriter, r *http.Request) {
	_, err := wr.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}

func HelloV2(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}


