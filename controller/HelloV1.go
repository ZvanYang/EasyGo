package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HelloV1(w http.ResponseWriter, r *http.Request) http.Handler {
	params := httprouter.ParamsFromContext(r.Context())

	fmt.Fprintf(w, "hello, %s!\n", params.ByName("name"))
	fmt.Println("hello worldV1")
	return nil
}
