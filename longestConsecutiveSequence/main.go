package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	unique := make(map[int]bool)

	// Variables to hold the actual count and the temp count.
	var count int = 0
	var tempCount int = 0

	// Saving all the unique elements in a map.
	for _, num := range nums {
		unique[num] = true
	}

	// Searching for consecutive sequences by
	// searching for the start of a sequence.
	// Explanation: num - 1 doesn't exist in the map.
	for num := range unique {
		if unique[num-1] {
			continue
		}

		// We found a start of a sequence, searching for the next
		// consecutive element.
		tempCount = 1
		var toSearch int = num + 1

		for unique[toSearch] {
			tempCount++
			toSearch++
		}

		if tempCount > count {
			count = tempCount
		}
	}

	return count
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
