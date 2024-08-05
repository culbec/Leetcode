package main

func subarraysDivByK(nums []int, k int) int {
    count := 0
    prefixSum := 0
    prefixSumCount := make([]int, k)
    prefixSumCount[0] = 1

    for _, num := range nums {
        prefixSum = (prefixSum + num) % k
        if prefixSum < 0 {
            prefixSum += k
        }
        count += prefixSumCount[prefixSum]
        prefixSumCount[prefixSum]++
    }

    return count
}
