package main

import "fmt"

func treeeA() {
	letter := ""
	for i := 0; i < 100; i++ {
		letter += "A"
		fmt.Println(letter)
	}
}
func main(){
  treeeA()
}
