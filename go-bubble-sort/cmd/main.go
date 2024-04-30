package main

import (
	"fmt"
)

func BubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}

	return data
}

func PrintData(data []int) {
	for _, v := range data {
		fmt.Printf("%v ", v)
	}

	fmt.Println()
}

func main() {
	data := []int{89, 11, 21, 6, 53, 35}
	sortedData := BubbleSort(data)
	PrintData(sortedData)
}
