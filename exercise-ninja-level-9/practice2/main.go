package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Akash",
		last:  "Parmar",
	}
	saySomething(&p)
}
