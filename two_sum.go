package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				if i != j {
					solution := make([]int, 2)
					solution[0] = i
					solution[1] = j
					return solution
				}
			}
		}
	}
	solution := make([]int, 2)
	return solution
}

func main() {
	numArray := []int{2, 7, 11, 15}
	fmt.Println(twoSum(numArray, 9))
}