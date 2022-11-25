package main

import "fmt"

func print(name string) string {
	return "You are brilliant " + name
}

func detail(p func(string) string, name string) {
	val := p(name)
	fmt.Println(val)
}

func main() {
	detail(print, "Akash Parmar")
}
