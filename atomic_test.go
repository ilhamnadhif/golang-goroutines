package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var group sync.WaitGroup
	//var mutex sync.Mutex
	var x int64 = 0
	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				//mutex.Lock()
				//x = x + 1
				//mutex.Unlock()
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter : ", x)
}
