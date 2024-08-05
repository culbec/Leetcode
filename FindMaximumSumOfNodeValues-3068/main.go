package main

import (
	"fmt"
	"math"
	"sort"
)

func dpTabulation(nums []int, k int, edges [][]int) int64 {
	n := len(nums)

	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, 2)
	}

	dp[n][1] = 0 // no operation on an odd number of elements
	dp[n][0] = math.MinInt

	for i := n - 1; i >= 0; i-- {
		for isEven := 0; isEven <= 1; isEven++ {
			// Case 1: we can perform the xor operation
			var xor int64 = dp[i+1][isEven^1] + int64(nums[i]^k)

			// Case 2: we cannot perform the xor operation
			var noXor int64 = dp[i+1][isEven] + int64(nums[i])

			// Saving the maximum value.
			dp[i][isEven] = int64(math.Max(float64(xor), float64(noXor)))
		}
	}

	return dp[0][1]
}

func greedySort(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	change := make([]int, n)
	var sum int64 = 0

	for i := 0; i < n; i++ {
		change[i] = (nums[i] ^ k) - nums[i]
		sum += int64(nums[i])
	}

	// Sorting the array in reverse order to pair the bigger values.
	sort.Slice(change, func(i, j int) bool {
		return change[i]-change[j] > 0
	})

	// Trying to pair the numbers and get the maximum sum.
	for i := 0; i < n; i += 2 {
		if i+1 >= n {
			break
		}

		var pairSum int64 = int64(change[i]) + int64(change[i+1])
		if pairSum > 0 {
			sum += pairSum
		}
	}

	return sum
}

func greedyMaxMin(nums []int, k int, edges [][]int) int64 {
	var sum int64 = 0
	var count int = 0

	mini, maxi := math.MaxInt, math.MinInt

	for num := range nums {
		xor := num ^ k
		sum += int64(num)
		change := xor - num

		if change > 0 {
			mini = int(math.Min(float64(mini), float64(change)))
			sum += int64(change)
			count++
		} else {
			maxi = int(math.Max(float64(maxi), float64(change)))
		}
	}

	if count%2 == 0 {
		return sum
	}

	return int64(math.Max(float64(sum-int64(mini)), float64(sum+int64(maxi))))
}

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	dp, gs, gmm := dpTabulation(nums, k, edges),
		greedySort(nums, k, edges),
		greedyMaxMin(nums, k, edges)

	return int64(dp - gs - gmm)
}

func main() {
	fmt.Println(maximumValueSum([]int{1, 2, 1}, 3, [][]int{{0, 1}, {0, 2}}))
}
