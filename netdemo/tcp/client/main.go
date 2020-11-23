/**
 *
 * @author: echomuof
 * @created: 2020/11/21
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Printf("connection to server error : %v", err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for true {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Errorf("read user input error : %v", err)
			break
		}

		trimInput := strings.TrimSpace(input)
		if trimInput == "Q" {
			break
		}

		_, err = conn.Write([]byte(trimInput))
		if err != nil {
			fmt.Errorf("send msg error : %v", err)
			break
		}
	}
}
