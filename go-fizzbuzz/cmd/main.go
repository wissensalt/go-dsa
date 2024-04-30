package main

import (
	"fmt"
)

func main() {
	n := 100
	FizzBuzz(n)
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
