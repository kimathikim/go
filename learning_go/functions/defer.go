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

func main() {
	fmt.Println(floatAv([]float64{1.3, 40.23}))
	fmt.Println(twoAscNum(3, 1))
	loopA()
	myfunc(1, 3, 53, 3)
}
