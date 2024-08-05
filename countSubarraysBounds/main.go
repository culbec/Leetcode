package main

import (
	"fmt"
	"math"
)

func countSubarrays(nums []int, minK int, maxK int) int64 {
	if maxK < minK {
		return 0
	}

	left, right := 0, 0
	n := len(nums)
	var count int64 = 0
	foundMinK, foundMaxK := -1, -1

	for right = 0; right < n; right++ {
		// Checking if the current number is in the range.
		if nums[right] >= minK && nums[right] <= maxK {
			if nums[right] == minK {
				foundMinK = right
			}
			if nums[right] == maxK {
				foundMaxK = right
			}

			// Min because we want the minimum value where a bound has been found.
			if foundMinK != -1 && foundMaxK != -1 {
				count += int64(int(math.Min(float64(foundMinK), float64(foundMaxK))) - left + 1)
			}
		} else {
			left = right + 1
			foundMinK = -1
			foundMaxK = -1
		}
	}

	return count
}

func main() {
	fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	fmt.Println(countSubarrays([]int{1, 1, 1, 1}, 1, 1))
}
