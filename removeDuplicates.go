package main

func removeDuplicates(nums []int) int {
	// create an empty Map where it is map[int]int
	// if the value does not exist in the map yet, set nums[current] = val

	myMap := map[int]int{}
	currentIndex := 0

	for _, v := range nums {
		if myMap[v] == 0 {
			nums[currentIndex] = v
			myMap[v] = 1
			currentIndex++
		}
	}

	return currentIndex
}
