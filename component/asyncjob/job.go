package asyncjob

import (
	"context"
	"time"
)

// Job requirement:
// 1. Job can do something (handler)
// 2. Job can retry
// 2.1. Config retry times and duration
// 3. Should be stateful
// 4. Whe should have job manager to message jobs

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDuration(times []time.Duration)
}

const (
	defaultMaxTimeout    = time.Second * 10
	defaultMaxRetryCount = 3
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 5, time.Second * 10}
)

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeOut
	StateCompleted
	StateRetryFailed
)

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

// NewJob ...
func NewJob(handler JobHandler) *job {
	return &job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1, // retry =-1 do chỉ mới khởi tạo, chưa được retry lần nào
		stopChan:   make(chan bool),
	}
}

// Execute ...
func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	var err error
	if err = j.handler(ctx); err != nil {
		j.state = StateFailed

		return err
	}

	j.state = StateCompleted
	return nil
}

// Retry ...
func (j *job) Retry(ctx context.Context) error {
	//if j.retryIndex == len(j.config.Retries)-1 {
	//	return nil
	//}

	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	//j.state = StateRunning
	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 { // Check có phải lần retry cuối cùng không?
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed

	return err
}

// State ...
func (j *job) State() JobState {
	return j.state
}

// RetryIndex ...
func (j *job) RetryIndex() int {
	return j.retryIndex
}

// SetRetryDuration ...
func (j *job) SetRetryDuration(times []time.Duration) {
	if len(times) == 0 {
		return
	}

	j.config.Retries = times
}
