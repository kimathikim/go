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
