package main

import "fmt"


func foo(x ...int) int {
	var sum int 
	
	for _,i := range x {
		sum+=i
	}
	return sum
}

func bar(x []int) int{
	var sum int 
	
	for _,i := range x {
		sum+=i
	}
	return sum
}

func main(){
	x:=[]int{1,2,3,4,5,56}
	sum1:=foo(x...)
	fmt.Println(sum1)
	sum1=bar(x)
	fmt.Println(sum1)
}