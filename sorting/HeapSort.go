package sorting

import "sync"

// sink performs the 'sink' operation of the heapsort algorithm, helping to maintain the heap property.
func sink(arr []int, start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	root := start

	for {
		child := root*2 + 1
		if child > end {
			break
		}
		if child+1 <= end && arr[child] < arr[child+1] {
			child++
		}
		if arr[root] < arr[child] {
			arr[root], arr[child] = arr[child], arr[root]
			root = child
		} else {
			break
		}
	}
}

// heapify builds a max heap from a slice of integers.
func heapify(arr []int, wg *sync.WaitGroup) {
	n := len(arr)
	for start := (n - 2) / 2; start >= 0; start-- {
		wg.Add(1)
		go sink(arr, start, n-1, wg)
	}
	wg.Wait()
}

// heapsort concurrently sorts an array using the heapsort algorithm.
func HeapSort(arr []int) []int {
	wg := sync.WaitGroup{}
	heapify(arr, &wg)

	end := len(arr) - 1
	for end > 0 {
		arr[end], arr[0] = arr[0], arr[end]
		end--
		wg.Add(1)
		go sink(arr, 0, end, &wg)
		wg.Wait()
	}
	return arr
}
