package main

func findMax(nums []int) int {
	maxi := nums[0]

	for i := 1; i < len(nums); i++ {
		if maxi < nums[i] {
			maxi = nums[i]
		}
	}

	return maxi
}

func countSubarrays(nums []int, k int) int64 {
	maxi, n := findMax(nums), len(nums)

	start, end, count := 0, 0, 0
	var result int64 = 0

	for end = 0; end < n; end++ {
		// Growing the size of the array while the count is not k.
		if nums[end] == maxi {
			count++
		}

		// Shrinking the size of the array from the left while the count is k.
		// This results in counting all the arrays that shall contain the max element
		// for k times, containing the sequence that has the maxi element k times.

		for count == k {
			if nums[start] == maxi {
				count--
			}

			start++
		}

		result += int64(start)
	}

	return result
}
