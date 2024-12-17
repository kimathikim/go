package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

// a loop in disguise
func loop(n int) int {
	fmt.Println(n)

here:
	if n > 2 {
    n--

  fmt.Println(n)
		goto here
	}
  return n
}

func main() {
	fmt.Println(loop(7))
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
