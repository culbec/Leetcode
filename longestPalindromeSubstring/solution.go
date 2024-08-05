package main

import "fmt"

func isPalindrome(slice string) bool {
	left := 0               // left pointer
	right := len(slice) - 1 // right pointer

	for left < right {
		if slice[left] != slice[right] {
			return false
		}

		left += 1
		right -= 1
	}
	return true
}

func longestPalindrome(s string) string {
	var sLength int = len(s) // saving the length of the string

	for i := 0; i < sLength; i++ {
		for j := 0; j <= i; j++ {
			var slice string = s[j:(sLength - i + j)]

			if isPalindrome(slice) {
				return string(slice)
			}
		}
	}

	return ""
}

type indexPair struct {
	leftIndex  int
	rightIndex int
}

func evenPalindrome(s string) indexPair {
	maxLength := 1
	evenIndexPair := indexPair{
		leftIndex:  0,
		rightIndex: 0,
	}

	for i := 0; i < len(s); i++ {
		left := i
		right := i + 1

		for left >= 0 && right < len(s) && s[left] == s[right] {
			length := right - left + 1

			if length >= maxLength {
				maxLength = length
				evenIndexPair.leftIndex = left
				evenIndexPair.rightIndex = right
			}

			left--
			right++
		}
	}

	return evenIndexPair
}

func oddPalindrome(s string) indexPair {
	maxLength := 1
	oddIndexPair := indexPair{
		leftIndex:  0,
		rightIndex: 0,
	}

	for i := 0; i < len(s); i++ {
		left := i
		right := i

		for left >= 0 && right < len(s) && s[left] == s[right] {
			length := right - left + 1

			if length >= maxLength {
				maxLength = length
				oddIndexPair.leftIndex = left
				oddIndexPair.rightIndex = right
			}

			left--
			right++
		}
	}
	return oddIndexPair
}

func longestPalindromeFaster(s string) string {
	if len(s) <= 1 {
		return s
	}

	evenIndexPair := evenPalindrome(s)
	oddIndexPair := oddPalindrome(s)

	if (evenIndexPair.rightIndex - evenIndexPair.leftIndex + 1) > (oddIndexPair.rightIndex - oddIndexPair.leftIndex + 1) {
		return s[evenIndexPair.leftIndex : evenIndexPair.rightIndex+1]
	}

	return s[oddIndexPair.leftIndex : oddIndexPair.rightIndex+1]
}

func main() {
	s1 := "babad"
	s2 := "cbbd"
	s3 := ""

	fmt.Println(longestPalindrome(s1))
	fmt.Println(longestPalindrome(s2))
	fmt.Println(longestPalindrome(s3))

	fmt.Println(longestPalindromeFaster(s1))
	fmt.Println(longestPalindromeFaster(s2))
	fmt.Println(longestPalindromeFaster(s3))
}
