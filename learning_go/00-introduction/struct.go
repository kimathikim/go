package main

import "fmt"

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
