package main

import (
	"fmt"
	"log"
  "file"

)

func main() {

	sum := 0

  for i := 0; i < 10; i++ {

  sum += i;

  }
  fmt.Println(sum)

  if err := file.Chmod(0664); err != nil {
    log.Println(err)
    println(err)
  }
}
