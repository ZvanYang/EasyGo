package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HelloV1(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	_, err := fmt.Fprintf(w, "hello, %s!\n", params.ByName("name"))
	if err != nil {
		return 
	}
	fmt.Println("hello worldV1")
}
