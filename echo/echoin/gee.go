/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package echoin

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

type HandlerFunc func(ctx *Context)

// 路由组
type RouterGroup struct {
	//组前缀
	prefix      string
	//中间件handler
	middlewares []HandlerFunc
	//父组
	parent      *RouterGroup

	engine      *Engine
}

type Engine struct {
	*RouterGroup
	r      *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{r: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	//统一异常处理
	engine.Use(Recovery())
	return engine
}

/*
 * 添加一个group
 */
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	//最上级的group
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	//添加到group列表中 /v1、/v2、/v1/v3
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

/*
 * 为group绑定中间件
 */
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

/*
 * 接收到请求后的逻辑
 */
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		//遍历所有group，如果当前请求属于某个group，就在其context的处理器列中加上该group的中间件
		if strings.HasPrefix(request.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	ctx := newContext(writer, request)
	ctx.handlers = middlewares
	//执行请求的处理逻辑
	engine.r.handle(ctx)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (group *RouterGroup) addRoute(method string, path string, handler HandlerFunc) {
	pattern := group.prefix + path
	group.engine.r.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(path string, handler HandlerFunc) {
	group.addRoute("GET", path, handler)
}

func (group *RouterGroup) POST(path string, handler HandlerFunc) {
	group.addRoute("POST", path, handler)
}

/*
 * 打印异常堆栈
 */
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

/*
 * 统一异常处理
 */
func Recovery() HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				ctx.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		ctx.Next()
	}
}
