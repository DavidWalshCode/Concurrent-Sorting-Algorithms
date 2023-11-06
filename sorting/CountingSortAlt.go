package sorting

import "sync"

// Concurrent counting sort alt algorithm
func CountingSortAlt(input []int, maxVal int) []int {
	// Find the range of input values by identifying the smallest value
	minVal := maxVal
	for _, num := range input {
		if num < minVal {
			minVal = num
		}
	}

	// Initialize the slices for counting and the result
	// The offset is used to adjust negative indices since slices cannot have negative indices.
	offset := -minVal
	// Size is the number of possible values in the input
	size := maxVal - minVal + 1
	// Count slice keeps track of the number of times each value appears
	count := make([]int, size)
	// Result slice will contain the sorted elements
	result := make([]int, len(input))

	// Concurrently count the occurrences of each number
	var wait sync.WaitGroup
	for _, num := range input {
		// Increment the WaitGroup counter before starting a goroutine
		wait.Add(1)
		// Start a new goroutine for each number in the input
		go func(n int) {
			defer wait.Done() // Decrement the WaitGroup counter when the goroutine completes
			count[n+offset]++ // Increment the count for the number
		}(num)
	}
	// Wait for all goroutines to finish
	wait.Wait()

	// Sum up the counts to get the starting index for each number
	for i := 1; i < size; i++ {
		count[i] += count[i-1]
	}

	// Build the output array using the counts to determine the positions
	for _, num := range input {
		// Place the number in the result slice based on the count, then decrement the count
		result[count[num+offset]-1] = num
		count[num+offset]--
	}

	// Return the sorted result
	return result
}
