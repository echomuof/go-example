/**
 *
 * @author: echomuof
 * @created: 2020/11/21
 */
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("listener tcp port error\n")
	}
	defer listener.Close()

	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("conn error %v", err)
			continue
		}
		defer conn.Close()
		go process(conn)
	}
}

func process(conn net.Conn) {
	for true {
		buf := make([]byte, 128)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Errorf("read from connection error : %v", err)
			break
		}
		msg := string(buf[:n])
		if msg == "Q" {
			break
		}
		fmt.Printf("revicve from client: %s\n", msg)
	}
}
