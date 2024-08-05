func dictContainsWord(word string, wordDict []string) bool {
    for _, w := range wordDict {
        if word == w {
            return true
        }
    }
    return false
}

func wordBreak(s string, wordDict []string) []string {
    var dp map[int][]string = make(map[int][]string)
    var n int = len(s)

    for start := n; start >= 0; start-- {
        var validSentences []string = make([]string, 0)

        for end := start; end < n; end++ {
            var curr string = s[start:end+1]

            if dictContainsWord(curr, wordDict) {
                if end == n - 1 {
                    validSentences = append(validSentences, curr)
                } else {
                    var sentencesFromNextIndex []string = dp[end + 1]
                    for _, sentence := range sentencesFromNextIndex {
                        validSentences = append(validSentences, curr + " " + sentence)
                    }
                }
            }
        }
        dp[start] = validSentences
    }

    return dp[0]
}
