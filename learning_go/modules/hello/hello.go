package main

import (
  "fmt"
  "log"
  "trial/greetings"
)
func main (){
  log.SetPrefix("Greetings: ")
  log.SetFlags(0)

  message, err:= greetings.Hello("Denis")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(message)

}
