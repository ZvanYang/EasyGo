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
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("error message:", err)
	}
}

func hello(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	_, err := wr.Write([]byte("Hello World"))
	if err != nil {
		return 
	}

	timeElapsed := time.Since(timeStart)
	fmt.Println("print message:", timeElapsed)
}
