package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) append(val int) {
	newNode := &ListNode{val, nil}
	currentNode := node

	for currentNode.Next != nil {
		nextNode := currentNode.Next
		currentNode = nextNode
	}

	currentNode.Next = newNode
}

func (node *ListNode) print() {
	currentNode := node

	for currentNode.Next != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}

	fmt.Println(currentNode.Val)
}

type Solution struct{}

func (sol *Solution) addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	var carry, val int

	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		// for example. 15 = 1, 5
		carry, val = divMod(v1+v2+carry, 10)
		curr.Next = &ListNode{val, nil}

		// move to the next digit
		curr = curr.Next
	}

	return dummy.Next
}

func divMod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	l1 := &ListNode{2, nil}
	l1.append(6)
	l1.append(9)
	fmt.Println("l1 list:")
	l1.print()

	l2 := &ListNode{2, nil}
	l2.append(6)
	l2.append(9)
	fmt.Println("l2 list:")
	l2.print()

	sol := Solution{}
	l3 := sol.addTwoNumbers(l1, l2)
	fmt.Println("l3 list:")
	l3.print()

}
