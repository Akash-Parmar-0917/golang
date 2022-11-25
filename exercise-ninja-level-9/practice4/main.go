package main

import (
	"fmt"
	"runtime"
	"sync"
)


var wg sync.WaitGroup

func main(){
	k:=0
	var mu sync.Mutex
	wg.Add(100)
	for i:=0;i<100;i++{
		go func(){
			mu.Lock()
			t:=k
			runtime.Gosched()
			t++
			k=t
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(k)
}