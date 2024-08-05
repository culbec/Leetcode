package main

import "fmt"

func pivotInteger(n int) int {
	if n < 1 {
		return -1
	}

	// Initializing a slice that will contain the partial sums of the integers.
	partialSums := make([]int, 0)
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
		partialSums = append(partialSums, sum)
	}

	// Searching for the integer.

	for x := 0; x < n; x++ {
		if partialSums[x] == sum-partialSums[x]+x+1 {
			return x + 1
		}
	}

	return -1
}

func pivotIntegerFormulas(n int) int {
	if n < 1 {
		return -1
	}

	// Sum of elements between 1 and x: (x / 2) * (1 + x)
	// Sum of elements between x and n: ((n - x + 1) / 2) (x + n)
	// Gauss sum

	for x := 1; x <= n; x++ {
		sum1x := (x * (x + 1)) / 2
		sumxn := ((n - x + 1) * (x + n)) / 2

		if sum1x == sumxn {
			return x
		}
	}

	return -1
}

func main() {
	fmt.Println(pivotInteger(8))
	fmt.Println(pivotIntegerFormulas(8))
	fmt.Println(pivotInteger(1))
	fmt.Println(pivotIntegerFormulas(1))
	fmt.Println(pivotInteger(4))
	fmt.Println(pivotIntegerFormulas(4))
}
