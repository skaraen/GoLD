package utils

import "container/list"

type Queue[T any] struct {
	list *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: list.New()}
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.PushBack(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.list.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := q.list.Front()
	q.list.Remove(elem)
	return elem.Value.(T), true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.Len() == 0
}