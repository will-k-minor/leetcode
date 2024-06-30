package main

func maxProfit(prices []int) int {
	// keep track of the max
	// keep track of the min on the left
	// if you find a new min, use that one
	// compare the current profit to max profit

	maxProfit := 0
	currentMin := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < currentMin {
			currentMin = prices[i]
			continue
			//found a new min. ignore calculating current profit
		}

		currentProfit := prices[i] - currentMin

		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}

	return maxProfit
}
