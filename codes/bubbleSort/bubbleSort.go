package bubbleSort

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func GenerateData(numOfInt int) []int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Build sequential slice
	data := make([]int, numOfInt)
	for i := range data {
		data[i] = i
	}

	// Shuffle with local rand
	r.Shuffle(numOfInt, func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	return data
}

func Sort(list []int) {
	start := time.Now()
	sorted := false
	for !sorted {
		for i, v := range list {
			if i >= 0 && i < len(list)-1 {
				if v > list[i+1] {
					temp := list[i+1]
					list[i+1] = v
					list[i] = temp
				}
			}
		}
		sorted = sort.SliceIsSorted(list, func(i, j int) bool {
			return list[j] > list[i]
		})
	}
	endTime := time.Since(start)
	fmt.Printf("Sorted the list with %d in %fs.\n", len(list), endTime.Seconds())
}
