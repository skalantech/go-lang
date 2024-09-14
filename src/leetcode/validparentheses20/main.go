package main

import "fmt"

// ////////////////////////////////////
type Item rune

type Stack struct {
	items [MAX]Item
	top   int
}

const MAX = 10000

func NewStack() *Stack {
	return &Stack{top: 0}
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) IsFull() bool {
	return s.top == MAX
}

func (s *Stack) Push(item Item) bool {
	if s.top < MAX {
		s.items[s.top] = item
		s.top++
		return true
	}
	return false
}

func (s *Stack) Peek() Item {
	if s.top > 0 {
		return s.items[s.top-1]
	}
	return '0' // Possible bug here!
}

func (s *Stack) Pop() bool {
	if s.top > 0 {
		s.top--
		return true
	}
	return false
}

type Solution struct{}

func (sol *Solution) IsValid(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.Push(Item(s[i]))
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			switch s[i] {
			case ')':
				if stack.Peek() == '(' {
					stack.Pop()
				} else {
					stack.Push(Item(s[i]))
				}
			case ']':
				if stack.Peek() == '[' {
					stack.Pop()
				} else {
					stack.Push(Item(s[i]))
				}
			case '}':
				if stack.Peek() == '{' {
					stack.Pop()
				} else {
					stack.Push(Item(s[i]))
				}
			}
		}
		fmt.Printf("%c", s[i])
	}
	return stack.IsEmpty()
}

///////////////////////////////////////////////////

type Node struct {
	Val  rune
	Next *Node
}

func (node *Node) Push(data rune) *Node {
	return &Node{Val: data, Next: node}
}

func (node *Node) Pop() (data rune, n *Node) {
	if node == nil {
		fmt.Println("stack overflow")
		panic("stack overflow")
	}

	return node.Val, node.Next
}

func isValid_Node(s string) bool {
	var stack *Node
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = stack.Push(char)
		default:
			if stack == nil {
				return false
			}
			top, newStack := stack.Pop()
			stack = newStack
			switch char {
			case ')':
				if top != '(' {
					return false
				}
			case ']':
				if top != '[' {
					return false
				}
			case '}':
				if top != '{' {
					return false
				}
			}
		}
	}
	return stack == nil
}

func PrintStack(top *Node) {
	for top != nil {
		fmt.Printf("%c ", top.Val)
		top = top.Next
	}
	fmt.Println()
}

// ////////////////////////////////////////////
func push(top **Node, data rune) {
	node := &Node{data, *top}
	*top = node
}

func pop(top **Node) rune {
	if *top == nil {
		fmt.Println("stack overflow")
		panic("stack overflow")
	}

	topNode := *top
	res := topNode.Val
	*top = topNode.Next
	return res
}

func isValid(s string) bool {
	var stack *Node
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			push(&stack, char)
		default:
			if stack == nil {
				return false
			}
			top := pop(&stack)
			switch char {
			case ')':
				if top != '(' {
					return false
				}
			case ']':
				if top != '[' {
					return false
				}
			case '}':
				if top != '{' {
					return false
				}
			}
		}
	}
	return stack == nil
}

func printStack(top **Node) {
	t := *top
	for t != nil {
		tmpNode := t
		fmt.Printf("%c ", tmpNode.Val)
		t = tmpNode.Next
	}
	fmt.Println()
}

func main() {
	s := "{{{[[]]}}}"
	sol := &Solution{}
	ans := "false"
	if sol.IsValid(s) {
		ans = "true"
	}
	fmt.Println(ans)

	fmt.Println(isValid_Node(s))
	node := &Node{Val: 'a', Next: &Node{Val: 'b', Next: &Node{Val: 'c', Next: nil}}}
	PrintStack(node)

	fmt.Println(isValid(s))
	printStack(&node)
}

/*
In the Go code, the Stack class is represented as a struct with an array of Item values and a top variable to keep track of the top index. The functions are implemented as methods on the Stack struct.

The Solution class is converted to a Solution struct, and the IsValid function is implemented as a method on that struct.

The main function is similar, with the necessary adjustments for Go syntax and using fmt.Println instead of cout.

Note: The C++ code had a possible bug in the Stack::Peek() function where it assigned '0' to items[0]. In the Go code, I've maintained the same behavior, but it's worth reviewing that logic to ensure it behaves as expected.
*/
