// creating a simple variadic funtion for sum
package main

import "fmt"

func sum(num ...int) int {
	// add all the numbers in the num slice
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}

func main(){
  fmt.Printf("Sum of 1, 4, 5 is %v", sum(1, 4 ,5))
}
