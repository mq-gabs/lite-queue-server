package queue

import (
	"fmt"
	job "lite_queue_server/queue/job"
	"sync"
	"time"
)

type Queue struct {
	mutex    sync.Mutex
	size     uint
	head     *job.Job
	tail     *job.Job
	inFlight map[string]*job.Job
}

func New() *Queue {
	q := &Queue{
		inFlight: make(map[string]*job.Job),
	}

	go q.checkInFlights()

	return q
}

func (q *Queue) checkInFlights() {
	for k, v := range q.inFlight {
		if time.Now().After(v.RequeueAt) {
			v.Requeue()
			q.Push(v)
			delete(q.inFlight, k)
		}
	}

	time.Sleep(time.Minute)

	go q.checkInFlights()
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

	j := q.head
	q.head = q.head.Next
	q.size -= 1

	if q.size == 0 {
		q.tail = nil
	}

	q.inFlight[j.Id] = j

	return j
}

func (q *Queue) Size() uint {
	return q.size
}

func (q *Queue) Ack(id string) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	_, ok := q.inFlight[id]

	if !ok {
		return fmt.Errorf("there is no in-flight job with id: %v", id)
	}

	delete(q.inFlight, id)

	return nil
}
