package main

import (
	"fmt"
	"sync"
	"time"
)

// basicWaitGroup demonstrates the fundamental usage of sync.WaitGroup.
// WaitGroup is used to wait for a collection of goroutines to finish executing.
//
// Key concepts:
//   - Add(n) increments the WaitGroup counter by n
//   - Done() decrements the counter by 1 (typically called with defer)
//   - Wait() blocks until the counter becomes 0
//
// Common pattern:
//  1. Call Add() before starting goroutine
//  2. Call defer Done() at the start of goroutine
//  3. Call Wait() to block until all goroutines complete
func basicWaitGroup() {
	fmt.Println("=== Basic WaitGroup Example ===")

	var wg sync.WaitGroup

	// Add 1 to the WaitGroup counter before launching goroutine
	wg.Add(1)
	go func() {
		// defer ensures Done() is called even if panic occurs
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Task 1 done (2s)")
	}()

	// Add another goroutine to the WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Task 2 done (1s)")
	}()

	// Wait blocks until WaitGroup counter reaches 0
	// This ensures both goroutines complete before proceeding
	wg.Wait()
	fmt.Println("All tasks completed!\n")
}

// waitGroupWithLoop demonstrates using WaitGroup with multiple goroutines in a loop.
// Important: Add() must be called before launching goroutines to avoid race conditions.
func waitGroupWithLoop() {
	fmt.Println("=== WaitGroup with Loop ===")

	var wg sync.WaitGroup
	numWorkers := 5

	// Add all workers at once
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// Important: Pass i as parameter to avoid closure issues
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Worker %d completed\n", id)
		}(i) // Pass i as argument
	}

	wg.Wait()
	fmt.Println("All workers completed!\n")
}

// workerPool demonstrates a worker pool pattern using WaitGroup.
// Fixed number of workers process jobs from a channel.
func workerPool() {
	fmt.Println("=== Worker Pool Pattern ===")

	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	numWorkers := 3

	// Start worker pool
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(100 * time.Millisecond)
				results <- job * 2
			}
		}(w)
	}

	// Send jobs
	go func() {
		for j := 1; j <= 9; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Close results channel after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("  %d\n", result)
	}
	fmt.Println()
}

func main() {
	basicWaitGroup()
	waitGroupWithLoop()
	workerPool()
}
