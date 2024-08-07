package main

import "fmt"

func loopA() {
	// Defer statements are executed in LIFO
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// using variadic functions
func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func main() {
	loopA()
	myfunc(1, 3, 53, 3)
}
