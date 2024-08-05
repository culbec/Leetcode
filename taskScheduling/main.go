package main

import (
	"fmt"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	// Creating a map to store the frequency of each task
	taskFreq := make(map[byte]int)
	for _, task := range tasks {
		taskFreq[task]++
	}

	// Saving the frequency of each task
	freq := make([]int, 0, len(taskFreq))
	for _, f := range taskFreq {
		freq = append(freq, f)
	}

	// Sorting the tasks in ascending order
	sort.Ints(freq)

	// We don't need to wait after the execution of the most frequent task
	mostFreq := freq[len(freq)-1] - 1
	// The number of idle times at the beginning is equal to the frequency of the most frequent task * the time of waiting
	idleSlots := mostFreq * n

	// Substracting from the idleSlots as we find lower freqs.
	for i := len(freq) - 2; i >= 0; i-- {
		idleSlots -= min(mostFreq, freq[i])
	}

	// If we still have idleSlots, we return the sum of the length of the tasks list + idleSlots to incorporate them in the number of minIntervals.
	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}

	return len(tasks)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}
