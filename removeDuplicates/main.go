package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numsSet := make([]int, 1)
	k := 1

	numsSet[0] = nums[0]
	currElem := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != currElem {
			numsSet = append(numsSet, nums[i])
			currElem = nums[i]
			k++
		}
	}

	for i := 1; i < len(numsSet); i++ {
		nums[i] = numsSet[i]
	}

	return k
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	currElem := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] != currElem {
			currElem = nums[i]
			nums[k] = currElem

			k++
		}
	}

	return k
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}), " ", removeDuplicates2([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}), " ", removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}), " ", removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
