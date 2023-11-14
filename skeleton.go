package main

// Importing all the necessary packages needed.
import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Read numbers from the input CSV file.
	numbers, err := readNumbers("numbers.csv")
	if err != nil {
		log.Fatalf("Error reading numbers: %v", err)
	}

	coreCount := runtime.NumCPU()
	runtime.GOMAXPROCS(coreCount)

	// Record the start time.
	start := time.Now()

	// This is where your sort method on the numbers array goes.
	// Sort numbers here.
	sortedNumbers := countingSort(numbers)

	// Calculate the elapsed time.
	elapsed := time.Since(start)

	// Write the sorted numbers to the output CSV file. Change the studentID to yours.
	err = writeNumbers("out(20276885).csv", sortedNumbers)
	if err != nil {
		log.Fatalf("Error writing numbers: %v", err)
	}

	// Print the number of sorted elements and the time taken.
	fmt.Printf("Sorted %d numbers in %s.\n", len(sortedNumbers), elapsed)
}

// readNumbers reads integers from a CSV file and returns them as a slice.
func readNumbers(filename string) ([]int, error) {
	// Open the input file.
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader.
	reader := csv.NewReader(file)

	// Read the numbers from the CSV file and store them in a slice.
	numbers := []int{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for _, value := range record {
			number, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, number)
		}
	}

	return numbers, nil
}

// writeNumbers writes a slice of integers to a CSV file.
func writeNumbers(filename string, numbers []int) error {
	// Create the output file.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the numbers to the CSV file.
	for _, number := range numbers {
		err := writer.Write([]string{strconv.Itoa(number)})
		if err != nil {
			return err
		}
	}

	return nil
}

func countingSort(data []int) []int {
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
	numCPU := runtime.NumCPU()
	chunkSize := (len(data) + numCPU - 1) / numCPU

	// Create a slice of slices to hold local counts
	localCounts := make([][]int, numCPU)
	for i := range localCounts {
		localCounts[i] = make([]int, size)
	}

	var wait sync.WaitGroup

	// Process chunks concurrently
	for i := 0; i < numCPU; i++ {
		wait.Add(1)
		go func(chunkStart int) {
			defer wait.Done()
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

	wait.Wait()

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
