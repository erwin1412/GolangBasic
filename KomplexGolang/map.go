package main

import "fmt"

func main() {

	mymap := map[string]string{
		"Banana" : "Pisang",
		"Apple" : "Apel",
	}
	fmt.Println(mymap)

	mymap2 := map[int]int{
		3 : 1,
		4 : 2,
	}
	for index,maps := range mymap2{
		
		fmt.Println(index, maps)	
	}

	for maps := range mymap2{
	
		fmt.Println(maps)	
	}

	for _,maps := range mymap2{
	
		fmt.Println(maps)	
	}

	for index,maps := range mymap2{
		
		fmt.Println(index, " < " , maps)	
	}
	
}