package main

func ContainerWithMostWater(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	k := len(height) - 1
	i := 0
	result := 0

	min := func(num1, num2 int) int {
		if num1 > num2 {
			return num2
		}

		return num1
	}

	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}

		return num2
	}

	// start from two ends
	// move on from the lower one (i.e. if height[i] )
	for i < k {
		width := k - i
		minHeight := min(height[i], height[k])

		result = max(result, minHeight*width)
		if height[i] < height[k] {
			i++
		} else if height[i] > height[k] {
			k--
		} else {
			i++
		}
	}

	return result
}
