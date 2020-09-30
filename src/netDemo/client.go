package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8088")
	fmt.Println(conn)
	log.Println("conn success")
	defer conn.Close()
	log.Println(os.Stdout)
	io.Copy(os.Stdout, conn)
}
