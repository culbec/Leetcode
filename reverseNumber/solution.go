package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	// Doesn't work for all test cases.
	reversed := 0

	is_signed := x < 0

	first_digit := x % 10
	if is_signed {
		first_digit = first_digit * -1
	}
	reversed = reversed*10 + first_digit
	x /= 10

	passes := 1

	for x != 0 && passes < 9 {
		digit := x % 10

		if is_signed {
			digit = digit * -1
		}

		reversed = reversed*10 + digit

		x /= 10
		passes++
	}

	if passes == 9 && x != 0 {
		if first_digit > 2 {
			return 0
		}

		digit := x % 10

		if is_signed {
			digit = digit * -1
		}

		reversed = reversed*10 + digit
	}

	if is_signed {
		return reversed * -1
	}

	return reversed
}

func reverseOk(x int) int {
	reversed := 0

	for x != 0 {
		digit := x % 10

		reversed = reversed*10 + digit

		if reversed < math.MinInt32 || reversed > math.MaxInt32 {
			return 0
		}

		x /= 10
	}

	return reversed
}

func run() {
	fmt.Println(123, " ", reverseOk(123))
	fmt.Println(-123, " ", reverseOk(-123))
	fmt.Println(120, " ", reverseOk(120))
	fmt.Println(1534236469, " ", reverseOk(0))
	fmt.Println(-2147483412, " ", reverseOk(-2147483412))
	fmt.Println(-10, " ", reverseOk(-10))
}
