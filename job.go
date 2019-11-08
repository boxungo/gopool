package gopool

// Job 工作
type Job struct {
	f func() error
}

// NewJob 新工作
func NewJob(f func() error) *Job {
	if f == nil {
		return nil
	}
	job := &Job{
		f: f,
	}
	return job
}

// Execute 开始工作
func (j *Job) Execute() {
	err := j.f()
	if err != nil {
		// todo something
	}
}
