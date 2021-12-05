package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func BasicAuth(h http.Handler) httprouter.Handle {
	return Ware(func(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		_, _, hasAuth := r.BasicAuth()

		if hasAuth {
			Logger.Println("auth   ....")
			// Delegate request to the given handle
			h.ServeHTTP(wr, r)
		} else {
			// Request Basic Authentication otherwise
			wr.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(wr, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	})
}
