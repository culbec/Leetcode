package main

func reverseString(s []byte)  {
    var i, j int = 0, len(s) - 1
    var temp byte

    for i < j {
        temp = s[i]
        s[i] = s[j]
        s[j] = temp

        i++
        j--
    }
}
