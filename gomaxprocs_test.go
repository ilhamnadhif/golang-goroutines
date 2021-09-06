package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	var group sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	var group sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCpu)

	//runtime.GOMAXPROCS(200)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoroutine)
	group.Wait()
}
