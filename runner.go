package main

import "fmt"

func main() {
	result1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("Results: %v \n", result1)

	result5 := longestPalindrome("babad")
	fmt.Printf("Longest Palindrom Substring is: %v", result5)

	result11 := ContainerWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Printf("Container with most water contains %v area", result11)

	result12 := IntegerToRoman(1994)
	fmt.Printf("%v is %v", 1994, result12)
}
