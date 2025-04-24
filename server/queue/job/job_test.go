package job

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	startTime := time.Now()
	j := New([]byte("example"))

	if j.Data == nil {
		t.Error("data cannot be nil")
	}

	if !bytes.Equal(j.Data, []byte("example")) {
		t.Errorf("data is invalid: %v", j.Data)
	}

	if err := uuid.Validate(j.Id); err != nil {
		t.Errorf("invalid uuid: %v", err)
	}

	expectedTime := startTime.Add(time.Minute)
	if j.RequeueAt.Sub(expectedTime) > time.Second {
		t.Errorf("requeue at has invalid date. expected: %v, received: %v", expectedTime, j.RequeueAt)
	}

	if j.Next != nil {
		t.Error("next must be nil")
	}
}
