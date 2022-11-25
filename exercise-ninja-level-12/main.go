package main

import (
	"fmt"
	"./dog"
)



type canin struct{
	name string
	age int
}

func main(){
	a:=canin{
		name: "faido",
		age: dog.Years(10),
	}
	fmt.Print(a)
}