package main

import "fmt"


func foo() func(){
	return func ()  {
		fmt.Println("Akash Parmar")
	}
}
func main(){
	Mam:=foo()
	Mam()
}