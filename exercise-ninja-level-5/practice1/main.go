package main

import "fmt"

type person struct{
	firstName string
	lastName string
	iceCream string
}

func main(){
	p1:= person{"Akash","Parmar","Vanilla"}
	p2:=person{lastName:"Hiren",firstName: "Parmar",iceCream: "Strawberry"}
	fmt.Println(p1,p2)
}