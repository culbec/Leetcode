package main

import "fmt"

func customSortString(order string, s string) string {
	if order == "" || s == "" {
		return s
	}

	orderedS := ""

	// Getting the frequency of every character in the string 's'.

	frequencies := make(map[rune]int)

	for _, ch := range s {
		frequencies[ch] += 1
	}

	// Adding each character in the ordered string based on it's index in the 'order' string.
	for _, ch := range order {
		for frequencies[ch] > 0 {
			orderedS += string(ch)
			frequencies[ch]--
		}
	}

	// Adding the rest of the characters, the order does not matter.
	for ch, val := range frequencies {
		for val > 0 {
			orderedS += string(ch)
			val--
		}
	}

	return orderedS
}

func main() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("bcafg", "abcd"))
}
