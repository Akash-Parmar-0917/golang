package main

import "fmt"


func main(){
	x:=[]string{"James", "Bond", "Shaken, not stirred"}
	y:=[]string{"Miss", "Moneypenny", "Helloooooo, James."}
	z:=[][]string{x,y}

	for _,v:=range z{
		for _,w:=range v{
			fmt.Printf("%v ",w)
		}
	}
}