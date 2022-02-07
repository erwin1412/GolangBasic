package main

import (
	"Teknik_Dasar/function"
	"fmt"
)
func main(){
	fmt.Println("Per Function")
	result := function.Per(2,5,"a")
	fmt.Println(result)
	result2 := function.Per(3,5,"b")
	fmt.Println(result2)
	result3 := function.Per(12,5,"c")
	fmt.Println(result3)
	result4 := function.Per(33,5,"k")
	fmt.Println(result4)
	println("\n")
	println("\n")
	fmt.Println("Testing Array")
	arr:= [4]int{
		function.Per(2,5,"a"),
		function.Per(3,5,"b"),
		function.Per(12,5,"c"),
		function.Per(33,5,"k")}
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}
	println("\n")
	println("\n")
	fmt.Println("All Function")
	AllFunction := function.All(3,5)
	fmt.Println(AllFunction)
}

