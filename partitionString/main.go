package main

import (
	"fmt"
	"strings"
)

func partitionString(s string) int {
	if s == "" {
		return 0
	}

	// Trying to expand the current partition from left to right as much as possible.
	leftIndex := 0
	currentPartition := "" + string(s[leftIndex])

	partitions := 1
	leftIndex++

	for leftIndex < len(s) {
		if strings.ContainsRune(currentPartition, rune(s[leftIndex])) {
			partitions += 1
			currentPartition = ""
		}
		currentPartition = currentPartition + string(s[leftIndex])
		leftIndex++
	}

	return partitions
}

func partitionStringMap(s string) int {
	if s == "" {
		return 0
	}

	characterMap := make(map[byte]int)
	characterMap[s[0]] = 1
	partitions := 1

	for i := 1; i < len(s); i++ {
		if characterMap[s[i]] == 0 {
			characterMap[s[i]] = partitions
		} else if characterMap[s[i]] == partitions {
			partitions++
			characterMap[s[i]]++
			//clear(characterMap) // resetting the map -> this is faster
		} else {
			characterMap[s[i]] = partitions
		}
	}

	return partitions
}

func main() {
	fmt.Println(partitionString("abacaba"), " ", partitionStringMap("abacaba"))
	fmt.Println(partitionString("ssssss"), " ", partitionStringMap("ssssss"))
	fmt.Println(partitionString("hdklqkcssgxlvehva"), " ", partitionStringMap("hdklqkcssgxlvehva"))
	fmt.Println(partitionString("lvkmzlaeaxbprczpfarnlaptfvmutkfsatyywnxpmkpduwoqeeiltbdjipwihqi"), " ", partitionStringMap("lvkmzlaeaxbprczpfarnlaptfvmutkfsatyywnxpmkpduwoqeeiltbdjipwihqi"))
}
