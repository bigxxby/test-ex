package main

import "fmt"

func main() {
	first := 11
	second := 20
	res := returnPrimeNumbers(first, second)
	fmt.Println(res)
}
func returnPrimeNumbers(first, second int) []int {
	var res []int
	for number := first; number < second; number++ {
		current := number
		isPrime := true
		for i := 2; i < current; i++ {
			if current%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			res = append(res, current)
		}
	}
	return res

}
