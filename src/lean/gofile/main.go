//主程序必须是main包
package main

//导入其他类库
import (
	"fmt"
)

//定义全局常量
const name = "lianghw"

//全局变量
var age int = 18

//定义类型（类）
type T struct {
	//属性
	job string
}

//init函数，执行main之前
func init() {

}

//主程序方法
func main() {
	var p T = T{"1"}
	fmt.Println(p)
}
