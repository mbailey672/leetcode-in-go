package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	currentSum := 0
	bestSum := -2147483648

	for i, _ := range nums {
		if currentSum+nums[i] > nums[i] {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}

		if currentSum > bestSum {
			bestSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}
	return bestSum
}

func main() {
	numArray := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(numArray))
}
