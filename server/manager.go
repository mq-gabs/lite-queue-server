package main

import (
	"errors"
	"fmt"
)

type QueueManager struct {
	queues map[string]*Queue
}

func NewQueueManager() *QueueManager {
	return &QueueManager{
		queues: make(map[string]*Queue),
	}
}

func (qm *QueueManager) getQueue(name string) (*Queue, error) {
	q, ok := qm.queues[name]

	if !ok {
		return nil, fmt.Errorf("queue does not exists: %v", name)
	}

	return q, nil
}

func (qm *QueueManager) Push(name string, value []byte) error {
	q, err := qm.getQueue(name)

	if err != nil {
		return err
	}

	j := Job{
		Data: value,
	}

	q.Push(&j)

	return nil
}

func (qm *QueueManager) Pop(name string) ([]byte, error) {
	q, err := qm.getQueue(name)

	if err != nil {
		return nil, err
	}

	j := q.Pop()

	if j == nil {
		return nil, errors.New("job is empty")
	}

	return j.Data, nil
}
