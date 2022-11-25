package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	go gen(c)
	for v := range c {
		fmt.Println(v)
	}
}

func gen(c chan int) {
	wg.Add(10)
	for j := 0; j < 10; j++ {
		go func(k int) {
			for i := 0; i < 10; i++ {
				//fmt.Println(j, i)
				c <- k*10 + i
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
	close(c)
}
