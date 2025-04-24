package job

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	Next      *Job
	Id        string
	RequeueAt time.Time
	Data      []byte
}

func New(data []byte) *Job {
	return &Job{
		Id:        uuid.New().String(),
		RequeueAt: time.Now().Add(time.Minute),
		Data:      data,
	}
}

func (j *Job) Requeue() {
	j.RequeueAt = time.Now().Add(time.Minute)
}
