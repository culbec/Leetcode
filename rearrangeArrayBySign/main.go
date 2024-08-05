package main

import "fmt"

func rearrangeArray(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	positive := make([]int, 0)
	negative := make([]int, 0)

	for _, elem := range nums {
		if elem < 0 {
			negative = append(negative, elem)
		} else {
			positive = append(positive, elem)
		}
	}

	j := 0
	k := 0

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = positive[j]
			j++
		} else {
			nums[i] = negative[k]
			k++
		}
	}

	return nums
}

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
	fmt.Println(rearrangeArray([]int{-1, 1}))
}
