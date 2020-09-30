package main

import (
	"fmt"
)

func main() {
	var car Car = new(BtCar)
	car.dirve()
	car = new(FtCar)
	car.dirve()
	car = new(Dog)
	car.dirve()
}

type Car interface {
	dirve()
}

type BtCar struct{}
type FtCar struct{}

func (v BtCar) dirve() {
	fmt.Println("开本田")
}
func (v FtCar) dirve() {
	fmt.Println("开丰田")
}

type Dog struct {
}
