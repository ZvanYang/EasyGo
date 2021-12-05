package middleware

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

var Logger = log.New(os.Stdout, "", 0)

func TimeMiddleware(h http.Handler) httprouter.Handle {
	return Ware(func(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		timeStart := time.Now()
		p := httprouter.ParamsFromContext(r.Context())
		if len(p) > 0 {
			ctx := r.Context()
			ctx = context.WithValue(ctx, httprouter.ParamsKey, p)
			r = r.WithContext(ctx)
		}

		// next handler
		h.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		Logger.Println("print message:", timeElapsed)
	})
}


func TimeMiddlewareV2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		Logger.Println("print message:", timeElapsed)
	})
}
