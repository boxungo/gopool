package gopool

import "sync"

// Pool 池
type Pool struct {
	maxWorkers uint
	maxJobs    uint
	jobChan    chan *Job
	wg         sync.WaitGroup
}

// NewPool 新池
// maxJobs 如果 > 0, 在队列中还有任务未执行完时终止程序会丢数据
func NewPool(opt Option) *Pool {
	if opt.MaxWorkers <= 0 {
		opt.MaxWorkers = 1
	}
	pool := &Pool{
		maxWorkers: opt.MaxWorkers,
		maxJobs:    opt.MaxJobs,
		jobChan:    make(chan *Job, opt.MaxJobs),
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
	for i := 0; i < int(p.maxWorkers); i++ {
		worker := NewWorker(p.jobChan)
		go worker.Start()
	}
}

// StopWait 停止工作协程,并且等待响应结果
func (p *Pool) StopWait() {

}
