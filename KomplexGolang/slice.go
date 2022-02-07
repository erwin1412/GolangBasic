package main

import "fmt"

//slice digunakan untuk menambah data , ketika kita buat sebuah array
//kita gk tau mw brp bnyk data yg ingin kita gunakan maka kita dapat menggunakan slice
func main() {
  var slice []string
	slice = append(slice, "Saya")
	slice = append(slice, "Baca" , "Sedang")
	slice = append(slice, "Buku")

	for index , slicing := range slice {
		fmt.Println("Index = " ,index , " Slice = ", slicing)
	}
}