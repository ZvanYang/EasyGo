package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func TimeMiddleware(next Handler) Handler {
	return HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		logger.Println("print message:", timeElapsed)
	})
}
