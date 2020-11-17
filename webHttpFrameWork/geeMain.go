/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello~~~</h1>")
	})

	engine.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s", ctx.Query("name"))
	})

	engine.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	engine.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	engine.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	engine.Run(":9090")
}
