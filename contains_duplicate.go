package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	cache := make(map[int]int)
	for _, num := range nums {
		fmt.Println(num)
		cache[num] += 1
		if cache[num] > 1 {
			return true
		}
	}
	return false
}

func main() {
	numArray := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(numArray))
}
