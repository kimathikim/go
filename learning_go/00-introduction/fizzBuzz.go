package main

import "fmt"

func fizzBuz(){
  i := 100

  for n := 0; n <= i; n++ {
    if n % 3 == 0 {
      fmt.Println("Fizz")
      continue
    }
    if n % 5 == 0{
      fmt.Println("BUzz")
      continue
    }
    if n % 5 == 0 && n % 3 == 0{
      fmt.Println("fizzBuzz")
      continue
    }
    fmt.Println(n)
  }
}

func main(){
  fizzBuz()
}
