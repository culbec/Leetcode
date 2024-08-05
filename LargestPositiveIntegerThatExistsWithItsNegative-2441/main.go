package main

import (
	"fmt"
	"sort"
)

func findMaxK(nums []int) int {
	sort.Ints(nums)

	l, r := 0, len(nums)-1

	for l < r && nums[l] < 0 {
		if nums[r] == -nums[l] {
			return nums[r]
		} else if nums[r] > -nums[l] {
			r -= 1
		} else {
			l += 1
		}
	}

	return -1
}

func main() {
	fmt.Println(findMaxK([]int{-1, 2, -3, 3}))
	fmt.Println(findMaxK([]int{-1, 10, 6, 7, -7, 1}))
	fmt.Println(findMaxK([]int{-10, 8, 6, 7, -2, -3}))
}
