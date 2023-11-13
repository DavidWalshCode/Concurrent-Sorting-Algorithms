package main

// Importing all the necessary packages needed.
import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Read numbers from the input CSV file.
	numbers, err := readNumbers("numbers.csv")
	if err != nil {
		log.Fatalf("Error reading numbers: %v", err)
	}

	// Record the start time.
	start := time.Now()

	// This is where your sort method on the numbers array goes.
	// Sort numbers here.

	// Calculate the elapsed time.
	elapsed := time.Since(start)

	// Write the sorted numbers to the output CSV file. Change the studentID to yours.
	err = writeNumbers("out(20276885).csv", numbers)
	if err != nil {
		log.Fatalf("Error writing numbers: %v", err)
	}

	// Print the number of sorted elements and the time taken.
	fmt.Printf("Sorted %d numbers in %s.\n", len(numbers), elapsed)
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
