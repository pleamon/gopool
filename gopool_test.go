package gopool

import (
	"testing"
	"time"
)

func TestGopool(t *testing.T) {
	worker := Worker{}

	// 设置callback与协程池大小
	worker.Init(func(t *testing.T, counter int) {
		t.Log(counter)
		time.Sleep(1 * time.Second)
	}, 10)

	for i := 0; i < 10000; i++ {
		worker.Push(t, i)
	}

	worker.Start()
	worker.Wait()
}
