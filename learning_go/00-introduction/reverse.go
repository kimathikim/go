package main

import (
	"fmt"
)

// reverse reboot to toober
func main() {
	str := "reboot"
  newstr := ""
	runes := []rune(str)

	for key, _ := range runes {
    newstr += string(runes[len(str) - key - 1])
	}
  fmt.Println(newstr)
}
// or 
// import "fmt"
// func main() {
// s := "foobar"
// a := [] rune (s)
// ← Again a conversion
// f o r i, j := 0, len (a)-1 ; i < j ; i, j = i+1, j-1 {
// a[i], a[j] = a[j], a[i]
// ← Parallel assignment
// }
// fmt.Printf("%s\n", s t r i n g (a))
// }
