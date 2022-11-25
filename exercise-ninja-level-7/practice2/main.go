package main

import "fmt"


type person struct{
	first string
	last string
	age int
}

func changeMe(p *person){
	//fmt.Println(p.first,(*p).first)
	p.first="Hiren"
}

func main(){
	p:=person{
		first: "Akash",
		last: "Parmar",
		age:21,
	}
	changeMe(&p)
	fmt.Println(p)
}