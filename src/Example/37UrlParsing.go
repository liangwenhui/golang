package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	u := "http://admin:12345@localhost:8080/eureka?opt=add#k"
	uurl, err := url.Parse(u)
	defer func() {
		if err!=nil{
			panic(err)
		}
	}()
	fmt.Printf("%p \n",uurl)
	fmt.Println(uurl)
	fmt.Println(*uurl)
	fmt.Println("-----------------------")
	fmt.Println(uurl.Scheme)
	fmt.Println(uurl.User)
	fmt.Println(uurl.User.Username())
	fmt.Println(uurl.User.Password())
	fmt.Println(uurl.Host)
	fmt.Println(net.SplitHostPort(uurl.Host))
	fmt.Println(uurl.Fragment)
	fmt.Println(uurl.RawQuery)
	//params map
	query, err := url.ParseQuery(uurl.RawQuery)
	fmt.Println(query)
}
