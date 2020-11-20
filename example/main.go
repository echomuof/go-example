package main

import (
	"io"
	"log"
	"os"
)

type animal struct {
	name string
	age  int
}

func main() {
	srcFile, err := os.Open("./example/aa.txt")
	if err != nil {
		log.Println("打开源文件失败")
		panic(err)
	}
	defer srcFile.Close()

	targetFile, err := os.Create("./example/bb.txt")
	if err != nil {
		log.Println("创建输出文件失败")
		panic(err)
	}
	defer targetFile.Close()

	buf := make([]byte, 1024)
	for true {
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			log.Println("读取完毕")
			break
		}
		if err != nil {
			log.Println("读取失败")
			break
		}
		targetFile.Write(buf[:n])
	}
}
