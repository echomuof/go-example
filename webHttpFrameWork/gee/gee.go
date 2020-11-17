/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package gee

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	r *router
}

func New() *Engine {
	return &Engine{r: newRouter()}
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	engine.r.handle(newContext(writer, request))
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	engine.r.addRoute(method, path, handler)
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRoute("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRoute("POST", path, handler)
}
