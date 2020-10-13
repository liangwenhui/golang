package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host","localhost","ip address")
	port := flag.Int("port",8080,"ip port")
	opt := flag.String("opt","get","[get post put delete]")
	optset := flag.NewFlagSet("opt", flag.ExitOnError)
	optset.Bool("ssl",false,"use https?")
	optset.String("opt","get","[get post put delete]")
	var uname string
	flag.StringVar(&uname,"u","","login user name")
	flag.Parse()
	fmt.Println(*host)
	fmt.Println(*port)
	fmt.Println(*opt)
	fmt.Println(uname)
	fmt.Println(flag.Args())

}
