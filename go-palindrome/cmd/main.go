package main

import (
	"fmt"
	"strings"
)

// IsPalindrome return true if value is Palindrome
func IsPalindrome(value string) bool {
	sanitizedValue := sanitizeString(value)
	counter := 1
	result := false
	for {
		leftScan := sanitizedValue[counter-1]
		rightScan := sanitizedValue[len(sanitizedValue)-counter]
		if leftScan == rightScan {
			result = true
		}

		if counter == len(sanitizedValue) {
			return result
		}

		counter++
	}
}

func sanitizeString(value string) string {
	if len(value) == 0 {
		return ""
	}

	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, " ", "")

	return value
}

func main() {
	fmt.Println("ABCDEF: ", IsPalindrome("ABCDEF"))
	fmt.Println("ABCDEFFEDCBA: ", IsPalindrome("ABCDEFFEDCBA"))
	fmt.Println("RACE CAR", IsPalindrome("RACE CAR"))
}
