package main

import "sync"

type Job struct {
	Next *Job
	Data []byte
}

type Queue struct {
	size  uint
	head  *Job
	tail  *Job
	Mutex sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(n *Job) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.size += 1

	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.Next = n
	q.tail = n
}

func (q *Queue) Pop() *Job {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	if q.size == 0 {
		return nil
	}

	h := q.head
	q.head = q.head.Next
	q.size -= 1

	if q.size == 0 {
		q.tail = nil
	}

	return h
}

func (q *Queue) Size() uint {
	return q.size
}
