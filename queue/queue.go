package queue

import (
	"errors"
)

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Size gets the number of items in queue
func (q Queue) Size() int {
	return len(q.items)
}

// Check if queue is empty
func (q Queue) isEmpty() bool {
	return q.Size() == 0
}

// Get first item in queue
func (q Queue) Peek() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("queue is empty")
	}

	return q.items[0], nil
}

// Enqueue function, add item in queue, O(1) or O(n)
func (q *Queue) Enqueue(n int) {
	q.items = append(q.items, n)
}

// Dequeue function, removes first item in queue
func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("queue is empty")
	}

	itemRemoved := q.items[0]
	q.items = q.items[1:]

	return itemRemoved, nil
}
