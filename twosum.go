package main

func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int)

	for i, num := range nums {
		remainder := target - num

		if idx, found := intMap[remainder]; found {
			return []int{idx, i}
		}

		intMap[num] = i
	}

	return []int{}
}
