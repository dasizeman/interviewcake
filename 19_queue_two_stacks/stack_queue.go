package main

import (
	"fmt"
	"github.com/dasizeman/tools"
)

// StackQueue is a queue made from two stacks
type StackQueue struct {
	enqueueStack, dequeueStack tools.Stack
}

func (queue *StackQueue) enqueue(item interface{}) {
	queue.enqueueStack.Push(item)
}

func (queue *StackQueue) dequeue() interface{} {
	if queue.dequeueStack.IsEmpty() &&
		!queue.enqueueStack.IsEmpty() {
		queue.swapStacks()
	}

	if queue.dequeueStack.IsEmpty() &&
		queue.enqueueStack.IsEmpty() {
		return nil
	}

	return queue.dequeueStack.Pop()
}

func (queue *StackQueue) swapStacks() {
	for !queue.enqueueStack.IsEmpty() {
		queue.dequeueStack.Push(queue.enqueueStack.Pop())
	}
}

func (queue *StackQueue) IsEmpty() bool {
	return queue.dequeueStack.IsEmpty() &&
		queue.enqueueStack.IsEmpty()
}

func main() {
	queue := &StackQueue{}
	for i := 0; i < 10; i++ {
		num := tools.RandomInt(1, 10)
		fmt.Printf("Pushing %d\n", num)
		queue.enqueue(num)
	}

	fmt.Printf("\n")

	for !queue.IsEmpty() {
		num := queue.dequeue().(int)
		fmt.Printf("Popping %d\n", num)
	}
}
