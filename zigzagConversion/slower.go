// https://leetcode.com/problems/zigzag-conversion/

package main

import "fmt"

func convert_slower(s string, numRows int) string {
  if numRows == 1 || numRows > len(s) {
    return s
  }

  converted := ""
 
  matrix := make([][]rune, numRows)
  rowIndex := 0
  direction := 1 // 1 - down, -1 - up

  for i := 0; i < len(s); i++ { 
    ch := rune(s[i])

    if rowIndex == numRows {
      direction = -1
      rowIndex--
    } else if rowIndex == 0 {
      direction = 1
    }

    if direction == 1 {
      matrix[rowIndex] = append(matrix[rowIndex], ch)
      rowIndex++
    } else {
      if rowIndex == numRows - 1 {
        for j := 0; j < numRows; j++ {
          matrix[rowIndex] = append(matrix[rowIndex], 0)
        }
        rowIndex--
        i--
      } else {
        for j := 0; j < numRows - rowIndex - 2; j++ {
          matrix[rowIndex] = append(matrix[rowIndex], 0)
        }
        matrix[rowIndex] = append(matrix[rowIndex], ch)
        rowIndex--
      }
    }
  }

  for i := 0; i < numRows; i++ {
    for _, ch := range matrix[i] {
      if ch != 0 {
        converted = converted + string(ch)
      }
    }
  } 

  return converted
}

func runSlower() {
  fmt.Println(convert_slower("PAYPALISHIRING", 2))
  fmt.Println(convert_slower("PAYPALISHIRING", 3))
  fmt.Println(convert_slower("PAYPALISHIRING", 4))
}
