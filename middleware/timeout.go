package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

var Logger = log.New(os.Stdout, "", 0)

func TimeOut(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(w, r)

		timeElapsed := time.Since(timeStart)
		Logger.Println("print message:", timeElapsed)
	}

	return http.HandlerFunc(fn)
}
