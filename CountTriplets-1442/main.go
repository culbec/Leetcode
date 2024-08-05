package main

func countTriplets(arr []int) int {
    var n int = len(arr)
    var count, prefix int = 0, 0
 
    var counts, totals map[int]int = make(map[int]int), make(map[int]int)
    counts[0] = 1

    for i := 0; i < n; i++ {
        prefix ^= arr[i]
        count += counts[prefix] * i - totals[prefix]
        
        counts[prefix]++
        totals[prefix] += i + 1
    }

    return count
}
