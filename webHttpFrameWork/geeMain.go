/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
	"time"
)

func main() {
	engine := gee.New()
	engine.Use(gee.Logger())
	v1 := engine.Group("/v1")
	v1.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	v1.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	v2 := engine.Group("/v2")
	v2.Use(func4V2())
	v2.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v2.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"password": c.PostForm("password"),
		})
	})
	v3 := v1.Group("/v3")
	v3.GET("/hello/:name", func(c *gee.Context) {
		name := []string{c.Params["name"]}
		fmt.Printf(name[100])
	})

	fmt.Println("hello")
	engine.Run(":9090")
}

func func4V2() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Request.RequestURI, time.Since(t))
	}
}
