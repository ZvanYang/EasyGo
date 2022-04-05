package easygo

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Router struct {
	middlewareChain []Middleware
	mux             map[string]http.Handler
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.ServeHTTP(writer, request)
}

func NewRouter() *Router {
	return &Router{
		mux: make(map[string]http.Handler),
	}
}

func (r *Router) Use(m Middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergeHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergeHandler = r.middlewareChain[i](mergeHandler)
	}

	r.mux[route] = mergeHandler
}
