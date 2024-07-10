package main

import "slices"

func largestPerimeter(nums []int) int64 {
	/**
	  A valid polygon, ordered by size of each side
	      [x, y, z]
	      x < y + z

	  Sort the numbers desc

	  keep track of the max possible at each index.

	  if a number is larger than the rest combined, remove it,
	      move to the next number
	  The result has to have a minimum of 3
	*/
	slices.Sort(nums)
	maxCombined := []int{}
	currentMax := 0

	for _, num := range nums {
		currentMax += num
		maxCombined = append(maxCombined, currentMax)
	}

	slices.Reverse(nums)
	slices.Reverse(maxCombined)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < maxCombined[i+1] {
			return int64(maxCombined[i])
		}
	}

	return -1
}