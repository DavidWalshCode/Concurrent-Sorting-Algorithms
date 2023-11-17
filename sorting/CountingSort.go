package sorting

import (
	"runtime"
	"sync"
)

// David Walsh 20276885
func CountingSort(data []int) []int {
	if len(data) == 0 {
		return data
	}

	// Find the range of data values
	maxVal := data[0]
	minVal := data[0]
	for _, num := range data {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	offset := -minVal
	size := maxVal - minVal + 1

	// Use the number of CPU cores to limit concurrency
	numCPU := runtime.NumCPU()
	chunkSize := (len(data) + numCPU - 1) / numCPU

	// Create a slice of slices to hold local counts
	localCounts := make([][]int, numCPU)
	for i := range localCounts {
		localCounts[i] = make([]int, size)
	}

	var waitGroup sync.WaitGroup

	// Process chunks concurrently
	for i := 0; i < numCPU; i++ { // Loop iterates for each available CPU Core (4 for me).  The idea is to create a separate goroutine for each chunk of the data to be sorted, allowing these chunks to be processed in parallel
		waitGroup.Add(1)
		go func(chunkStart int) { // chunkStart represents the starting index of the data chunk that this particular goroutine will process
			defer waitGroup.Done()
			chunkEnd := chunkStart + chunkSize // Each goroutine calculates the end index of its chunk (chunkEnd). If the calculated end exceeds the length of the data, it's adjusted to be the length of the data to avoid out-of-bounds access
			if chunkEnd > len(data) {
				chunkEnd = len(data)
			}
			localCount := localCounts[chunkStart/chunkSize] // Each goroutine accesses its own local count array from localCounts. This local array is used to count occurrences of each element within the chunk
			for _, num := range data[chunkStart:chunkEnd] {
				localCount[num+offset]++ // The goroutine iterates over its assigned chunk of the data (data[chunkStart:chunkEnd]). For each element num in the chunk, it increments the corresponding index in the local count array. The offset is used to adjust the index for cases where the data contains negative numbers
			}
		}(i * chunkSize) // This ensures that each goroutine gets a different chunk of the data to process.
	}

	waitGroup.Wait()

	// Merge local counts
	globalCount := make([]int, size)
	for _, localCount := range localCounts {
		for i := 0; i < size; i++ {
			globalCount[i] += localCount[i]
		}
	}

	// Sum up the counts
	for i := 1; i < size; i++ {
		globalCount[i] += globalCount[i-1]
	}

	sortedResult := make([]int, len(data))

	// Build the output array
	for _, num := range data {
		sortedResult[globalCount[num+offset]-1] = num
		globalCount[num+offset]--
	}

	return sortedResult
}
