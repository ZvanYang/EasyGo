package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HelloV2(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	fmt.Fprintf(w, "hello, %s!\n", params.ByName("name"))
	fmt.Println("hello worldV2")
}
