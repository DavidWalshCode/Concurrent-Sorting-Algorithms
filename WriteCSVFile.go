package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a file named in.csv
	file, err := os.Create("inbig.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Generate and write 10,000 random integers
	var record []string
	for i := 0; i < 10000000; i++ {
		// Generate a random integer between -9999 and 9999
		num := rand.Intn(19999) - 9999
		record = append(record, strconv.Itoa(num))
	}

	// Write the record to the CSV file
	if err := writer.Write(record); err != nil {
		panic(err)
	}
}
