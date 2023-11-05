package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sync"
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
		myMergeSort(slice)

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

// Concurrent implementation of the merge sort algorithm
func myMergeSort(data []uint64) []uint64 {
	// If the data slice is empty or has one element, it's already sorted
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2    // Find the middle index
	var wait sync.WaitGroup // Declare a wait group to synchronize goroutines
	wait.Add(2)             // Add two counts to the wait group for the two goroutines we will launch

	var left, right []uint64 // Declare slices to hold the left and right halves

	// Concurrently sort the left half
	go func() {
		left = myMergeSort(data[:mid])
		wait.Done()
	}()
	// Concurrently sort the right half
	go func() {
		right = myMergeSort(data[mid:])
		wait.Done()
	}()

	wait.Wait() // Wait for both halves to be sorted

	return myMerge(left, right) // Merge the sorted halves and return the result
}

// Merges the two sorted slices together
func myMerge(left, right []uint64) []uint64 {

	result := make([]uint64, 0, len(left)+len(right)) // Create a slice to hold the merged result
	i, j := 0, 0                                      // Initialize indices for iterating over the left and right slices

	// Merge the slices until one is exhausted
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from left and right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
