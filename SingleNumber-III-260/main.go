package main

func singleNumber(nums []int) []int {
  var result []int = make([]int, 2)
  result[0], result[1] = 0, 0
  
  var xor int = 0

  for _, num := range nums {
    xor ^= num
  }

  var k int = 0
  for xor != 0 && xor & 1 != 1 {
    k++
    xor >>= 1
  }

  for _, num := range nums {
    if num & (1 << k) == 1 {
      result[0] ^= num
    } else {
      result[1] ^= num
    } 
  }

  return result
}
