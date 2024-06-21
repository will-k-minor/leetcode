package main

import "fmt"

func main() {
	result1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("Results: %v \n", result1)

	result5 := longestPalindrome("babad")
	fmt.Printf("Longest Palindrom Substring is: %v \n", result5)

	const problem6String = "PAYPALISHIRING";
	result6 := zigzag(problem6String, 3);
	fmt.Printf("ZigZag with %v number of rows: %v \n", problem6String, result6);
}
