package main

func rearrangeArray(nums []int) []int {
	/**
	  keep two arrays, as these will help preserve order
	      pos
	      neg

	  as we loop through the arrays, we alternate between them
	      pos[0]
	      neg[0]
	      post[1]
	      neg[1]
	      ....
	      pos[n-1]
	      neg[n-1]
	*/
	pos := []int{}
	neg := []int{}
	for _, num := range nums {
		if num >= 0 {
			pos = append(pos, num)
		} else {
			neg = append(neg, num)
		}
	}

	res := []int{}
	posIdx := 0
	negIdx := 0
	dir := 1
	for posIdx+negIdx < len(nums) {
		// we've hit the end of pos
		if posIdx == len(pos) {
			res = append(res, neg[negIdx])
			negIdx++
		} else if negIdx == len(neg) {
			res = append(res, pos[posIdx])
			posIdx++
		} else if dir > 0 {
			res = append(res, pos[posIdx])
			posIdx++
		} else {
			res = append(res, neg[negIdx])
			negIdx++
		}

		dir = -dir
	}

	return res
}