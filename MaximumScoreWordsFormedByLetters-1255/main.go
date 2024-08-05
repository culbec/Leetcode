var finalScore int
var freq []int

func isValidUsage(usedLetters []int) bool {
    for c := 0; c < 26; c++ {
        if freq[c] < usedLetters[c] {
            return false
        }
    }
    return true
}

func backtrack(n int, words []string, score, usedLetters []int, totalScore int) {
    if n == -1 {
        if finalScore < totalScore {
            finalScore = totalScore
        }

        return
    }

    backtrack(n - 1, words, score, usedLetters, totalScore)
    var l int = len(words[n])

    for i := 0; i < l; i++ {
        var c int = int(words[n][i] - 'a')
        usedLetters[c]++;
        totalScore += score[c]
    }

    if isValidUsage(usedLetters) {
        backtrack(n - 1, words, score, usedLetters, totalScore)
    }

    for i := 0; i < l; i++ {
        var c int = int(words[n][i] - 'a')
        usedLetters[c]--
        totalScore -= score[c]
    }
}

func maxScoreWords(words []string, letters []byte, score []int) int {
    var n int = len(words);
    freq = make([]int, 26)
    var usedLetters []int = make([]int, 26)
    finalScore = 0
    
    for _, c := range letters {
        freq[int(c - 'a')]++
    }

    backtrack(n - 1, words, score, usedLetters, 0)
    return finalScore
}
