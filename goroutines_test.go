package goroutines_learn

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func TestCreateGoroutines(t *testing.T) {
	go HelloWorld()
	fmt.Println("Test completed")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number:", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	// send data
	// channel <- "Hello, Channel!"

	// make channel inside the variabel
	// data := <-channel

	// or just put it
	// fmt.Println(<-channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "arlingga"
		fmt.Println("Selesai test create channel") // depends the long of the data, if this longer than this can be go first than the channel
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
