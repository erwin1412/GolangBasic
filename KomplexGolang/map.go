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

	for value := range mymap2{
	
		fmt.Println(value)	
	}

	for _,value := range mymap2{
	
		fmt.Println(value)	
	}

	for key,value := range mymap2{
		
		fmt.Println(key, " < " , value)	
	}
	
}