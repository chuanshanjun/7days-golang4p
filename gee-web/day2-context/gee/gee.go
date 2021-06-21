package gee

import "net/http"

type HandlerFunc func(c *Context)

type Engine struct {
	routers *Router
}

func NEW() *Engine {
	return &Engine{routers: newRouter()}
}

func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.routers.addRouter(method, pattern, handler)
}

func (e *Engine) GET(patter string, handler HandlerFunc) {
	e.addRouter("GET", patter, handler)
}

func (e *Engine) POST(patter string, handler HandlerFunc) {
	e.addRouter("POST", patter, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.routers.handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}