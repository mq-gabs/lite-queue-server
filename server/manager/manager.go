package manager

import (
	"errors"
	"fmt"
	"lite_queue_server/protocol"
	queue "lite_queue_server/queue"
	job "lite_queue_server/queue/job"
	"lite_queue_server/utils"
)

type QueueManager struct {
	queues map[string]*queue.Queue
}

func New() *QueueManager {
	return &QueueManager{
		queues: make(map[string]*queue.Queue),
	}
}

func (qm *QueueManager) getQueue(name string) (*queue.Queue, error) {
	q, ok := qm.queues[name]

	if !ok {
		return nil, fmt.Errorf("queue does not exists: %v", name)
	}

	return q, nil
}

func (qm *QueueManager) NewQueue(name string) error {
	if _, err := qm.getQueue(name); err == nil {
		return fmt.Errorf("queue %v already exists", name)
	}

	qm.queues[name] = queue.New()

	return nil
}

func (qm *QueueManager) Push(name string, value []byte) error {
	q, err := qm.getQueue(name)

	if err != nil {
		return err
	}

	j := job.New(value)

	q.Push(j)

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

	return utils.FlattenBytes([][]byte{
		[]byte(j.Id),
		[]byte{protocol.Seperator},
		j.Data,
	}), nil
}

func (qm *QueueManager) Ack(name, id string) error {
	q, err := qm.getQueue(name)

	if err != nil {
		return err
	}

	err = q.Ack(id)

	if err != nil {
		return err
	}

	return nil
}
