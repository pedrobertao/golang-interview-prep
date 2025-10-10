package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateData creates a slice of integers from 0 to numOfInt-1 and shuffles them randomly.
// This is useful for testing sorting algorithms with randomized data.
//
// Parameters:
//   - numOfInt: the number of integers to generate
//
// Returns:
//   - A shuffled slice of integers from 0 to numOfInt-1
func GenerateData(numOfInt int) []int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	data := make([]int, numOfInt)
	for i := range data {
		data[i] = i
	}

	r.Shuffle(numOfInt, func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	return data
}

// Sort performs an optimized bubble sort on the given slice of integers in-place.
// Bubble sort repeatedly steps through the list, compares adjacent elements and swaps them
// if they are in the wrong order. The pass through the list is repeated until no swaps are needed.
//
// Algorithm characteristics:
//   - Time Complexity: O(n²) worst/average case, O(n) best case (already sorted)
//   - Space Complexity: O(1) - sorts in-place
//   - Stable: Yes - maintains relative order of equal elements
//   - In-place: Yes - modifies the input slice directly
//
// Optimization: Uses a 'swapped' flag to detect when the list is sorted and exit early.
//
// The function also measures and prints the sorting time and the sorted result.
//
// Parameters:
//   - list: the slice of integers to sort (modified in-place)
func Sort(list []int) {
	start := time.Now()
	for {
		swapped := false // Track if any swaps occurred in this pass
		for i := 0; i < len(list)-1; i++ {
			// Compare adjacent elements
			if list[i] > list[i+1] {
				// Swap if they are in wrong order
				temp := list[i+1]
				list[i+1] = list[i]
				list[i] = temp
				swapped = true
			}
		}
		// If no swaps occurred, the list is sorted
		if !swapped {
			break
		}
	}
	endTime := time.Since(start)
	fmt.Printf("Sorted the list with %d in %fs.\n", len(list), endTime.Seconds())
	fmt.Println("> After: ", list)
}

func main() {
	Sort(GenerateData(20))
}
