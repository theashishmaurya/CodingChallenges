package main

// TOODO: FIX the types and functions to use our custom HTTP server
import (
	"net/http"
	"strings"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(method, pattern string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{method, pattern, handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && strings.HasPrefix(req.URL.Path, route.Pattern) {
			route.Handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
