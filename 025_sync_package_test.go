package go_hello_world

import (
	"sync"
	"testing"
	"time"
)

// sync包下并发控制
func TestCounter(t *testing.T) {
	var counter int
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)     // 这里去掉后非 5000
	t.Logf("counter = %d", counter) // counter = 4773 < 5000
}

// sync.Mutex 锁机制
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	var counter int
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)     // 这里去掉后非 5000
	t.Logf("counter = %d", counter) // 5000
}

// sync.WaitGroup 等待完成后执行
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	var counter int
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter) // 5000
}
