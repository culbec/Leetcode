package main

import (
	"sort"
	"strconv"
)

type Athlete struct {
	Score int
	Rank  int
}

func findRelativeRanks(scores []int) []string {
	var n int = len(scores)

	var ranks []string = make([]string, n)
	var athletes []Athlete = make([]Athlete, n)

	// Saving the athletes for sorting.
	for i, score := range scores {
		athletes[i] = Athlete{
			Score: score,
			Rank:  i,
		}
	}

	// Sorting the athletes based on their scores.
	sort.Slice(athletes, func(i, j int) bool {
		return athletes[i].Score > athletes[j].Score
	})

	for i, athlete := range athletes {
		if i == 0 {
			ranks[athlete.Rank] = "Gold Medal"
		} else if i == 1 {
			ranks[athlete.Rank] = "Silver Medal"
		} else if i == 2 {
			ranks[athlete.Rank] = "Bronze Medal"
		} else {
			ranks[athlete.Rank] = strconv.Itoa(i + 1)
		}
	}

	return ranks
}
