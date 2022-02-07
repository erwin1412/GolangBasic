package function

func Per(number int, numberTwo int, flag string) int {
	if flag == "a" {
		return number + numberTwo
	} else if flag == "b" {
		return number - numberTwo
	} else if flag == "c" {
		return number * numberTwo
	}
	return number / numberTwo

}

func All(number int, numberTwo int) int {
	println(number + numberTwo)
	println(number - numberTwo)
	println(number * numberTwo)
	return number / numberTwo
}

// func Min(angka1 int, angka2 int) int {
// 	return angka1 - angka2
// }