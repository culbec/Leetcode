package main

func maxSubarrayLength(nums []int, k int) int {
	left, right, n, result := 0, 0, len(nums), 0
	freqs := make(map[int]int)

	for left < n {
		for right < n {
			freqs[nums[right]]++
			if freqs[nums[right]] > k {
				break
			}
			right++
		}

		if result < right-left {
			result = right - left
		}

		if right >= n {
			break
		}

		for left < n && freqs[nums[right]] > k {
			freqs[nums[left]]--
			left++
		}
		right++
	}

	return result
}
