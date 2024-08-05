package main

import "fmt"

func dfs(grid [][]byte, start_i, start_j int, n, m int) {
	// Base case: out of bounds or positioning in water
	if (0 > start_i || start_i >= n) || (0 > start_j || start_j >= m) || grid[start_i][start_j] == '0' {
		return
	}

	// Completing all the other adjacent cells with 0
	grid[start_i][start_j] = '0'
	dfs(grid, start_i+1, start_j, n, m)
	dfs(grid, start_i-1, start_j, n, m)
	dfs(grid, start_i, start_j+1, n, m)
	dfs(grid, start_i, start_j-1, n, m)
}

func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j, n, m)
				count++
			}
		}
	}

	return count
}

func main() {
	mat := [][]byte{[]byte{'1', '1', '1', '1', '0'}}
	fmt.Println(numIslands(mat))
}
