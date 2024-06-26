package main

func canJump(nums []int) bool {
	// start from the end
	// from the end, we know that we need the previous number to be 1
	// if it is not 1, move to the next number, and increase the number by 1
	// if we reach a number >= jumpsNeeded, reset jumpsNeeded
	// if we reach the end, and the number >= jumpsNeeded, return true
	if len(nums) <= 1 {
		return true
	}
	jumpsNeeded := 1
	k := len(nums) - 2

	for k >= 0 {
		if nums[k] >= jumpsNeeded {
			if k == 0 {
				return true
			}
			jumpsNeeded = 1
		} else {
			jumpsNeeded++
		}
		k--
	}

	return false
}
