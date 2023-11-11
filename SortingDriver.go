package main

import (
	"Concurrent-Sort-Algorithms/readandwrite"
	"Concurrent-Sort-Algorithms/sorting"
	"time"
)

func main() {
	// Read data from in.csv
	data, err := readwrite.ReadCSV("input/in.csv")
	if err != nil {
		panic(err)
	}

	startTime := time.Now() // Start the timer to measure execution time

	//sortedData := sorting.MergeSort(data) // Sort the data using the concurrent merge sort algorithm
	sortedData := sorting.MergeSortAlt(data) // Sort the data using the concurrent merge sort algorithm
	//sortedData := sorting.CountingSort(data, len(data)) // Sort the data using the concurrent counting sort algorithm
	//sortedData := sorting.CountingSortAlt(data, len(data)) // Sort the data using the concurrent counting sort algorithm
	//sortedData := sorting.HeapSort(data) // Sort the data using the concurrent heap sort algorithm
	//sortedData := sorting.ShellSort(data) // Sort the data using the concurrent shell sort algorithm

	// Calculate the elapsed time since the timer started
	elapsed := time.Since(startTime)
	println("Execution time:", elapsed.Seconds(), "seconds")
	println("Execution time:", elapsed.Milliseconds(), "milliseconds")
	println("Execution time:", elapsed.Microseconds(), "microseconds")
	println("Execution time:", elapsed.Nanoseconds(), "nanoseconds")

	// Write the sorted data to out.csv
	if err := readwrite.WriteCSV("output/outMergeSortAlt.csv", sortedData); err != nil {
		panic(err)
	}
}
