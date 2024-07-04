package main

func search(nums []int, target int) int {
	/**
	  determine the index of the rotation
	  if the end is not more than the beginning, there is no rotation

	  determine the lhs and the rhs after determining the pivot

	  determinePivot is where the currentNum > nextNum

	*/
	// the "pivot" is the true beginning ofthe array

	if len(nums) <= 2 {
		for idx, val := range nums {
			if val == target {
				return idx
			}
		}

		return -1
	}

	findPivot := func() int {
		left, right := 0, len(nums)-1

		if nums[left] < nums[right] {
			return 0 // there is no pivot
		}

		// binary search
		for left <= right {
			mid := left + (right-left)/2

			if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
				return mid + 1 // we found the beginning
			}

			if nums[left] <= nums[mid] {
				left = mid + 1 // shifts our window to the right
			} else {
				right = mid - 1 // shifts our window to the left
			}
		}
		return 0
	}

	binarySearch := func(left, right int) int {
		for left <= right {
			mid := left + (right-left)/2

			if nums[mid] == target {
				return mid
			}

			if nums[mid] < target {
				left = mid + 1 // search rhs
			} else {
				right = mid - 1 // search lhs
			}
		}
		return -1
	}

	pivotIndex := findPivot()

	if pivotIndex == 0 {
		return binarySearch(0, len(nums)-1)
	}

	if target >= nums[0] && target <= nums[pivotIndex-1] {
		// pick lhs
		return binarySearch(0, pivotIndex-1)
	} else {
		// pick rhs
		return binarySearch(pivotIndex, len(nums)-1)
	}
}
