package middleware

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func Handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %s", ps.ByName("name"))
}

func Ware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)

		// call registered handler
		n(w, r, ps)
	}
}

func TimeMiddleware(h http.Handler) httprouter.Handle {
	return Ware(func(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		timeStart := time.Now()
		//p := httprouter.ParamsFromContext(r.Context())
		//if len(p) > 0 {
		//	ctx := r.Context()
		//	ctx = context.WithValue(ctx, httprouter.ParamsKey, p)
		//	r = r.WithContext(ctx)
		//}

		// next handler
		h.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		logger.Println("print message:", timeElapsed)
	})
}
