package main

import "fmt"

func Buddy(start, limit int) []int {
  delSum := 0
  delX   := 0
  
  for x:= start; x <= limit ; x++ {
    delX = 0
    for i:= 1; i < x; i++ {
        if x % i == 0 {
          delX+= i
        }
    }
    delSum = 0
    for i:= 1; i < delX - 1; i++ {
      if (delX - 1)% i == 0 {
        delSum += i
        }
    }
    if delSum - 1== x{
      ans := []int{x, delX - 1}
      return ans
    } 
  }
  return nil
}

func	main() {
	fmt.Println(Buddy(2693, 7098))
}