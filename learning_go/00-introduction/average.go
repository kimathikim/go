package main

import "fmt"

func main() {
	runes := []float64{3.2, 234.2, 323.34}
	sum := 0.0
	for _, value := range runes {
    sum += value
	}
  fmt.Printf("%d\n", int(sum) / len(runes))
}
