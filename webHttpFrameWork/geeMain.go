/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	fmt.Printf("%p\n",engine)
	engine.GET("/hello", helloHandler)
	engine.POST("/abc", indexHandler)
	engine.Run(":8080")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "header[%q] = %q\n", k, v)
	}
}
