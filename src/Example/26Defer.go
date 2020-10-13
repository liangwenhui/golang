package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("D:\\ok.txt")
	defer closeFile(file)
	writeFile(file,"ok verry match")

}

func createFile(s string) *os.File{
	fmt.Println("createFile")
	file, err := os.OpenFile(s, os.O_CREATE|os.O_RDWR, 0766)
	if err!=nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File,s string){
	fmt.Println("writeFile")
	fmt.Fprintln(file,s)
}
func closeFile(file *os.File){
	fmt.Println("closeFile")
	file.Close()

}
