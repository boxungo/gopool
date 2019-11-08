package gopool

import "testing"

func TestOption(t *testing.T) {
	opt := Option{}

	if opt.MaxWorkers != 0 {
		t.Error("Option的值错误")
	}
}
