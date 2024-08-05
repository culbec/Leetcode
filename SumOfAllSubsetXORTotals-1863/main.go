package main

func subsetXORSum(nums []int) int {
    var xorTotal int = 0

    for _, num := range nums {
        // Save the bits set to 1 in the result.
        xorTotal |= num
    }

    // Shift the result by n - 1.
    return xorTotal << (len(nums) - 1)
}
