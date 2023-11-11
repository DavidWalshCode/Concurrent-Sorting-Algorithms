package sorting

import "sync"

// Concurrent merge sort algorithm
func MergeSortAlt(data []int) []int {
	return mergeSortConcurrent(data, 0)
}

// Helper function that manages concurrency based on slice size
func mergeSortConcurrent(data []int, depth int) []int {
	if len(data) <= 1 {
		return data
	}

	// Use a threshold to limit concurrency
	if len(data) < 2048 || depth > 4 { // Adjust the threshold and depth limit as needed
		return mergeSortSequential(data)
	}

	mid := len(data) / 2
	var left, right []int
	var wait sync.WaitGroup
	wait.Add(2)

	// Concurrently sort the left half
	go func() {
		left = mergeSortConcurrent(data[:mid], depth+1)
		wait.Done()
	}()
	// Concurrently sort the right half
	go func() {
		right = mergeSortConcurrent(data[mid:], depth+1)
		wait.Done()
	}()

	wait.Wait()
	return mergeAlt(left, right)
}

// Sequential merge sort for smaller slices
func mergeSortSequential(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := mergeSortSequential(data[:mid])
	right := mergeSortSequential(data[mid:])
	return mergeAlt(left, right)
}

// Merges two sorted slices
func mergeAlt(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
