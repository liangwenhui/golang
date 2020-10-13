package main

import (
	"testing"
)

func Test(t *testing.T){
	i :=0
	t.Error("1 err",i)
	i++
	t.Error("2 err",i)

}
