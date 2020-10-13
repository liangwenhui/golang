package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello",hello)
	http.HandleFunc("/header",headers)
	http.ListenAndServe(":9090",nil)


}

func headers(repr http.ResponseWriter,req *http.Request){
	for name,headers := range req.Header{
		for _,h := range headers{
			fmt.Fprintf(repr,"%v : %v \n",name,h)
		}
	}
}


func hello(repr http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	req.ParseForm()
	var name = req.Form.Get("name")
	fmt.Println(name)
	fmt.Fprintf(repr,"hello!")
}