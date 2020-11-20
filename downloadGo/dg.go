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
	preStr := "https://lzdliving.alicdn.com/live_hp/840184c1-4388-4274-8f1a-363d40f3fb9f/"
	postStr := ".ts?auth_key=1608433228-0-0-9086b345162058b2a519a34ba920d1e2"
	dir := "/Users/wdr/Downloads/"
	file, _ := os.OpenFile(dir+"gooo.mp4", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	w := bufio.NewWriter(file)
	defer file.Close()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	for i := 1; i < 350; i++ {
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
