package main

/*
 			go							java
IO			bufio/io					io/nio
字符串操作	strings						string
容器			container(heap/list/ring)	Collection
锁			sync						juc
时间			time						time/date
算数			math						Math
unsafe		unsafe						unsafe
*/

//常用基础类型
func main() {
	//1.boolean
	var _ bool = true
	//2.有符号整形
	var _ int = 4      //4或8字节 32位 64位
	var _ int8 = 0xf   //1字节 8位
	var _ int16 = 0xff //2字节 16位
	var _ int32 = 0xffff
	var _ int64 = 0xffffffff
	//3 无符号整形
	var _ uint = 4 //4或8字节
	var _ uint8
	var _ uint16
	var _ uint32
	var _ uint64
	//4 浮点型
	var _ float32
	var _ float64
	//5 字符串
	var _ string
	//6 字节
	var _ byte //等于uint8
	var _ rune //等于int32
	//7 派生数据类型
	/*
		指针	pointer
		数组	array
		结构体 struct
		管道 channel
		函数 func
		切片 slice
		接口 interface
		哈希 map
	*/

}
