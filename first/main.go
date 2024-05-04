package main

import "fmt"

func main() {
	testNum := 3

	result := getCorrectWord(testNum)
	fmt.Println(result)
}
func getCorrectWord(num int) string {
	word := "компьютер"
	dict := make(map[int]string)
	dict[1] = ""
	dict[2] = "а"
	dict[3] = "а"
	dict[4] = "а"
	dict[5] = "ов"
	dict[6] = "ов"
	dict[0] = "ов"
	lastDigit := num % 10

	return word + dict[lastDigit]
}
