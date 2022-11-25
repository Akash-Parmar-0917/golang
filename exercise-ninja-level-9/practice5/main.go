package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


var wg sync.WaitGroup

func main(){
	var k int64
	wg.Add(100)
	for i:=0;i<100;i++{
		go func(){
			k=atomic.AddInt64(&k,1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(k)
}