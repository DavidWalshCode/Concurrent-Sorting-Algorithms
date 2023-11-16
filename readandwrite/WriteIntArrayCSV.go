package readwrite

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func WriteIntArrayCSV() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a file named in.csv
	file, err := os.Create("input/inbig.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Generate and write 10,000 random integers
	var record []string
	for i := 0; i < 100000; i++ {
		// Generate a random integer between -99999 and 99999
		num := rand.Intn(199999) - 99999
		record = append(record, strconv.Itoa(num))
	}

	// Write the record to the CSV file
	if err := writer.Write(record); err != nil {
		panic(err)
	}
}
