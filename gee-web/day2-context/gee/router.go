package gee

import (
	"log"
	"net/http"
)

type Router struct {
	hadnlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{hadnlers: make(map[string]HandlerFunc)}
}

func (r *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.hadnlers[key] = handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Req.URL.Path
	if handler, ok := r.hadnlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

