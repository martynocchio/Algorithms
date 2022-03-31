package main

type Queue struct {
	collection []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		collection: []interface{}{},
	}
}

func (q *Queue) Enqueue(value interface{}) {
	q.collection = append(q.collection, value)
}

func (q *Queue) Dequeue() interface{} {
	deletedElement := q.collection[0]
	q.collection = q.collection[1:]
	return deletedElement
}
