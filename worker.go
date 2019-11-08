package gopool

// Worker 工作者
type Worker struct {
	jobChan chan *Job
}

// NewWorker 新工作协程
func NewWorker(jobChan chan *Job) *Worker {
	worker := &Worker{
		jobChan: jobChan,
	}
	return worker
}

// Start 启动工作协程
func (w *Worker) Start() {
	for {
		select {
		case job, ok := <-w.jobChan:
			if ok {
				job.Execute()
			}
		}
	}
}
