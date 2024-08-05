package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	p, n, count := 1, len(nums), 0
	var i, j int

	for i = 0; i < n; i++ {
		if nums[i] >= k {
			continue
		}

		count++
		p = nums[i]

		for j = i + 1; j < n; j++ {
			p *= nums[j]

			if p >= k {
				break
			}

			count++
		}
	}

	return count
}

func numSubarrayProductLessThanKEfficient(nums []int, k int) int {
	// Sliding window technique
	if k < 0 {
		return 0
	}

	left, count := 0, 0
	p := 1

	for right := 0; right < len(nums); right++ {
		p *= nums[right]

		// Dividing the sum by the left element until we are again in bounds
		for left <= right && p >= k {
			p /= nums[left]
			left++
		}

		count += right - left + 1
	}

	return count
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanKEfficient([]int{10, 5, 2, 6}, 100))
}
