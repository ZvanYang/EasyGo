package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := httprouter.ParamsFromContext(r.Context())

	fmt.Fprintf(w, "hello, %s!\n", params.ByName("name"))
}
