package main

import (
	"io/ioutil"
	"os"
)

func main() {
	var err error
	defer func() {
		if err!=nil{
			panic(err)
		}
	}()
	dir, _ := ioutil.ReadDir("D:\\subdir")
	if dir==nil{
		err = os.Mkdir("D:\\subdir", 0766)
	}
	//不能重复创建文件
	file, _ := ioutil.ReadFile("D:\\subdir\\a.txt")
	if(file==nil){
		ioutil.WriteFile("D:\\subdir\\a.txt",nil,0766)
	}
/*
os.Chdir 可cd到指定dir
通过.  .. 操作目录
   filepath.Walk() 递归访问目录，visit的每一个文件夹文件都会调用传入的func
	ioutil.TempFile 创建临时文件，会自动选择临时目录，以及自动给文件加上一串数字，用作并发
*/



}
