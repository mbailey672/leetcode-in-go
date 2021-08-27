package main

import (
	"fmt"
)

//O(n^2) brute force
func maxProfit(prices []int) int {
	maximizedProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if (prices[j] - prices[i]) > maximizedProfit {
				maximizedProfit = prices[j] - prices[i]
			}
		}
	}
	return maximizedProfit
}

func maxProfitOptimized(prices []int) int {
	smallestPrice := prices[0]
	maximizedProfit := 0
	for i, _ := range prices {
		if prices[i] < smallestPrice {
			smallestPrice = prices[i]
		} else if currentProfit := prices[i] - smallestPrice; currentProfit > maximizedProfit {
			maximizedProfit = currentProfit
		}
	}
	return maximizedProfit
}

func main() {
	numArray := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(numArray))
	fmt.Println(maxProfitOptimized(numArray))
}
