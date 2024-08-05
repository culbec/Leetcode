package main

import (
	"fmt"
	"math"
)

func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}

	/*
		The number of steps if n = 2^k => 2^(k + 1) - 1

		Example: 4 = 100
		* 100 -> 101
		* 101 -> 111
		* 111 -> 110
		* 110 -> 010
		* 010 -> 011
		* 011 -> 001
		* 001 -> 000

		Example: 4 = 1000
		* 1000 -> 1001
		* 1001 -> 1011
		* 1011 -> 1010
		* 1010 -> 1110
		* 1110 -> 1111
		* 1111 -> 1101
		* 1101 -> 1100
		* 1100 -> 0100
		* 0100 -> 0101
		* 0101 -> 0111
		* 0111 -> 0110
		* 0110 -> 0010
		* 0010 -> 0011
		* 0011 -> 0001
		* 0001 -> 0000

		If n = 2^k + 1 => we treat each 1 bit as a power of 2

		Example: 5 = 101
		1 -> 1 op
		100 -> 7 op

		=> 8 op
	*/

	// This variable holds the number of operations done.
	var ops int = 0

	// Counting the number of bits of the binary representation of n.

	noBits := 0

	for nn := n; nn != 0; nn >>= 1 {
		noBits++
	}
	// This variable holds the mask for extracting the current right bit.
	mask := 0x00000001 << (noBits - 1)

	// This variable holds the sign of the addition.
	sign := -1

	// Computing the number of operations for each bit.
	for i := 0; i < noBits; i++ {
		bit := (n & mask) >> (noBits - i - 1)

		if bit == 1 {
			ops += sign * (-1) * (int(math.Pow(2, float64(noBits-i))) - 1)
			sign *= -1
		}

		mask >>= 1
	}

	return ops
}

func main() {
	//fmt.Println(minimumOneBitOperations(3))
	//fmt.Println(minimumOneBitOperations(6))
	fmt.Println(minimumOneBitOperations(9))
}
