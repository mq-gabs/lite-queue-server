package queue

import (
	job "lite_queue_server/queue/job"
	"testing"
)

var q *Queue = New()
var j *job.Job

func TestNew(t *testing.T) {
	q = New()

	if q.size != 0 {
		t.Errorf("size must be zero: %v", q.size)
	}

	if q.head != nil {
		t.Errorf("head must be nil: %v", q.head)
	}

	if q.tail != nil {
		t.Errorf("tail must be nil: %v", q.tail)
	}

	if q.Size() != q.size {
		t.Errorf("method Size must returns equal size | expected: %v, got: %v", q.size, q.Size())
	}
}

func TestPush(t *testing.T) {
	j = job.New([]byte("data"))

	q.Push(j)

	if q.size != 1 {
		t.Errorf("size must be one: %v", q.size)
	}

	if q.head == nil {
		t.Errorf("head must not be empty: %v", q.head)
	}

	if q.tail == nil {
		t.Errorf("tail must not be empty: %v", q.tail)
	}
}

func TestPop(t *testing.T) {
	res := q.Pop()

	if res != j {
		t.Errorf("pop must return job pushed")
	}

	if q.size != 0 {
		t.Errorf("size must be zero: %v", q.size)
	}

	if q.head != nil {
		t.Errorf("head must be empty: %v", q.head)
	}

	if q.tail != nil {
		t.Errorf("tail must be empty: %v", q.tail)
	}
}
