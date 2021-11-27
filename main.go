package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("error message:", err)
	}
}

func hello(wr http.ResponseWriter, r *http.Request)  {
	write, err := wr.Write([]byte("Hello World"))
	if err != nil {
		return 
	}
}