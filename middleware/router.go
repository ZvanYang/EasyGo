package middleware

import "net/http"

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
