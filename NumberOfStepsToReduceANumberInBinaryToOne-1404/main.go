func numSteps(s string) int {
    var n int = len(s)
    var num, carry int = 0, 0

    for i := n - 1; i > 0; i-- {
        if ((int(s[i] - '0') + carry) % 2)  == 1 {
            num += 2
            carry = 1
        } else {
            num++
        }
    }

    return num + carry
}
