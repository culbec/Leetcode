package main

import (
	"cmp"
	"fmt"
	"slices"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	var score int64 = 0
	var loss int = 0

	slices.SortStableFunc(happiness, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	for i := 0; i < len(happiness) && k > 0; i++ {
		if happiness[i]-loss > 0 {
			score += int64(happiness[i] - loss)
		}

		k--
		loss++
	}

	return score
}

func main() {
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))
	fmt.Println(maximumHappinessSum([]int{1, 1, 1, 1}, 2))
	fmt.Println(maximumHappinessSum([]int{2, 3, 4, 5}, 1))
	fmt.Println(maximumHappinessSum([]int{2, 98, 45}, 1))
}
