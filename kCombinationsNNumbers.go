package main

func combine(n int, k int) [][]int {
	/**
	  there is less and less options as we reach the end
	  func backtrack(list)

	      if we reach a list of size k, store that list in the result

	      otherwise

	      loop through the range (1 -> n)
	          run current list through this function i.e [1, 1+currentIndex]
	      -----------------------
	      second iteration [1 ,2]
	      see if this is size k
	          if so, store it in results
	      else
	          loop through 2 -> n
	              run [1, 2, 2+currentIndex] through the next iteration
	      ----
	      add infinitum
	*/
	result := [][]int{}

	var dfsCombo func(cur []int)
	dfsCombo = func(cur []int) {
		if len(cur) == k {
			temp := make([]int, len(cur))
			copy(temp, cur)
			result = append(result, temp)
		} else {
			idx := -1
			if len(cur)-1 >= 0 {
				idx = cur[len(cur)-1]
			} else {
				idx = 0
			}
			idx++

			for idx <= n {
				dfsCombo(append(cur, idx))
				idx++
			}
		}
	}

	dfsCombo([]int{})
	return result
}