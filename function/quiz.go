package main

import "fmt"

func main() {
	fmt.Println(sum())
}

// cara 1
// func sum() int {
// 	scores := []int{10, 5, 8, 9, 7}
// 	total := 0
// 	for _,score := range scores{
// 		total = total + score
// 	}
// 	return total
// }

// cara 2
func sum() int {
	scores := []int{10, 5, 8, 9, 7}
	panjang := len(scores)
	total := 0
	for i := 0; i < panjang; i++ {
		total = total + scores[i]
	}
	return total
}