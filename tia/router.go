package tia

import (
	"net/http"
	"log"
)
type router struct {
	handlers map[string]HandlerFunc
}


func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}
func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	log.Printf("Add route %4s - %s", method, path)
	key := method + "-" + path
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}