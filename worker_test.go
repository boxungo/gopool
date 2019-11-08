package gopool

import "testing"

func TestWorker(t *testing.T) {
	jobChan := make(chan *Job)
	worker := NewWorker(jobChan)
	go worker.Start()
	i := 10
	job := NewJob(func() {
		i += 10
	})

	for j := 0; j < 5; j++ {
		jobChan <- job
	}

	// 速度这么快？怎么不会延迟呢？奇怪
	if i != 60 {
		t.Error("Worker执行失败")
	}
}
