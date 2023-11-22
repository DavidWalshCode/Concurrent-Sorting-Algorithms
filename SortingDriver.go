package main

import (
	"Concurrent-Sort-Algorithms/readandwrite"
	"Concurrent-Sort-Algorithms/sorting"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Read data from .csv
	data, err := readwrite.ReadCSV("input/numbers.csv")
	if err != nil {
		panic(err)
	}

	coreCount := runtime.NumCPU() // 4 Cores for my laptop
	//coreCount := 3
	//coreCount := 2
	//coreCount := 1
	runtime.GOMAXPROCS(coreCount)

	startTime := time.Now() // Start the timer to measure execution time

	sortedData := sorting.MergeSort(data) // Sort the data using the concurrent merge sort algorithm
	//sortedData := sorting.CountingSort(data) // Sort the data using the concurrent counting sort algorithm
	//sortedData := sorting.HeapSort(data) // Sort the data using the concurrent heap sort algorithm
	//sortedData := sorting.ShellSort(data) // Sort the data using the concurrent shell sort algorithm

	// Calculate the elapsed time since the timer started
	elapsed := time.Since(startTime)

	// Printing
	fmt.Printf("Running with %d core(s)\n", coreCount)
	fmt.Printf("Sorted %d numbers in %s\n", len(sortedData), elapsed)
	println(" -", elapsed.Seconds(), "seconds")
	println(" -", elapsed.Milliseconds(), "milliseconds")
	println(" -", elapsed.Microseconds(), "microseconds")
	println(" -", elapsed.Nanoseconds(), "nanoseconds")

	// Write the sorted data to .csv
	if err := readwrite.WriteCSV("output/out(20276885).csv", sortedData); err != nil {
		panic(err)
	}

	// Check if the numbers are sorted or not
	if isSorted(sortedData) {
		fmt.Println("The numbers are sorted correctly")
	} else {
		fmt.Println("The numbers are not sorted correctly")
	}
}

// Check is the numbers are sorted correctly
func isSorted(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return false
		}
	}
	return true
}
