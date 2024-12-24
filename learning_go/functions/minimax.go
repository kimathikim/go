package main

import "fmt"

// functions to find max and min values
func max(list []int) int {
	result := list[0]
	for _, value := range list {
    if value > result{
      result = value
    }
	}
	return result
}
func mini(list []int) int {
	result := list[0]
	for _, value := range list {
    if value < result{
      result = value
    }
	}
	return result
}

func main() {
	list := []int{13,5654,5,2,5654544}
	fmt.Println(max(list))
	fmt.Println(mini(list))
}
