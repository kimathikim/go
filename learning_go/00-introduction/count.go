package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str,
		len([]byte(str)), utf8.RuneCount([]byte(str)))
	run := []rune(str)

  run = append(run[:4], append([]rune("abc"), run[5:]...)...)
  str = string(run)
	fmt.Println(str)
}
