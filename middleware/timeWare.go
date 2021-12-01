package middleware

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func TimeMiddleware(next Handler) Handler {
	return HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		p := httprouter.ParamsFromContext(r.Context())
		if len(p) > 0 {
			ctx := r.Context()
			ctx = context.WithValue(ctx, httprouter.ParamsKey, p)
			r = r.WithContext(ctx)
		}

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		logger.Println("print message:", timeElapsed)
	})
}
