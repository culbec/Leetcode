func subsets(nums []int) [][]int {
    var subsets [][]int = make([][]int, 0)
    var subset []int = make([]int, 0)
    var numsSize int = len(nums)

    var bkt func(int, int) 
    bkt = func(first, length int) {
        if len(subset) == length {
            subsets = append(subsets, append([]int{}, subset...))
            return
        }

        for i := first; i < numsSize; i++ {
            subset = append(subset, nums[i])
            bkt(i + 1, length)
            subset = subset[:len(subset) - 1]
        }
    }

    for l := 0; l <= numsSize; l++ {
        bkt(0, l)
    }

    return subsets
}
