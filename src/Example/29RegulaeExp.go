package main

import (
	"fmt"
	"regexp"
)

//正则
func main() {
	matchString, _ := regexp.MatchString("p([a-z]*)ch", "patch")
	fmt.Println(matchString)
	compile, _ := regexp.Compile("p([a-z]*)ch")
	fmt.Println(compile.MatchString("pathc"))
	fmt.Println(compile.FindString("patch pahyc passch"))
	fmt.Println(compile.FindAllString("patch pahyc passch",-1))
	fmt.Println(compile.FindStringIndex("paqch pahyc passch"))
	fmt.Println(compile.FindStringSubmatch("paqch pahyc passch"))
	fmt.Println(compile.FindStringSubmatchIndex("patch pahyc passch"))

}
