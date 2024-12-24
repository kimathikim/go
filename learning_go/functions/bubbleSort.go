package main

import (
  "fmt"
  "even"
)

func bubble(lists []int) []int {
	length := len(lists)
	for i := 0; i < length-1; i++ {
		swapped := false
		for j := 0; j < length-i-1; j++ {
			if lists[j] > lists[j+1] {
				lists[j], lists[j+1] = lists[j+1], lists[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return lists
}

func main() {
	list := []int{66, 34, 25, 12, 22, 11, 90}
	sortedList := bubble(list)
	fmt.Println("Sorted list:", sortedList)
}
