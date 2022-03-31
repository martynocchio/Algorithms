package main

// first in - last out

type Stack struct {
	collection map[int]interface{}
	count      int
}

func NewStack() *Stack {
	return &Stack{
		collection: map[int]interface{}{},
		count:      0,
	}
}

func (s *Stack) Push(value interface{}) {
	s.collection[s.count] = value
	s.count++
}

func (s *Stack) Pop() interface{} {
	s.count--
	deletedElement := s.collection[s.count]
	delete(s.collection, s.count)
	return deletedElement
}
