package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func (p person) speak (){
	fmt.Println("my name is:",p.first,p.last,"and my age is:",p.age)
}
func main(){
	p:=person{"Akash","Parmar",21}
	p.speak()
}