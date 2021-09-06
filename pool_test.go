package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New" // default value jika datanya nil
		},
	}
	var group sync.WaitGroup

	pool.Put("Muhammmad")
	pool.Put("Ilham")
	pool.Put("Nadhif")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}
	group.Wait()
	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
