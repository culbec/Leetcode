package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) < 1 {
		return -1
	}

	// If a candidate reaches 0 votes, than it is guarranteed that it's appearances can't get over N / 2.
	// Boyer-Moore majority voting algorithm

	votes := 0
	candidate := -1

	for i := 0; i < len(nums); i++ {
		if votes == 0 {
			candidate = nums[i]
			votes = 1
		} else {
			if candidate == nums[i] {
				votes++
			} else {
				votes--
			}
		}
	}

	voteCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate {
			voteCount++
		}
	}

	if voteCount > int(len(nums)/2) {
		return candidate
	}

	return -1
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
