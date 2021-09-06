package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ilham Nadhif"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	go func() {
		data := <-channel
		fmt.Println(data)
		fmt.Println("Selesai Menerima data")
	}()

	time.Sleep(5 * time.Second)
}

func GiveMeRespone(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ilham Nadhif"
	fmt.Println("Selesai mengirim data ke channel")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespone(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	channel <- "Ilham Nadhif"
	fmt.Println("Selesai Mengirim data")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
	fmt.Println("Selesai Menerima data")
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)
	go func() {
		channel <- "Muhammad"
		channel <- "Ilham"
		channel <- "Nadhif"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()
	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go GiveMeRespone(channel1)
	go GiveMeRespone(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channle 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channle 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go GiveMeRespone(channel1)
	go GiveMeRespone(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channle 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channle 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
