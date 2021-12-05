package middleware

import "net/http"

type rateLimitHandlerFunc func(http.ResponseWriter, *http.Request)

func (f rateLimitHandlerFunc) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	f(wr, r)
}