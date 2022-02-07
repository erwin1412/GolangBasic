package main

import "fmt"

func main(){

	//cara 1
	var nilai int
	fmt.Println("INPUT NILAI")
	fmt.Scanln(&nilai)

	var grade string
	if nilai >= 80{
		grade ="A"
	}else if nilai <80 && nilai >69 {
		grade ="B"
	}else if nilai <70 && nilai >59 {
		grade ="C"
	}else{
		grade ="D"
	}
	fmt.Println(grade)

	//cara 2
	// var nilai = 80
	// if nilai >= 80{
	// 	fmt.Println("Grade A")
	// }else if nilai <80 && nilai >69 {
	// 	fmt.Println("Grade B")
	// }else if nilai <70 && nilai >59 {
	// 	fmt.Println("Grade C")
	// }
}