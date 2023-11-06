package sorting

import "sync"

// Concurrent counting sort algorithm
func CountingSort(data []int, maxVal int) []int {
	// Find the range of data values
	minVal := maxVal
	for _, num := range data {
		if num < minVal {
			minVal = num
		}
	}

	// Initialize the slices for counting and the result
	offset := -minVal
	size := maxVal - minVal + 1
	count := make([]int, size)
	result := make([]int, len(data))

	// Concurrently count the occurrences of each number
	var wait sync.WaitGroup
	for _, num := range data {
		wait.Add(1)
		go func(n int) {
			defer wait.Done()
			count[n+offset]++
		}(num)
	}
	wait.Wait()

	// Sum up the counts
	for i := 1; i < size; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for _, num := range data {
		result[count[num+offset]-1] = num
		count[num+offset]--
	}

	return result
}
