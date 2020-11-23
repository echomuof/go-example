package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	preStr := "https://lzdliving.alicdn.com/live_hp/71a840e3-ca35-41ab-b89c-28f1eb29db1e/"
	postStr := ".ts?auth_key=1608562391-0-0-651d64fdd203db7e41c1edc8d7a0529a"
	dir := "/Users/darius/Documents/video/GOOO/"
	file, _ := os.OpenFile(dir+"gooo2.mp4", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	w := bufio.NewWriter(file)
	defer file.Close()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	for i := 1; i <= 420; i++ {
		time.Sleep(20 * time.Millisecond)
		func(i int) {
			url := fmt.Sprintf(preStr+"%d"+postStr, i)
			resp, err := client.Get(url)
			if err != nil {
				log.Printf("%d 请求失败\n", i)
				log.Println(err)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("读取失败，%s", err.Error())
			}
			w.Write(body)
			fmt.Printf("%d 成功\n", i)
		}(i)
	}
	w.Flush()
}
