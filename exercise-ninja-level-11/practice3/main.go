package main

import (
	"fmt"
)

type customErr struct {
	err string
}

func (cu customErr) Error() string {
	return "calling Error function: " + cu.err
}

func foo(er error) {
	fmt.Println(er)
}

func main() {

	c1 := customErr{"custom error 1"}
	foo(c1)
}
