package main

import "fmt"

func main() {
	result1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("Results: %v", result1)

	result5 := longestPalindrome("babad")
	fmt.Printf("Longest Palindrom Substring is: %v", result5)
}
