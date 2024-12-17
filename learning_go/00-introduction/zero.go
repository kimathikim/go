package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j int = 3, 4
	var f float64 = math.Sqrt(float64(i*i + j + j))
	z := uint(f)

	fmt.Println(i, j, z, f)
}
