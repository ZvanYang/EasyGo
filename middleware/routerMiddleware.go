package middleware

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %s", ps.ByName("name"))
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	f(wr, r)
}

func Ware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)
		n(w, r, ps)
	}
}

// 想做的是一个路由识别的工具
type middleware func(handler http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		middlewareChain: nil,
		mux:             make(map[string]http.Handler),
	}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {

	var mergeHandler = h

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergeHandler = r.middlewareChain[i](mergeHandler)
	}

	r.mux[route] = mergeHandler
}
