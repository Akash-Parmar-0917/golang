package main

import "fmt"


func foo() int {
	return 17
}


func bar() (int,string){
	return 9,"Akash"
}

func main(){
	a:=foo()
	b,c:=bar()
	fmt.Println(a,b,c)
}