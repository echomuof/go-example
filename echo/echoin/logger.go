/**
 *
 * @author: echomuof
 * @created: 2020/11/18
 */
package echoin

import (
	"log"
	"time"
)

// 日志记录中间件
func Logger() HandlerFunc {
	return func(ctx *Context) {
		t := time.Now()
		ctx.Next()
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Request.RequestURI, time.Since(t))
	}
}
