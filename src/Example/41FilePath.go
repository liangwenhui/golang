package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("etc","jk","ok.txt")
	fmt.Println(p)
	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.Base(p))
	fmt.Println(filepath.IsAbs("/a/b"))//true
	fmt.Println(filepath.IsAbs("a/b"))
	fmt.Println(filepath.Ext(p))
	fmt.Println(strings.TrimSuffix(p,".txt"))
	//获取basepath 于 目标路径的相对路径
	rel, _ := filepath.Rel("etc/", p)
	fmt.Println(rel)
	rel, _ = filepath.Rel("etc/r", p)
	fmt.Println(rel)
}
