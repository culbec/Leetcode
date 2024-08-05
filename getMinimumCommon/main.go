package main

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return -1
	}

	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}

		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return -1
}

func main() {
	fmt.Println(getCommon([]int{1, 2, 3}, []int{2, 4}))
	fmt.Println(getCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}))
}
