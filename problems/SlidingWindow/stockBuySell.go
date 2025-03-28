package main

import (
	"fmt"
	"math"
)

/*
In the Sliding Window (Two-Pointer) Approach, we use:

Left Pointer (minPrice) → Tracks the lowest price seen so far (best buying price).
Right Pointer (price[i]) → Moves through the array to find the best selling price.

🔹 Summary of How Two Pointers Work
1️⃣ Left Pointer (minPrice) → Keeps track of the lowest price seen so far.
2️⃣ Right Pointer (price[i]) → Moves forward, checking if selling today gives a higher profit.
3️⃣ Update maxProfit whenever a new higher profit is found.
4️⃣ Continue till the end of the array.

⏳ Time Complexity
O(N) → Single pass through the array.
O(1) space → Uses only two variables (minPrice, maxProfit).
*/
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Maximum Profit:", maxProfit(prices)) // Output: 5
}

func maxProfitBruteForce(prices []int) int {
	maxProfit := 0
	n := len(prices)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // Update minPrice
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit // Update maxProfit
			}
		}
	}
	return maxProfit
}
