package main

import (
	"fmt"
	"gaaaamma/vertex"
)

type ListNode[T any] struct {
	val  T
	next *ListNode[T]
}

func traverseLinkedList[T any](root *ListNode[T]) {
	traverser := root
	for traverser != nil {
		fmt.Printf("%v ", traverser.val)
		traverser = traverser.next
	}
	fmt.Println()
}

func createLinkedList[T any](s []T) *ListNode[T] {
	var dummy *ListNode[T] = nil
	traverse := dummy
	for _, val := range s {
		if dummy == nil {
			dummy = &ListNode[T]{val, nil}
			traverse = dummy
		} else {
			traverse.next = &ListNode[T]{val, nil}
			traverse = traverse.next
		}
	}
	return dummy
}
func main() {
	inputA := []int{2, 4, 1, 3, 5, 0}
	traverseLinkedList(createLinkedList(inputA))
	inputB := []string{"Apple", "Banana", "Canada", "Difference"}
	traverseLinkedList(createLinkedList(inputB))
	inputVertex := []vertex.Vertex{{X: 1, Y: 10}, {X: 2, Y: 20}, {X: 3, Y: 30}}
	traverseLinkedList(createLinkedList(inputVertex))
}
