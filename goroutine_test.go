package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func DiplayNumber(number int) {
	fmt.Println("Diplay", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DiplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}