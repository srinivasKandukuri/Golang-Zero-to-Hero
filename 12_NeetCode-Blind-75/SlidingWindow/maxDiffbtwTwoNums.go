package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxDiff([]int{2, 3, 10, 6, 4, 8, 1})) // Output: 8 (Buy at 2, Sell at 10)
	fmt.Println(maxDiff([]int{7, 1, 5, 3, 6, 4}))     // Output: 5 (Buy at 1, Sell at 6)
	fmt.Println(maxDiff([]int{7, 6, 4, 3, 1}))        // Output: 0 (No profit)
	fmt.Println(maxDiff([]int{5}))                    // Output: 0 (Not enough prices)
}

func maxDiff(prices []int) int {
	if len(prices) < 2 {
		return 0 // No valid transaction possible
	}

	maxProfit := 0
	minPrice := math.MaxInt32

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
