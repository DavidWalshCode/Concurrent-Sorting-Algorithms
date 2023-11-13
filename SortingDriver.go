package main

import (
	"Concurrent-Sort-Algorithms/readandwrite"
	"Concurrent-Sort-Algorithms/sorting"
	"fmt"
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
	//sortedData := sorting.MergeSortAlt(data) // Sort the data using the concurrent merge sort algorithm
	//sortedData := sorting.CountingSort(data, len(data)) // Sort the data using the concurrent counting sort algorithm
	sortedData := sorting.CountingSortAlt(data, len(data)) // Sort the data using the concurrent counting sort algorithm
	//sortedData := sorting.HeapSort(data) // Sort the data using the concurrent heap sort algorithm
	//sortedData := sorting.ShellSort(data) // Sort the data using the concurrent shell sort algorithm

	// Calculate the elapsed time since the timer started
	elapsed := time.Since(startTime)
	fmt.Printf("Sorted %d numbers in:\n", len(data))
	println(" -", elapsed.Seconds(), "seconds")
	println(" -", elapsed.Milliseconds(), "milliseconds")
	println(" -", elapsed.Microseconds(), "microseconds")
	println(" -", elapsed.Nanoseconds(), "nanoseconds")

	// Write the sorted data to out.csv
	if err := readwrite.WriteCSV("output/outCountingSortAlt.csv", sortedData); err != nil {
		panic(err)
	}
}
