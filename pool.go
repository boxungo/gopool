package gopool

import "sync"

// Pool 池
type Pool struct {
	maxWorkers int
	jobChan    chan *Job
	wg         sync.WaitGroup
}

// NewPool 新池
func NewPool(maxWorkers, maxJobs int) *Pool {
	pool := &Pool{
		maxWorkers: maxWorkers,
		jobChan:    make(chan *Job, maxJobs),
	}

	go pool.start()

	return pool
}

// Submit 把任务提交到池子
func (p *Pool) Submit(job *Job) {
	if job != nil {
		p.jobChan <- job
	}
}

// start 池子开启
func (p *Pool) start() {
	for i := 0; i < p.maxWorkers; i++ {
		worker := NewWorker(p.jobChan)
		go worker.Start()
	}
}

// StopWait 停止工作协程,并且等待响应结果
func (p *Pool) StopWait() {

}
