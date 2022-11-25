package main

import (
	"fmt"
	"runtime"
	"sync"
)


var wg sync.WaitGroup

func main(){
	k:=0
	wg.Add(100)
	for i:=0;i<100;i++{
		go func(){
			t:=k
			runtime.Gosched()
			t++
			k=t
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(k)
}