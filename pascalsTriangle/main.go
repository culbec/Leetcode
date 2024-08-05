package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex < 0 || rowIndex > 33 {
		return []int{}
	}

	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i < rowIndex; i++ {
		row[i] = 1
		for j := i; j > 0; j-- {
			row[j] = row[j] + row[j-1]
		}
	}

	row[rowIndex] = 1

	return row
}

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(21))
	fmt.Println(getRow(33))
}
