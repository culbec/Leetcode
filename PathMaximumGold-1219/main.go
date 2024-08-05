package main

import (
	"math"
)

func recFind(grid [][]int, n, m, i, j int) int {
	if (i < 0 || i >= n) || (j < 0 || j >= m) || grid[i][j] == 0 {
		return 0
	}

	gold, old := 0, grid[i][j]
	dirs := []int{0, 1, 0, -1, 0}

	grid[i][j] = 0

	for dir := 0; dir < 4; dir++ {
		gold = int(math.Max(float64(gold), float64(recFind(grid, n, m, dirs[dir]+i, dirs[dir+1]+j))))
	}

	grid[i][j] = old
	return gold + old
}

func getMaximumGold(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	maxi := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			maxi = int(math.Max(float64(maxi), float64(recFind(grid, n, m, i, j))))
		}
	}

	return maxi
}
