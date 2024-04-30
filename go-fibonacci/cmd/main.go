package main

import (
	"fmt"
)

func main() {
	n := 10
	Fibonacci(n)
}

func Fibonacci(n int) {
	x := 0
	y := 1
	z := x + y

	for i := 0; i <= n; i++ {
		if i == x {
			fmt.Println(i)
		} else if i == y {
			fmt.Println(y)
		} else {
			fmt.Println(z)
			x = y
			y = z
			z = x + y
		}
	}
}
