package main

import "fmt"

func wonderfulSubstrings(word string) int64 {
	freq := make(map[int]int64)

	// The empty prefix.
	freq[0] = 1

	var mask int = 0
	var result int64 = 0

	for _, ch := range word {
		bit := int(ch - 'a')

		// Flip the value of the bit in the mask.
		mask ^= (1 << bit)

		result += freq[mask]
		freq[mask] += 1

		// Counting all the other characters that appeared for an odd number of times.
		for odd_c := 0; odd_c < 10; odd_c++ {
			result += freq[mask^(1<<odd_c)]
		}
	}

	return result
}

func main() {
	fmt.Println(wonderfulSubstrings("aba"))
	fmt.Println(wonderfulSubstrings("aabb"))
}
