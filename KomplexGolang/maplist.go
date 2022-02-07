package main

import "fmt"

func main() {
	var satuan = map[string]int{
		"Panjang": 7,
		"Lebar":   8,
		"tinggi":  6,
	}
	for key, val := range satuan {
		fmt.Println(key,"  \t:",val)
	}

	var chickens = []map[string]string{
		{"name":"Chicken blue","gender":"male"},
		{"name":"Chicken yellow","gender":"male"},
		{"name":"Chicken red","gender":"female"},
	}
	for _,chicken := range chickens {
		fmt.Println(chicken["name"],chicken["gender"])
	}
}