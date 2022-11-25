package main

import "fmt"

func main(){
	s:= struct {
		firstName string
		lastName string
	}{
		firstName: "Dr. Akash",
		lastName: "Parmar",
	}

	fmt.Printf("%v %v",s.firstName,s.lastName)
}