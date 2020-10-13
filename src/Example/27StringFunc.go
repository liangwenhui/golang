package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println
	p("Contains:",s.Contains("abc","bc"))
	p("Count",s.Count("abcabcabca","ab"))
	p("HasPrefix",s.HasPrefix("foobar","foo"))
	p("HasSuffix",s.HasSuffix("foobar","bar"))
	p("Index",s.Index("asrsgas","s"))
	p("Join",s.Join([]string{"a","b"},"-"))
	p("Repeat",s.Repeat("*",5))
	p("Replace",s.Replace("ojbkojbk","jb","**",1))
	p("Replace",s.ReplaceAll("ojbkojbk","jb","**"))
	p("Split",s.Split("a,b,c",","))
	p("Lower",s.ToLower("ABCD"))
	p("Upper",s.ToUpper("abcd"))
	p("Len",len("abcdefg"))
	p("Char","good"[2])



}
