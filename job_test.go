package gopool

import "testing"

func TestJob(t *testing.T) {
	i := 10
	f := func() {
		i += 10
	}

	job := NewJob(f)
	job.Execute()

	if i != 20 {
		t.Error("Job测试不通过")
	}
}

func TestNil(t *testing.T) {
	job := NewJob(nil)
	if job != nil {
		t.Error("Job接收nil参数错误")
	}
}
