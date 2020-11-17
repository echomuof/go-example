/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package main

import (
	context "context"
	"go-example/grpc/proto/hello"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func init() {
	grpc.EnableTracing = true
}

const (
	port = ":50052"
)

type server struct {
}

func (s server) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Msg: "Hello, " + request.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("faild to listen :%v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, new(server))
	go startTrace()
	s.Serve(listen)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	log.Printf("trace listen on 50051")
}
