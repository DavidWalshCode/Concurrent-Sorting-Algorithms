package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"
	"time"
)

// ReadCSV reads integers from a CSV file.
func ReadCSV(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []int
	for _, record := range records {
		for _, value := range record {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			data = append(data, num)
		}
	}

	return data, nil
}

// WriteCSV writes integers to a CSV file.
func WriteCSV(filename string, data []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		record := []string{strconv.Itoa(value)}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// mergeSort is a concurrent implementation of the merge sort algorithm.
func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	var wait sync.WaitGroup
	wait.Add(2)

	var left, right []int
	go func() {
		left = mergeSort(data[:mid])
		wait.Done()
	}()
	go func() {
		right = mergeSort(data[mid:])
		wait.Done()
	}()
	wait.Wait()

	return merge(left, right)
}

// merge merges two sorted slices.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	// Read data from in.csv
	data, err := ReadCSV("in1.csv")
	if err != nil {
		panic(err)
	}

	// Start the timer
	startTime := time.Now()

	// Sort data
	sortedData := mergeSort(data)

	// Stop the timer and print the execution time
	elapsed := time.Since(startTime)
	println("Execution time:", elapsed.Seconds(), "seconds")
	println("Execution time:", elapsed.Milliseconds(), "milliseconds")
	println("Execution time:", elapsed.Microseconds(), "microseconds")
	println("Execution time:", elapsed.Nanoseconds(), "nanoseconds")

	// Write sorted data to out.csv
	if err := WriteCSV("out.csv", sortedData); err != nil {
		panic(err)
	}
}
