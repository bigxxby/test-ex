package main

import "fmt"

func main() {
	foo(5)
}

func foo(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println("")
	}
}
