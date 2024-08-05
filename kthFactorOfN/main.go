package main

import "fmt"

func kthFactor(n int, k int) int {
	if n < 1 || k < 1 {
		return -1
	}

	size := 0
	factors := make([]int, 0)
	d := 1

	for d*d <= n {
		if n%d == 0 {
			if size == k-1 {
				return d
			}
			factors = append(factors, d)
			size++
		}
		d++
	}

	// A perfect square would have an odd number of factors, and any other number an even number of factors.
	if (d-1)*(d-1) == n {
		size = (size-1)*2 + 1
	} else {
		size = size * 2
	}

	// Checking if k is greater than the total number of factors
	if size < k {
		return -1
	}

	return n / factors[size-k]
}

func kthFactorLessSpace(n int, k int) int {
	if n < 1 || k < 1 {
		return -1
	}

	size := 0
	d := 1

	for d*d <= n {
		if n%d == 0 {
			size++
		}

		if size == k {
			return d
		}

		d++
	}

	total := size * 2

	if (d-1)*(d-1) == n {
		total--
		size--
	}

	if total < k {
		return -1
	}

	d--
	for d >= 1 {
		if n%d == 0 {
			size++
		}

		if size == k {
			return n / d
		}

		d--
	}

	return -1
}

func main() {
	fmt.Println(kthFactorLessSpace(12, 3)) // [1, 2, 3, 4, 6, 12]
	fmt.Println(kthFactorLessSpace(7, 2))  // [1, 7]
	fmt.Println(kthFactorLessSpace(4, 4))  // [1, 2, 4]
	fmt.Println(kthFactorLessSpace(36, 5)) // [1, 2, 3, 4, 6, 8, 12, 18, 36]
	fmt.Println(kthFactorLessSpace(49, 2)) // [1, 7, 49]
	fmt.Println(kthFactorLessSpace(81, 3)) // [1, 3, 9, 27, 81]
	fmt.Println(kthFactorLessSpace(16, 6)) // [1, 2, 4, 8, 16]
	fmt.Println(kthFactorLessSpace(1, 1))  // [1]
	fmt.Println(kthFactorLessSpace(2, 2))  // [1, 2]
	fmt.Println(kthFactorLessSpace(100, 7))
}
