package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("Akash Parmar")
	wg.Done()
}

func bar() {
	fmt.Println("Hiren Parmar")
	wg.Done()
}

func main() {
	wg.Add(2)

	go bar()
	go foo()
	wg.Wait()
}
