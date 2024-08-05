package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	k := 1
	maxApps := 2
	currApps := 1

	currElem := nums[0]

	for i := 1; i < len(nums); i++ {
		if currApps == maxApps {
			for i < len(nums) && nums[i] == currElem {
				i++
			}

			if i < len(nums) {
				currApps = 0
				currElem = nums[i]
				i--
			}
		} else if nums[i] != currElem {
			currApps = 1
			currElem = nums[i]

			nums[k] = currElem
			k++
		} else if nums[i] == currElem {
			currApps++
			nums[k] = currElem
			k++
		}
	}

	return k
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
