func isPalindrome(s string) bool {
    var i, j int = 0, len(s) - 1

    for i < j {
        if s[i] != s[j] {
            return false
        }

        i++
        j--
    }

    return true
}

func partition(s string) [][]string {
    var result [][]string
    var curr []string
    var sSize int = len(s)

    var bkt func(int)

    bkt = func(first int) {
        if first >= sSize {
            result = append(result, append([]string(nil), curr...))
            return
        }

        for i := first; i < sSize; i++ {
            ss := s[first : i + 1]

            if isPalindrome(ss) {
                curr = append(curr, ss)
                bkt(i + 1)
                curr = curr[:len(curr) - 1]
            }
        }
    }

    bkt(0)
    return result
}
