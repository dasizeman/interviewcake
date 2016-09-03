package main

import (
	"fmt"
	"github.com/dasizeman/binarytree"
	"math"
)

func main() {

	// Problem 9 BST Checker
	// --------------------------

	// Valid BST
	//       50
	//     /    \
	//    30     80
	//   /  \   /  \
	//  20  40 70  90
	validBST := binarytree.Create(50)
	validBST.InsertLeft(30)
	validBST.InsertRight(80)
	validBST.Left.InsertLeft(20)
	validBST.Left.InsertRight(40) // Can switch this to 60 for invalid
	validBST.Right.InsertLeft(70)
	validBST.Right.InsertRight(90)

	fmt.Printf("Valid: %t\n", isBinaryTreeValid(validBST))

}

func isBinaryTreeValid(tree *binarytree.Node) bool {
	// We will implement a slightly modified depth-first
	// search, keeping track of the acceptable bounds for
	// BST values as we go

	nodeStack := &Stack{}
	maxMinStack := &Stack{}

	nodeStack.push(tree)
	maxMinStack.push(MaxMinPair{math.MinInt64, math.MaxInt64})

	for !nodeStack.isEmpty() {
		node := nodeStack.pop().(*binarytree.Node)
		maxMin := maxMinStack.pop().(MaxMinPair)
		//fmt.Printf("%d\n", node.Data.(int))
		//fmt.Printf("%d,%d\n", maxMin.min, maxMin.max)

		if node.Data.(int) < maxMin.min {
			return false
		}
		if node.Data.(int) > maxMin.max {
			return false
		}

		if node.Left != nil {
			nodeStack.push(node.Left)
			maxMinStack.push(MaxMinPair{maxMin.min, node.Data.(int)})
		}
		if node.Right != nil {
			nodeStack.push(node.Right)
			maxMinStack.push(MaxMinPair{node.Data.(int), maxMin.max})
		}
	}

	return true
}

// Simple "generic" stack
// -----------------------------------------

// Stack is a stack that can hold any value, but
// you have to type assert yourself
type Stack struct {
	data   []interface{}
	topIdx int
}

func (stack *Stack) peek() interface{} {
	return stack.data[stack.topIdx-1]
}

func (stack *Stack) pop() interface{} {
	value := stack.peek()
	stack.topIdx--
	return value
}

func (stack *Stack) push(value interface{}) {
	stack.topIdx++
	if stack.topIdx > len(stack.data) {
		stack.data = append(stack.data, value)
	} else {
		stack.data[stack.topIdx-1] = value
	}
}

func (stack *Stack) isEmpty() bool {
	return stack.topIdx <= 0
}

// Max/min pair
// -----------------------------------------------

// MaxMinPair stores a max and min float value
type MaxMinPair struct {
	min, max int
}
