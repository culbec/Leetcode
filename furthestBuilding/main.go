package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// Greedy approach to use as much ladders as possible.
	// When we have no ladders, we replace the minimum height with bricks.
	currHeight := heights[0]
	noBuildings := 1
	gaps := &IntHeap{}
	heap.Init(gaps)

	for i := 1; i < len(heights); i++ {
		if currHeight > heights[i] {
			noBuildings++
		} else {
			if ladders == 0 && bricks == 0 {
				return noBuildings
			} else if ladders > 0 {
				gaps.Push(heights[i] - currHeight)
				ladders--
			} else {
				if gaps.Len() > 0 {
					x := gaps.Pop()
					if bricks < x.(int) {
						return noBuildings
					}

					bricks -= x.(int)
				} else {
					if bricks < heights[i]-currHeight {
						return noBuildings
					}
					bricks -= heights[i] - currHeight
				}
			}
			noBuildings++
		}

		currHeight = heights[i]
	}

	return noBuildings
}

func main() {
	fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	fmt.Println(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	fmt.Println(furthestBuilding([]int{14, 3, 19, 3}, 17, 0))
}
