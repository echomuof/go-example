/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
}

func (mh myHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "URL.path = %q\n", request.URL.Path)
		}(writer, request)
	case "/hello":
		func(writer http.ResponseWriter, request *http.Request) {
			for k, v := range request.Header {
				fmt.Fprintf(writer, "header[%q]=[%q]\n", k, v)
			}
		}(writer, request)
	default:
		fmt.Fprintf(writer, "404 NOT FOUNT %q", request.URL)
	}

}

func main() {
	person := &Person{
		Name: "judy",
		Age:  20,
	}
	person.UpdateName("mike")
	fmt.Println(person)
	person.UpdateAge(30)
	fmt.Println(person)

}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "header[%q] = %q\n", k, v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) UpdateName(name string) error {
	p.Name = name
	return nil
}

func (p *Person) UpdateAge(age int) error {
	p.Age = age
	return nil
}
