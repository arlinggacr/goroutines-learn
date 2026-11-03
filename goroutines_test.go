package goroutines_learn

import (
	"fmt"
	"strconv"
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

func ToResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Arlingga Response"
}

func TestChannelWithParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go ToResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func DataIn(channel chan<- string) { // only send data to channel
	time.Sleep(2 * time.Second)
	channel <- "Data In Channel"
}

func DataOut(channel <-chan string) { // only got data from channel
	data := <-channel
	fmt.Println(data)
}

func TestChannelInAndOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	DataIn(channel)
	DataOut(channel)

	time.Sleep(3 * time.Second)
}

// buffered channel
func TestBufferedChannel(t *testing.T) { // get in the data buffer even the data wasnt read yet, but if the buffer is full, the data will be deadlock  until the data is read
	channel := make(chan string, 2) // buffered channel with capacity of 2
	defer close(channel)

	go func() {
		channel <- "Data 1"
		channel <- "Data 2"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done Task")
}

// range channel
// for iterating the data that is goes into channel, and the channel will be closed when the data is done
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// anonymous function to send data to channel
	go func() {
		for i := 1; i <= 5; i++ {
			channel <- "Loop in channel " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("got data", data)
	}

	fmt.Println("Done Task")
}

// select channel
// selecting the channel that is ready to be read or write, if the channel is not ready, it will be blocked until the channel is ready
// if there are multiple channels that are ready, it will select one of them randomly
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Data from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "Data from channel 2"
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
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

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Data from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Data from channel 2"
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
