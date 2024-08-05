package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}

	if s == "" || p == "" {
		return false
	}

	pIndex := 0
	sIndex := 0

	for sIndex < len(s) && pIndex < len(p)-1 {
		if p[pIndex+1] != '*' {
			if p[pIndex] == '.' {
				sIndex++
			} else {
				if s[sIndex] != p[pIndex] {
					return false
				}
			}
			pIndex++
		} else if p[pIndex+1] == '*' {
			for sIndex < len(s) && s[sIndex] == p[pIndex] {
				sIndex++
			}
			pIndex += 2
		}
	}

	// Checking the last group of the reg expr
	if len(p)%2 == 0 {
		if p[pIndex-1] != '*' {
			if p[pIndex-2] == '.' {
				sIndex++
			} else {
				if s[sIndex] != p[pIndex-2] {
					return false
				}
			}
			pIndex++
		} else {
			for sIndex < len(s) && s[sIndex] == p[pIndex-1] {
				sIndex++
			}
			pIndex += 2
		}
	} else {
		if p[pIndex] == '.' {
			sIndex++
		} else if s[sIndex] != p[pIndex] {
			return false
		}
		pIndex++
	}

	if sIndex < len(s) && pIndex >= len(p) {
		return false
	}

	return true
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
}
