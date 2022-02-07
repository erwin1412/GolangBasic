package main

import "fmt"

func main() {
	var kalimat = []string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[0] + " " + kalimat[3] + " " + kalimat[4] + " " + kalimat[5] + " " + kalimat[6])
	var sayuran = []string{}
	sayuran = append(sayuran,
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	)
	for i, sayur := range sayuran{
		fmt.Printf("elemen %d : %s\n" , i , sayur)
	}
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
}
