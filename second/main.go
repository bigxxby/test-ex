package main

import (
	"fmt"
)

func main() {
	numbers := []int{42, 12, 18}
	result := findCommonDivisors(numbers)
	fmt.Println(result)
}

func findCommonDivisors(numbers []int) []int {
	divisorFreq := make(map[int]int)

	for _, num := range numbers {
		divisors := findDivisors(num)
		for _, divisor := range divisors {
			divisorFreq[divisor]++
		}
	}

	var commonDivisors []int
	for divisor, freq := range divisorFreq {
		if freq == len(numbers) {
			commonDivisors = append(commonDivisors, divisor)
		}
	}

	return commonDivisors
}

func findDivisors(num int) []int {
	var divisors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
