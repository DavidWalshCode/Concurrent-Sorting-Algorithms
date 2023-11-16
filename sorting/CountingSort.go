package sorting

import (
	"runtime"
	"sync"
)

func CountingSort(data []int) []int {
	// Find the range of data values
	maxVal := len(data)
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
	result := make([]int, len(data))

	// Use the number of CPU cores to limit concurrency
	numCPU := runtime.NumCPU() // 1
	chunkSize := (len(data) + numCPU - 1) / numCPU

	// Create a slice of slices to hold local counts
	localCounts := make([][]int, numCPU)
	for i := range localCounts {
		localCounts[i] = make([]int, size)
	}

	var waitGroup sync.WaitGroup

	// Process chunks concurrently
	for i := 0; i < numCPU; i++ {
		waitGroup.Add(1)
		go func(chunkStart int) {
			defer waitGroup.Done()
			chunkEnd := chunkStart + chunkSize
			if chunkEnd > len(data) {
				chunkEnd = len(data)
			}
			localCount := localCounts[chunkStart/chunkSize]
			for _, num := range data[chunkStart:chunkEnd] {
				localCount[num+offset]++
			}
		}(i * chunkSize)
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

	// Build the output array
	for _, num := range data {
		result[globalCount[num+offset]-1] = num
		globalCount[num+offset]--
	}

	return result
}
