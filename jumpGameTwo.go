package main

func jumpTwo(nums []int) int {
	/**
	  current destination should be the first index (0)
	  highest jump recorded is 0

	  determine the highest jump possible
	  loop through the array (except the end, because the end doesn't count)
	      determine the farthest you can go given the current value
	      determine if this is the farthest you can reach
	          given the highest jump recorded so far

	      if the index is at the destination
	          increase number of jumps needed
	          set destination to be the highest jump so far
	*/

	destination, highestJumpSoFar, jumps := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		highestJumpHere := i + nums[i]
		if highestJumpHere > highestJumpSoFar {
			highestJumpSoFar = highestJumpHere
		}

		if i == destination {
			jumps++
			destination = highestJumpSoFar
		}
	}

	return jumps
}