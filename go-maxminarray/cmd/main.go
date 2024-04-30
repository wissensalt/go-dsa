package main

import (
	"fmt"
)

func main() {
	arr := []int{11, 28, 95, 12, 90, 221, 82, 37, 18, 65, 89, 47, 62, 52, 69, 78, 83}
	maxVal := FindMaxElementInArray(arr)
	minVal := FindMinElementInArray(arr)

	fmt.Println("Max: ", maxVal)
	fmt.Println("min: ", minVal)
}

func FindMinElementInArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < result {
			result = arr[i]
		}
	}

	return result
}

func FindMaxElementInArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > result {
			result = arr[i]
		}
	}

	return result
}
