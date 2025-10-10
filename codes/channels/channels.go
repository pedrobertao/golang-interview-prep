package main

import (
	"fmt"
	"time"
)

// bufferedChannelSync demonstrates synchronous usage of a buffered channel.
// A buffered channel allows sending values without blocking until the buffer is full.
//
// Key concepts:
//   - Buffered channel with capacity of 3
//   - Can send 3 values without a receiver being ready
//   - defer close() ensures channel is closed when function exits
func bufferedChannelSync() {
	fmt.Println("Buffered Channel Sync ==========")

	// Create buffered channel with capacity of 3
	buffChannel := make(chan int, 3)
	defer close(buffChannel)

	// Send 3 values - none of these block because buffer has space
	buffChannel <- 1
	buffChannel <- 2
	buffChannel <- 3

	// Receive the 3 values
	fmt.Println(<-buffChannel)
	fmt.Println(<-buffChannel)
	fmt.Println(<-buffChannel)

}
// bufferedChannelAsync demonstrates asynchronous usage of a buffered channel with goroutines.
// Multiple goroutines send values to the channel at different times.
//
// Key concepts:
//   - Multiple goroutines sending to the same channel
//   - Channel is closed by the last sender (important to avoid panic)
//   - for-range loop receives until channel is closed
//   - Demonstrates proper coordination between senders and receivers
func bufferedChannelAsync() {
	fmt.Println("Buffered Channel ASync ==========")
	buffChannel := make(chan int, 3)

	// Goroutine 1: Sends after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		buffChannel <- 1
	}()

	// Goroutine 2: Sends after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		buffChannel <- 2
	}()

	// Goroutine 3: Sends after 3 seconds and closes the channel
	// Important: Close from sender side to avoid "send on closed channel" panic
	go func() {
		time.Sleep(3 * time.Second)
		buffChannel <- 3
		close(buffChannel) // Safe to close here - we know this is the last send
	}()

	// Receive values until channel is closed
	// The range loop exits when the channel is closed and empty
	for v := range buffChannel {
		fmt.Println("Received:", v)
	}
	fmt.Println("Finished")

}

// unbufferedChannel demonstrates an unbuffered channel with timeout protection.
// Unbuffered channels (capacity 0) require both sender and receiver to be ready simultaneously.
//
// Key concepts:
//   - Unbuffered channel: make(chan int) - no capacity specified
//   - Requires synchronization: sender blocks until receiver is ready
//   - select statement allows non-blocking operations and timeout handling
//   - time.After() provides timeout protection to prevent infinite blocking
func unbufferedChannel() {
	fmt.Println("Unbuffered Channel Async ==========")

	// Create unbuffered channel (no buffer)
	ch := make(chan int)

	// Goroutine sends a value after 3 seconds
	go func() {
		fmt.Println("Sleeping...")
		time.Sleep(3 * time.Second)
		ch <- 1 // This will block until main function is ready to receive
	}()

	// Use select to handle multiple channel operations
	select {
	case v := <-ch:
		// Received value successfully
		fmt.Println("Value received", v)
		close(ch)
		// Note: 'break' here is unnecessary - select only executes once
		break
	case <-time.After(5 * time.Second):
		// Timeout case - prevents infinite blocking if sender never sends
		fmt.Println("Timeout!")
	}
}

func main() {
	bufferedChannelSync()
	bufferedChannelAsync()
	unbufferedChannel()
}
