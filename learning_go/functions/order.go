package main

import "fmt"

// Write a function that returns its (two) parameters in the right, numerical (ascend-
// ing) order:
// f(7,2) → 2,7
// f(2,7) → 2,7
//
func order(a int, b int)( int, int){
  if a > b{
    return b, a
  }
  return a, b
}

func main(){
  fmt.Println(order(1,4))
  fmt.Println(order(3, 1))
}
