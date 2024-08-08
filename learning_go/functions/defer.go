package main

import (
	"fmt"
	"strconv"
)

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

// function that calculates the average of a float
func floatAv(s []float64) float64 {
	sum := 0.0

	for _, val := range s {
		sum += val
	}
	return sum / float64(len(s))
}

// a function that returns its (two) parameters in the right, numerical
func twoAscNum(s int, d int) (j int, k int) {
	if s < d {
		return s, d
	} else {
		return d, s
	}
}

// stack function
// implement pop, push.
// It should be LIFOs
func push(s int, stack *[5]int) {
	for key, value := range stack {
		if value == 0 {
			stack[key] = s
			return
		}
	}
}

func pop(stack *[5]int) {
	if stack[0] == 0 {
		fmt.Println("Stack Empty")
	}
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] != 0 {
			stack[i] = 0
			break
		}
	}
	fmt.Println(*stack)
}

// finction that converts an array to a string
func arrString(arr [5]int) string {
	var str string
	for key, val := range arr {
		str += "[" + strconv.Itoa(key) + "] : " + strconv.Itoa(val) + "\n"
	}
	return str
}

func main() {
	var stack [5]int
	fmt.Println(stack)
	(push(3, &stack))
	fmt.Println(stack)

	fmt.Println(arrString(stack))
	print("\n")
	fmt.Println(floatAv([]float64{1.3, 40.23}))
	fmt.Println(twoAscNum(3, 1))
	loopA()
	myfunc(1, 3, 53, 3)
}
