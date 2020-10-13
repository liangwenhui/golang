package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name
	Id int32
	Name string
	Origin []string
}

func main() {
	coffee := &Plant{Id: 12,Name: "coffee"}
	coffee.Origin = []string{"add","insert","write"}
	indent, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(xml.Header,string(indent))

	//s := "  <Plant>   <Id>13</Id>  <Name>java</Name>   <Origin>add</Origin>   <Origin>insert</Origin> <Origin>write</Origin> </Plant>"
	//x,_:=xml.Marshal(s)
	var java Plant
	err := xml.Unmarshal(indent, &java)
	if err!=nil{
		panic(err)
	}
	fmt.Println(java)
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	indent, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(indent))
}
