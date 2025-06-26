package bubbleSort

import (
	"fmt"
	"math/rand"
	"time"
)

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

func Sort(list []int) {
	fmt.Println("> Before: ", list)
	start := time.Now()
	for {
		swapped := false // Track if any swaps occurred
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				temp := list[i+1]
				list[i+1] = list[i]
				list[i] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	endTime := time.Since(start)
	fmt.Printf("Sorted the list with %d in %fs.\n", len(list), endTime.Seconds())
	fmt.Println("> After: ", list)
}

func Example() {
	Sort(GenerateData(20))
}
