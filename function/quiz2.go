package main

import "fmt"
func main(){
	tambah,kurang,bagi,kali := hitung(10,2)
	fmt.Println(tambah)
	fmt.Println(kurang)
	fmt.Println(bagi)
	fmt.Println(kali)
}
func hitung(number1,number2 int)(tambah,kali,bagi,kurang float32){
	tambah = float32(number1+number2)
	kurang = float32(number1-number2)
	bagi = float32(number1/number2)
	kali = float32(number1*number2)
	return
}