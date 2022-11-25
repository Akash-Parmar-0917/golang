package main

import (
	"fmt"
)

func main() {
	const (
		a int = iota + 2000
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}
