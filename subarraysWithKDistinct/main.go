package main

import "fmt"

func atMostK(nums []int, k int) int {
	left, right, count, distinct := 0, 0, 0, 0
	n := len(nums)

	freqs := make(map[int]int)

	for right = 0; right < n; right++ {
		if freqs[nums[right]] == 0 {
			distinct++
		}

		freqs[nums[right]]++

		for distinct > k {
			freqs[nums[left]]--

			if freqs[nums[left]] == 0 {
				distinct--
			}

			left++
		}

		count += right - left + 1
	}

	return count
}

func subarraysWithKDistinct(nums []int, k int) int {
	return atMostK(nums, k) - atMostK(nums, k-1)
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
}
