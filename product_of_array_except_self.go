package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftProduct := make([]int, n)
	rightProduct := make([]int, n)
	leftProduct[0] = 1
	rightProduct[n-1] = 1

	for i := 1; i < n; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
		leftProduct[i] *= rightProduct[i]
	}

	return leftProduct

}

func main() {
	numArray := []int{2, 7, 11, 15}
	fmt.Println(productExceptSelf(numArray))
}
