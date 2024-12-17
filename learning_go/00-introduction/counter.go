// a for loop function
package main

import "fmt"

func loopOne() {
	// normal loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func gotoloop() {

	i := 0
here:

	if i < 10 {

		fmt.Println(i)
		i++
		goto here
	}
}
func main() {
	loopOne()
	println()
	gotoloop()
}
