func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}

func commonChars(words []string) []string {
    var n int = len(words)
    var commonChars []int = make([]int, 26)
    var currentChars []int

    for _, ch := range words[0] {
        commonChars[int(ch) - int('a')]++
    }

    for i := 1; i < n; i++ {
        currentChars = make([]int, 26)
        
        for _, ch := range words[i] {
            currentChars[int(ch) - int('a')]++
        }

        for letter := 0; letter < 26; letter++ {
            commonChars[letter] = min(commonChars[letter], currentChars[letter])
        }
    }

    result := make([]string, 0)
    for letter := 0; letter < 26; letter++ {
        for i := 0; i < commonChars[letter]; i++ {
            ch := int('a') + letter
            result = append(result, string(rune(ch)))
        }
    }

    return result
}
