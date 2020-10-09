package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	handle := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	buf := make([]byte, 2048)
	file, err1 := os.Open("D:\\专业版激活码.txt")
	defer handle(err1)
	defer file.Close() //记得关闭流
	reader := bufio.NewReader(file)
	//如果没有文件则创建
	tag, err1 := os.OpenFile("D:\\专业版激活码1.txt", os.O_CREATE|os.O_RDWR, 0766)
	writer := bufio.NewWriter(tag) //标准输出

	defer writer.Flush()
	for {
		data, _ := reader.Read(buf)
		if data == 0 {
			break
		}
		writer.Write(buf[0:data])
	}
}
