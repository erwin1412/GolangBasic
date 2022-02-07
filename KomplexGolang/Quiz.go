package main

import "fmt"
func main() {
	//hitung rata2
	//hitung goodscore >= 90
	scores := [8]int{100,80,75,92,70,93,88,67}
	var nilai float32
	nilai = 0

	for _,score:=range scores{
		fmt.Print("nilai sekarang : ",nilai ," + ", score , " = ")
		nilai=nilai+float32(score)
		fmt.Println(nilai)
	}
	fmt.Println("\nTotal Nilai : ",nilai)
	fmt.Println("Nilai Rata2 : ",(nilai/float32((len(scores)))))

	goodscore := []int{}
	for _,scorez := range scores{
		if scorez >= 90 {
			goodscore = append(goodscore,scorez)
		}
	}
	fmt.Println(goodscore)
}