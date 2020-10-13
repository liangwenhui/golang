package main

import (
	"fmt"
	"sort"
)

type stu struct {
	age int
	name string
}
type stulist []stu

func (s stulist)Len()int{
	return len(s)
}
func(s stulist)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}
//func(s stulist) Less(i,j int)bool{
//	return (s[i].age)<(s[j].age)
//}
func(s stulist) Less(i,j int)bool{
	return len(s[i].name)<len(s[j].name)
}

func main() {
	var arr stulist = make(stulist,5)
	arr[0] = stu{11,"a"}
	arr[1] = stu{6,"ccc"}
	arr[2] = stu{17,"bb"}
	arr[3] = stu{10,"fffff"}
	arr[4] = stu{9,"dddd"}
	fmt.Println(arr)
	sort.Sort(arr)
	fmt.Println(arr)

}
