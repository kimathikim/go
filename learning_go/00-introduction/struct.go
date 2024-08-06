package main

import (
	"fmt"
	"unicode/utf8"
)

type student struct {
	name    string
	age     int
	id      int
	present bool
}

func main() {
	// Declaring a variable of type student
	var s student
	// Assigning values to the fields of the struct
	s.name = "John"
	s.age = 21
	s.id = 1001
	s.present = true
	// Printing the values of the struct fields
	fmt.Println("Name:", s.name)
	fmt.Println("Age:", s.age)
	println("ID:", s.id)
	println("Present:", s.present)
	simpleFOr()
	loopGoto()
	loopArray()
	fizzBuzz()
	tree()
	countchar("asSASA ddd dsjkdsjs dk")
	reverseString("kimathi Brian")
	print("\n")
	fmt.Println(averangeFloatSlice([]float64{1.2, 32.3}))
}

func simpleFOr() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func loopGoto() {
	i := 0
start:
	if i < 10 {

		fmt.Println(i)
		i++
		goto start
	}
}

func loopArray() {
	numbers := [10]int{}
	i := 0
start:
	if i < 10 {
		numbers[i] = i
		i++
		goto start
	}
	fmt.Println(numbers)
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func tree() {
	s := "A"
	i := 0
start:
	if i < 100 {
		for j := 0; j < i; j++ {
			fmt.Print(s)
		}
		fmt.Println("")
		i++
		goto start
	}
}

func countchar(s string) {
	byteslice := []byte(s)
	fmt.Println("Bytes: ", len(byteslice))
	byteCoubt := utf8.RuneCount(byteslice)
	fmt.Println("Runes: ", byteCoubt)
}

func reverseString(s string) {
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Println(string(r))
	for i := len(r) - 1; i >= 0; i-- {
		fmt.Printf(string(r[i]))
	}
}

func averangeFloatSlice(s []float64) float64 {
	sum := 0.0
	if len(s) == 0 {
		return 0.0
	}
	for _, val := range s {
		print(val)
		sum += val
	}
	return sum / float64(len(s))
}
