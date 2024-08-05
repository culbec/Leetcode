package main

import "fmt"

func convert_faster(s string, numRows int) string {
  rows := make([]string, numRows) // saving the strings of each row 

  converted := ""

  index := 0      // starting at position 0
  direction := 1  // 1 -> going down a row, -1 -> going up a row in the zig-zag matrix

  for _, ch := range s {
    rows[index] = rows[index] + string(ch)

    if index == 0 {
      direction = 1
    } else if index == numRows - 1 {
      direction = -1
    }

    index = index + direction
  }

  for i := 0; i < numRows; i++ {
    converted = converted + rows[i]
  }

  return converted
}

func runFaster() {
  fmt.Println(convert_faster("PAYPALISHIRING", 4))
  fmt.Println(convert_faster("PAYPALISHIRING", 3))
  fmt.Println(convert_faster("PAYPALISHIRING", 2))
}
