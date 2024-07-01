package main

func search(nums []int, target int) int {
	/**
	  determine the index of the rotation
	  if the end is not more than the beginning, there is no rotation

	  determine the lhs and the rhs after determining the pivot

	  determinePivot is where the currentNum > nextNum

	*/

	result := -1

	var hasPivot = func(first, last int) bool {
		return first > last
	}

	var binarySearch func(left, right int)
	binarySearch = func(left, right int) {
		newPivot := (right - left) / 2
		if nums[newPivot] == target {
			result = newPivot
			return
		}
		if nums[left] == target {
			result = left
			return
		}
		if nums[right] == target {
			result = right
			return
		}

		if nums[newPivot] > target {
			// lhs search
			binarySearch(left, newPivot)
		} else {
			// rhs search
			binarySearch(newPivot, right)
		}
	}

	pivot := hasPivot(nums[0], nums[len(nums)-1])
	pivotIndex := len(nums) / 2

	if pivot {
		for i := 0; i < len(nums); i++ {
			if nums[i] > nums[i+1] {
				pivotIndex = i
				break
			}
		}
	}

	//determine first loop of whether to select rhs or lhs
	// pick the side where the firstNum < target < lastNum

	if nums[0] < target && target < nums[pivotIndex] {
		binarySearch(0, pivotIndex)
	} else if nums[pivotIndex] <= target && target <= nums[len(nums)-1] {
		binarySearch(pivotIndex, len(nums)-1)
	}

	return result

}
