package main

import (
	"fmt"
)

//for
func main() {

	for i := 1; i < 7; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

//while in golang
a:=1

for a<10{
	fmt.Println(a)
	a++
}


//foreach
title:="Erwin"

for index,letter := range title{
	if index%2==0 {
		fmt.Println(index,string(letter))
	}
	
}
}