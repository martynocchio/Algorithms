package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	length int
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) append(node *node) {
	if l.head == nil {
		l.head = node
		l.length++
		return
	}
	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = node
	l.length++
}

func (l linkedList) PrintList() {
	for l.head != nil {
		fmt.Printf(" %d ->", l.head.value)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l *linkedList) reverseLinkedList() {
	cursor := l.head
	var prev *node = nil

	for cursor != nil {
		next := cursor.next
		cursor.next = prev
		prev = cursor
		cursor = next
	}
	l.head = prev
}

func main() {
	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node3 := &node{value: 3}
	node4 := &node{value: 4}

	ll := NewLinkedList()
	ll.append(node1)
	ll.append(node2)
	ll.append(node3)
	ll.append(node4)
	ll.PrintList()
	ll.reverseLinkedList()
	ll.PrintList()
}
