package gee

import (
	"fmt"
	"net/http"
)

type HandlerFun func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFun
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFun) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handler HandlerFun) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFun) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %s\n", r.URL)
	}
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFun)}
}
