package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y:=x[:5]
	z:=x[5:]
	w:=x[2:7]
	v:=x[1:6]
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
	fmt.Println(y,z,w,v)
}
