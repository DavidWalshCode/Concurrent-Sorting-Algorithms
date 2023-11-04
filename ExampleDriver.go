package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	const numRuns = 3
	rand.Seed(420) // adds determinism to Slice generation
	fmt.Println("Sorting 300 million numbers...")
	times := make([]time.Duration, numRuns)
	for i := 0; i < numRuns; i++ {
		slice := generateSlice(30000) //takes 5-10 seconds to sort for optimised algorithms
		startTime := time.Now()

		//Insert Algorithm Here
		mySort(slice)

		time := time.Since(startTime)

		if sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] <= slice[j] }) {
			fmt.Println("Sorted, Algorithm functional")
			times[i] = time
			fmt.Println("Run", i+1, "time:", times[i])
		} else {
			fmt.Println("Not sorted, Algorithm not functional")
			os.Exit(1) // terminate test if algo fails
		}
	}

	totalTime := time.Duration(0)
	for _, v := range times {
		totalTime += v
	}
	fmt.Println("Average time:", totalTime/numRuns)
}

// Generates a slice of size, size filled with random positive 64bit numbers
func generateSlice(size int) []uint64 {
	slice := make([]uint64, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Uint64()
	}
	return slice
}

func mySort(inArray []uint64) {
	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(inArray)-1 {
			if inArray[i] > inArray[i+1] {
				inArray[i], inArray[i+1] = inArray[i+1], inArray[i]
				isDone = false
			}
			i++
		}
	}
}
