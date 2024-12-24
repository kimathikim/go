package main

import "fmt"

// create a simple map function

func mapfunc(list interface{}, f interface{}) interface{} {
	switch v := list.(type) {
	case []int:
		result := make([]int, len(v))
		for key, value := range v {
			result[key] = f.(func(int) int)(value)
		}
		return result
	case []string:
		result := make([]string, len(v))
		for key, value := range v {
			result[key] = f.(func(string) string)(value)
		}
		return result
	default:
		return nil
	}
}

func main() {
	intList := []int{1, 4, 7, 8, 9, 0, 3, 2}
	intOutput := mapfunc(intList, func(i int) int {
		return i * 2
	}).([]int)
	fmt.Println(intOutput)

	stringList := []string{"a", "b", "c", "d"}
	stringOutput := mapfunc(stringList, func(s string) string {
		return s + s
	}).([]string)
	fmt.Println(stringOutput)
}
