// Create a simple stack which can hold a ﬁxed number of ints. It does not have to
// grow beyond this limit. Deﬁne push – put something on the stack – and pop –
// retrieve something from the stack – functions. The stack should be a LIFO (last
// in, ﬁrst out) stack.functions
//   - fixed number
//     define push. if value added exit emmidiately
package main

import (
	"fmt"
	"strconv"
)

var top int = -1

func pushin(num int, stack *[5]int) {
	if top == len(stack)-1 {
		fmt.Println("stack is full")
		return
	}
	top++
	stack[top] = num
	fmt.Println(*stack)
}

func popout(stack *[5]int) {
	if top == -1 {
		fmt.Println("stack is empty")
		return
	}
	fmt.Printf("Popped value %d at index [%d]\n", stack[top], top)
	stack[top] = 0
	top--
}

func printout(stack *[5]int) string {
	// convert the string to string
	str := ""
	for key, value := range stack {
		str += "[" + strconv.Itoa(key) + ":" + strconv.Itoa(value) + "] "

	}
	return str
}
func main() {
	var stack [5]int
	pushin(3, &stack)
	pushin(4, &stack)
	popout(&stack)
	pushin(6, &stack)
	pushin(35, &stack)
	popout(&stack)

	str := printout(&stack)
	fmt.Printf("This is the stack %s\n", str)
}
