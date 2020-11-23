package echoin

import (
	"net/http"
	"strings"
)

type router struct {
	//key-Method  value-每种method的根节点
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

/*
 * 拆分路径
 */
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

/*
 * 增加路由
 */
func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	//拆分路径
	parts := parsePattern(path)
	key := method + "_" + path
	//获取根节点
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	//插入前缀树
	r.roots[method].insert(path, parts, 0)
	r.handlers[key] = handler
}

/*
 * 获取路由
 */
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	//获取跟节点
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.wholePath)
		//替换占位符
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

/*
 * 执行请求的处理逻辑
 */
func (r *router) handle(ctx *Context) {
	n, params := r.getRoute(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Method + "_" + n.wholePath
		//获取该请求的处理方法，绑定到处理器列表中
		ctx.handlers = append(ctx.handlers, r.handlers[key])
	} else {
		ctx.handlers = append(ctx.handlers, func(ctx *Context) {
			ctx.String(http.StatusNotFound, "404 NOT FOUND %s\n", ctx.Path)
		})
	}
	//调用所有的处理器
	ctx.Next()
}
