package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"
	"time"
)

// Reads integers from a CSV file
func ReadCSV(filename string) ([]int, error) {
	// Open the CSV file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close() // Ensure the file is closed when the function returns

	reader := csv.NewReader(file) // Create a new CSV reader

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []int // Initialize a slice to hold the integers

	// Loop through each record (row) in the CSV
	for _, record := range records {
		// Loop through each field (column) in the record
		for _, value := range record {
			// Convert the string field to an integer
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			data = append(data, num) // Append the integer to the data slice
		}
	}

	return data, nil
}

// Writes integers to a CSV file
func WriteCSV(filename string, data []int) error {
	// Create the CSV file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close() // Ensure the file is closed when the function returns

	writer := csv.NewWriter(file) // Create a new CSV writer
	defer writer.Flush()          // Ensure any buffered data is written when the function returns

	// Loop through each integer in the data slice
	for _, value := range data {
		record := []string{strconv.Itoa(value)} // Convert the integer to a string and create a record
		// Write the record to the CSV
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// Concurrent implementation of the merge sort algorithm
func mergeSort(data []int) []int {
	// If the data slice is empty or has one element, it's already sorted
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2    // Find the middle index
	var wait sync.WaitGroup // Declare a wait group to synchronize goroutines
	wait.Add(2)             // Add two counts to the wait group for the two goroutines we will launch

	var left, right []int // Declare slices to hold the left and right halves

	// Concurrently sort the left half
	go func() {
		left = mergeSort(data[:mid])
		wait.Done()
	}()
	// Concurrently sort the right half
	go func() {
		right = mergeSort(data[mid:])
		wait.Done()
	}()

	wait.Wait() // Wait for both halves to be sorted

	return merge(left, right) // Merge the sorted halves and return the result
}

// Merges the two sorted slices together
func merge(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right)) // Create a slice to hold the merged result
	i, j := 0, 0                                   // Initialize indices for iterating over the left and right slices

	// Merge the slices until one is exhausted
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from left and right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	// Read data from in.csv
	data, err := ReadCSV("inbig.csv")
	if err != nil {
		panic(err)
	}

	startTime := time.Now() // Start the timer to measure execution time

	sortedData := mergeSort(data) // Sort the data using the concurrent merge sort algorithm

	// Calculate the elapsed time since the timer started
	elapsed := time.Since(startTime)
	println("Execution time:", elapsed.Seconds(), "seconds")
	println("Execution time:", elapsed.Milliseconds(), "milliseconds")
	println("Execution time:", elapsed.Microseconds(), "microseconds")
	println("Execution time:", elapsed.Nanoseconds(), "nanoseconds")

	// Write the sorted data to out.csv
	if err := WriteCSV("out.csv", sortedData); err != nil {
		panic(err)
	}
}
