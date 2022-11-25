package main

import (
	"fmt"
	"math"
)



type circle struct{
	radius int
}

type square struct{
	length int

}

func (c circle) area () float64{
	pi:=math.Pi
	return pi * float64(c.radius*c.radius)
}

func (s square) area() float64{
	return float64(s.length*s.length)
}

type shape interface{
	area() float64
}

func info(s shape){
	
	// switch s.(type){
	// case square:
	// 	fmt.Println(s.(square).area())
	// case circle:
	// 	fmt.Println(s.(circle).area())
	// }
	fmt.Println(s.area())
}
func main(){
	cir:=circle{10}
	info(cir)
	sq:=square{5}
	info(sq)
}