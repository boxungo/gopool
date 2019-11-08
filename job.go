package gopool

// Job 工作
type Job struct {
	f func()
}

// NewJob 新工作
func NewJob(f func()) *Job {
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
	j.f()
}
