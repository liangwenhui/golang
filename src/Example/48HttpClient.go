package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
/**
resp, err := http.Post("http://www.baidu.com",
        "application/json",
        strings.NewReader("name=cjb"))
 */
func main() {
	response,err := http.Get("http://www.baidu.com")
	defer func() {
		if(err!=nil){
			panic(err)
		}
	}()
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	all, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(all))

}
