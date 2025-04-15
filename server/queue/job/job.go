package job

type Job struct {
	Next *Job
	Data []byte
}

func New(data []byte) *Job {
	return &Job{
		Data: data,
	}
}
