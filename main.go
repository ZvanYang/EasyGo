package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func main() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))

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

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		logger.Println("print message:", timeElapsed)
	})
}
