package main

import (
	"fmt"
)

func main() {
	car := new(BtCar)
	car.dirve()
	fcar := new(FtCar)
	fcar.dirve()
	fcar.fix()

}

type Tool interface {
	fix()
}

type Car interface {
	Tool
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
func (v FtCar) fix() {
	fmt.Println("修丰田")
}

type Dog struct {
}
