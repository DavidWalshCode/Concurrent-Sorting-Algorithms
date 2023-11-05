package readwrite

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Reads an integer array from a CSV file
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

// Writes the integer array to a CSV file
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
