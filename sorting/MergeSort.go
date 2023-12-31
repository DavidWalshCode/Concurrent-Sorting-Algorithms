package sorting

import "sync"

func MergeSort(data []int) []int {
	return mergeSortConcurrent(data, 0)
}

// Concurrent merge sort that also manages concurrency based on slice size
func mergeSortConcurrent(data []int, depth int) []int {
	// If a slice of length 1 or less is already sorted
	if len(data) <= 1 {
		return data
	}

	// Using a threshold to limit concurrency
	if len(data) < 4000 || depth > 5 { // Adjust the threshold and depth limit as needed
		return mergeSortSequential(data)
	}

	mid := len(data) / 2
	var left, right []int

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	// Concurrently sort the left half
	go func() {
		left = mergeSortConcurrent(data[:mid], depth+1)
		waitGroup.Done()
	}()
	// Concurrently sort the right half
	go func() {
		right = mergeSortConcurrent(data[mid:], depth+1)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	return merge(left, right)
}

// Sequential merge sort for smaller slices
func mergeSortSequential(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := mergeSortSequential(data[:mid])
	right := mergeSortSequential(data[mid:])

	return merge(left, right)
}

// Merges two sorted slices
func merge(left, right []int) []int {
	sortedResult := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sortedResult = append(sortedResult, left[i])
			i++
		} else {
			sortedResult = append(sortedResult, right[j])
			j++
		}
	}

	sortedResult = append(sortedResult, left[i:]...)
	sortedResult = append(sortedResult, right[j:]...)

	return sortedResult
}
