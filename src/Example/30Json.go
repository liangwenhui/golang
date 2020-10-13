package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	b, _ := json.Marshal(true)
	fmt.Println(string(b))
	i,_ := json.Marshal(12)
	fmt.Println(string(i))
	arr,_ := json.Marshal([]int{1,2,3})
	fmt.Println(string(arr))
	m,_ := json.Marshal(map[string]string{"a":"ok","c":"fald"})
	fmt.Println(string(m))
	//json字符串转构造体
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := map[string]interface{}{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res["fruits"])
	file,err:=os.OpenFile("D:\\ok2.txt",os.O_RDWR|os.O_CREATE,0766)

	writer := bufio.NewWriter(file)
	enc := json.NewEncoder(writer)
	d := map[string]int{"apple": 5, "lettuce": 7}
	err = enc.Encode(d)
	writer.Flush()
	defer file.Close()
	defer func() {
		if err!=nil{
			panic(err)
		}
	}()

}
