package easygo

import (
	"EasyGo/middleware"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	router.Use(middleware.WithTracing)
	router.Use(middleware.TimeOut)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}
}
