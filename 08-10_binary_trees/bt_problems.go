package main

import (
	"fmt"
	"github.com/dasizeman/binarytree"
	"github.com/dasizeman/tools"
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

	// Problem 8 "superbalanced"
	//------------------------------
	validBST.Right.Right.InsertLeft(85)
	validBST.Left.Left.InsertLeft(10)
	fmt.Printf("Valid: %t\n", isBinarySearchTreeValid(validBST))
	fmt.Printf("Superbalanced: %t\n", isBinaryTreeSuperBalanced(validBST))
	fmt.Printf("Second largest: %d\n", secondLargestItemInBST(validBST).(int))

}

func isBinarySearchTreeValid(tree *binarytree.Node) bool {
	// We will implement a slightly modified depth-first
	// search, keeping track of the acceptable bounds for
	// BST values as we go

	type NodeInfo struct {
		ptr      *binarytree.Node
		min, max int
	}

	nodeStack := &Stack{}

	nodeStack.push(NodeInfo{tree, math.MinInt64, math.MaxInt64})

	for !nodeStack.isEmpty() {
		node := nodeStack.pop().(NodeInfo)

		if node.ptr.Data.(int) < node.min {
			return false
		}
		if node.ptr.Data.(int) > node.max {
			return false
		}

		if node.ptr.Left != nil {
			nodeStack.push(NodeInfo{node.ptr.Left, node.min,
				node.ptr.Data.(int)})
		}
		if node.ptr.Right != nil {
			nodeStack.push(NodeInfo{node.ptr.Right, node.ptr.Data.(int),
				node.max})
		}

	}

	return true
}

// Note from interviewcake solution:
// We could short circuit and not have to traverse the whole tree,
// if we added an additional condition: never more than 2 unique
// depths, and checked both conditions at every iteration
func isBinaryTreeSuperBalanced(tree *binarytree.Node) bool {
	// Again we will do a DFS, keeping track of the max and min
	// heights of all leaves.  If max-min > 1, the tree is not superbalanced

	nodeStack := &Stack{}
	min := math.MaxInt64
	max := math.MinInt64

	// Quick local struct because no tuples in golang :(
	type Node struct {
		ptr   *binarytree.Node
		depth int
	}

	nodeStack.push(Node{tree, 0})

	for !nodeStack.isEmpty() {
		node := nodeStack.pop().(Node)

		if node.ptr.Left != nil {
			nodeStack.push(Node{node.ptr.Left, node.depth + 1})
		}
		if node.ptr.Right != nil {
			nodeStack.push(Node{node.ptr.Right, node.depth + 1})
		}

		// We only care about leaf nodes
		if !(node.ptr.Left == nil &&
			node.ptr.Right == nil) {
			continue
		}

		max = tools.IntMax(max, node.depth)
		min = tools.IntMin(min, node.depth)

	}

	return (max - min) <= 1
}

func secondLargestItemInBST(tree *binarytree.Node) interface{} {
	searchNode := tree
	for searchNode.Right.Right != nil {
		searchNode = searchNode.Right
	}

	if searchNode.Right.Left == nil {
		return searchNode.Data
	}

	searchNode = searchNode.Right.Left

	for searchNode.Right != nil {
		searchNode = searchNode.Right
	}

	return searchNode.Data
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

func (stack *Stack) height() int {
	return stack.topIdx
}
