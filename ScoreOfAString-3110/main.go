package main

func abs (a int) int {
    if a < 0 {
        return -a
    }

    return a
}

func scoreOfString(s string) int {
    var score int = 0
    var n int = len(s)

    for i := 0; i < n - 1; i++ {
        score += abs(int(s[i]) - int(s[i + 1]))
    }

    return score
}
