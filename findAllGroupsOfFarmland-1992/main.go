package main

import (
	"fmt"
	"math"
)

func dfs(land [][]int, x, y, minx, miny, maxx, maxy int, w, h int, coords []int) {
	if x < 0 || x >= h || y < 0 || y >= w || land[x][y] == 0 {
		return
	}

	// We need to find the leftmost top point in the table
	// and the rightmost bottom point.

	minx = int(math.Min(float64(minx), float64(x)))
	miny = int(math.Min(float64(miny), float64(y)))
	maxx = int(math.Max(float64(maxx), float64(x)))
	maxy = int(math.Max(float64(maxy), float64(y)))

	coords[0], coords[1] = minx, miny
	coords[2], coords[3] = maxx, maxy

	land[x][y] = 0
	dfs(land, x+1, y, minx, miny, maxx, maxy, w, h, coords)
	dfs(land, x-1, y, minx, miny, maxx, maxy, w, h, coords)
	dfs(land, x, y+1, minx, miny, maxx, maxy, w, h, coords)
	dfs(land, x, y-1, minx, miny, maxx, maxy, w, h, coords)
}

func findFarmland(land [][]int) [][]int {
	farmlands := make([][]int, 0)
	h, w := len(land), len(land[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if land[i][j] == 1 {
				coords := []int{i, i, j, j}
				dfs(land, i, j, i, j, i, j, w, h, coords)
				farmlands = append(farmlands, coords)
			}
		}
	}

	return farmlands
}

func main() {
	l1 := [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}
	l2 := [][]int{{0, 1}, {1, 0}}
	fmt.Println(findFarmland(l1))
	fmt.Println(findFarmland(l2))
}
