func specialArray(nums []int) int {
    var n int = len(nums)
    var count int = 0

    for x := 0; x <= n; x++ {
        count = 0

        for _, num := range nums {
            if num >= x {
                count++
            }
        }
        if count == x {
            return x
        }
    }

    return -1
}
