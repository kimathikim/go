# functions

## functions that take variable number of a parameter are called `variadic` functions

```go
func myfunc(arg ...int){}
```

- arg- tells the function to it will receive variable number of arg

```go

package main

import "fmt"
func myfunc(arg ...int){
for _, n := range arg{
fmt.Printf("And the number is: %d\n", n)
}
}
func main(){
myfunc()
}
```
## Defer

They are written read in `LIFO` format this make it easier to manage resources

## Variadic functions
> This are functions that take a variable number of parameter

```go 
// How they are defined
func function(arg ...int) {
  // do something here
}
```

The arg is a slice of integers
