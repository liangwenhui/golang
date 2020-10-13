package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	address := flag.NewFlagSet("address",flag.ExitOnError)
	host := address.String("host", "localhost", "ip")
	port :=address.Int("port",8080,"port")
	net := flag.NewFlagSet("net",flag.ExitOnError)
	ssl:=net.Bool("ssl",false,"https?")
	opt := net.String("opt", "get", "get put post delete")
	fmt.Println(len(os.Args))
	fmt.Println(os.Args)
	if len(os.Args)<2{
		fmt.Println("expected address or ssl ")
		os.Exit(1)
	}
	for set := range os.Args{
		switch string(set) {
		case "address":

		case "net":

		}
	}

	fmt.Println(*host,*port,*ssl,*opt)
	
}
