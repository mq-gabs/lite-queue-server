package queue

import (
	job "lite_queue_server/queue/job"
	"sync"
)

type Queue struct {
	size  uint
	head  *job.Job
	tail  *job.Job
	mutex sync.Mutex
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(j *job.Job) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.size += 1

	if q.head == nil {
		q.head = j
		q.tail = j
		return
	}

	q.tail.Next = j
	q.tail = j
}

func (q *Queue) Pop() *job.Job {
	q.mutex.Lock()
	defer q.mutex.Unlock()

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
