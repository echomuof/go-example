/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(writer http.ResponseWriter, request *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "_" + request.URL.Path
	if handlerFunc, ok := engine.router[key]; ok {
		handlerFunc(writer, request)
	} else {
		fmt.Fprintf(writer, "404 NOT FOUND %q\n", request.URL)
	}
}

func (engine *Engine) Run(addr string) error {
	fmt.Printf("%p\n",engine)
	return http.ListenAndServe(addr, engine)
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	fmt.Printf("%p\n",engine)
	engine.addRoute("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	fmt.Printf("%p\n",engine)
	engine.addRoute("POST", path, handler)
}

func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	fmt.Printf("%p\n",engine)
	key := method + "_" + path
	engine.router[key] = handler
}
