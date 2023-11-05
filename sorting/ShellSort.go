package sorting

import "sync"

// Concurrent shell sort function
func ShellSort(arr []int) []int {
	n := len(arr)
	wg := sync.WaitGroup{}

	// Start with a big gap, then reduce the gap
	for gap := n / 2; gap > 0; gap /= 2 {
		// Do a gapped insertion sort for this gap size.
		for i := gap; i < n; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				temp := arr[i]
				j := i
				for ; j >= gap && arr[j-gap] > temp; j -= gap {
					arr[j] = arr[j-gap]
				}
				arr[j] = temp
			}(i)
		}
		wg.Wait() // Wait for all goroutines to finish
	}
	return arr
}
