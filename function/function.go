package main

import "fmt"

func main(){
	// varibel luas1 dan keliling1 tidak wajib mengikuti variabel output 
	// pada function , yg mana disana adalah luas dan keliling
	//maka di func main bisa di atur jadi luas1 dan keliling1 atau apa aja serah
	luas1,keliling1 := hitung(10,2)
	fmt.Println(luas1)
	fmt.Println(keliling1)

	//kalau cuman butuh luas maka keliling ganti dengan _ , begitu juga sebaliknya
	luas2,_:=hitung(10,2)
	fmt.Println(luas2)
}

//cara 1
// func namafunction(input variabel 1 type data , variabel 2 type data) (output typedata 1 , typedata 2 )
// func hitung(panjang int,lebar int)(int,int){
// 	luas:= panjang * lebar
// 	keliling := 2 * (panjang+lebar)
// 	return	luas,keliling
// }

// lebih enak cara 1 sih , cuman cara 2 pun enak juga jadi bebas pilih saja toh sama aja

//cara 2
// func namafunction(input variabel 1 type data , variabel 2 type data) (output variabel 1 typedata 1 , variabel 2 typedata 2 )
func hitung(panjang int,lebar int)(luas int, keliling int){
	luas= panjang * lebar
	keliling = 2 * (panjang+lebar)
	return
}