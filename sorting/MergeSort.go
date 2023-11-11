package sorting

import "sync"

// Concurrent merge sort algorithm
func MergeSort(data []int) []int {
	// If the data slice is empty or has one element, it's already sorted
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2     // Find the middle index
	var waitG sync.WaitGroup // Declare a wait group to synchronize goroutines
	waitG.Add(2)             // Add two counts to the wait group for the two goroutines we will launch

	var left, right []int // Declare slices to hold the left and right halves

	// Concurrently sort the left half
	go func() {
		left = MergeSort(data[:mid])
		waitG.Done()
	}()
	// Concurrently sort the right half
	go func() {
		right = MergeSort(data[mid:])
		waitG.Done()
	}()

	waitG.Wait() // Wait for both halves to be sorted

	return merge(left, right) // Merge the sorted halves and return the result
}

// Merges the two sorted slices together
func merge(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right)) // Create a slice to hold the merged result
	i, j := 0, 0                                   // Initialize indices for iterating over the left and right slices

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
