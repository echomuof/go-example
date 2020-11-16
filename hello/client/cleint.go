/**
 *
 * @author: echomuof
 * @created: 2020/11/17
 */
package main

import (
	"context"
	"go-example/proto/hello"
	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:50052"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect:%v", err)
	}
	defer conn.Close()

	c := hello.NewHelloClient(conn)
	res, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatal("could not say hello :%v", err)
	}
	log.Printf("call rpc res:%s", res.Msg)
}
