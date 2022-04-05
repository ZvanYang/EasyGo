package middleware

import "net/http"

type loggerMiddleware func(handler http.Handler) http.Handler

type loggerHandlerFunc func(http.ResponseWriter, *http.Request)

func (f loggerHandlerFunc) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	f(wr, r)
}