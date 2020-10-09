package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	buf := make([]byte, 2048)
	file, err1 := os.Open("D:\\专业版激活码.txt")
	defer errHandle(err1)
	defer file.Close() //记得关闭流
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout) //标准输出
	defer writer.Flush()
	for {
		data, _ := reader.Read(buf)
		if data == 0 {
			break
		}
		writer.Write(buf[0:data])
	}
}

//简单的异常捕获，抛出异常到控制台，中断程序
func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
