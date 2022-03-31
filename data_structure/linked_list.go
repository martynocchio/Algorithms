package main

type LinkedList struct {
	count int
	head  *Node
	tail  *Node
}

type Node struct {
	value interface{}
	next  *Node
}

func (l *LinkedList) Add(value interface{}) {
	newNode := &Node{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.count++
}

func (l *LinkedList) Remove(element interface{}) (int, bool) {
	current := l.head
	var previous *Node

	index := 0
	for current != nil {
		if current.value == element {
			if previous != nil {
				previous.next = current.next

				if current.next == nil {
					l.tail = previous
				}
			} else {
				l.head = current.next

				if l.head == nil {
					l.tail = nil
				}
			}
			l.count--
			return index, true
		}
		index++
		previous = current
		current = current.next
	}
	return index, false
}