package main

func luckyNumbers(matrix [][]int) []int {
	/**
	  loop through each row
	  rowMin[x] means the min for xth row
	  columnMin[y] means the min for yth row

	  loop through each row (x)
	      loop across the columns (y)
	          curr = mx[x][y]
	          if columnMin[y] is empty || columnMin[y] > curr
	              columnMin[y] = curr;
	          if rowMin[x] is empty || rowMin[x] > curr
	              rowMin[x] = curr;
	  store longer array as a map
	  loop through the smaller array
	      if there is a match in the map,
	          store it in the result array
	*/
	var rowMin, colMax []int

	for x, row := range matrix {
		for y, curr := range row {
			if len(colMax) <= y {
				colMax = append(colMax, curr)
			} else if colMax[y] < curr {
				colMax[y] = curr
			}

			if len(rowMin) <= x {
				rowMin = append(rowMin, curr)
			} else if rowMin[x] > curr {
				rowMin[x] = curr
			}
		}
	}

	numMap := make(map[int]bool)
	colIsLarger := len(rowMin) < len(colMax)
	result := []int{}

	if colIsLarger == true {
		for _, curr := range colMax {
			numMap[curr] = true
		}
		for _, curr := range rowMin {
			_, ok := numMap[curr]
			if ok {
				result = append(result, curr)
			}
		}
	} else {
		for _, curr := range rowMin {
			numMap[curr] = true
		}
		for _, curr := range colMax {
			_, ok := numMap[curr]
			if ok {
				result = append(result, curr)
			}
		}
	}

	return result

}