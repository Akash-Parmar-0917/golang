package main

import (
	"fmt"
)

type myAge int

var x myAge

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
