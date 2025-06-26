package main

import (
	"fmt"
	"gopreinterview/codes/bubbleSort"
)

func example_bubbleSort() {
	// Simple
	fmt.Println("====================Bubble Sort ")
	// Small Data Set
	bubbleSort.Sort(bubbleSort.GenerateData(100))
	// Massive Data Set
	bubbleSort.Sort(bubbleSort.GenerateData(100_000))
	fmt.Println("====================Bubble Sort")
}

func main() {
	fmt.Println("Here you will find several example of codes for interviews")
	// example_bubbleSort()
}
