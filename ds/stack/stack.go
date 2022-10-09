package ds

import (
	"fmt"
)

type Stack []string

func (stack *Stack) Push(str string) {
	*stack = append(*stack, str)
}

func (stack *Stack) Pop() (string, bool) {
	length := len(*stack) - 1
	element := (*stack)[length]
	*stack = (*stack)[:length]
	return element, true
}

func (stack *Stack) Peek() string {
	length := len(*stack) - 1
	element := (*stack)[length]
	return element
}

func (stack *Stack) Size() int {
	return len(*stack)
}

func (stack *Stack) Print() {
	fmt.Println(*stack)
}
