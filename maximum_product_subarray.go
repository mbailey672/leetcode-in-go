package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProduct(nums []int) int {
	minProduct := nums[0]
	maxProduct := nums[0]
	maximumProduct := nums[0]

	n := len(nums)

	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			minProduct, maxProduct = maxProduct, minProduct
		}
		minProduct = min(nums[i], minProduct*nums[i])
		maxProduct = max(nums[i], maxProduct*nums[i])

		if maxProduct > maximumProduct {
			maximumProduct = maxProduct
		}
	}
	return maximumProduct
}

func main() {
	numArray := []int{3, -1, 4}
	fmt.Println(maxProduct(numArray))
}
