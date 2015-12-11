// Package queue provides a simple thread safe queue.
// provides simple first in, first out functionality.
package queue

import (
	"errors"
	"sync"
)

var mutex = &sync.Mutex{}

// Queue Type for storing items in
// first in, first out order.
type Queue struct {
	queue []interface{}
}

// Pop the item at the front of the queue.
func (q *Queue) Pop() (interface{}, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if !q.hasLength() {
		return nil, errors.New("No Items in Queue")
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item, nil
}

// Push an item to the back of the queue.
func (q *Queue) Push(item interface{}) {
	mutex.Lock()
	defer mutex.Unlock()

	q.queue = append(q.queue, item)
}

// Return the current length of the queue.
func (q *Queue) Length() int {
	return len(q.queue)
}

func (q *Queue) hasLength() bool {
	return q.Length() > 0
}
