package main

import (
	"encoding/csv"
	"io"

	"fmt"
	"log"
	"os"
	// "strconv"
)

func main() {

	// open csv
	f, err := os.Open("data/shooting.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read in csv
	df := csv.NewReader(f)

	// FieldsPerRecord is set negative in case each row has a variable number of fields
	df.FieldsPerRecord = -1

	// rawCSVData will hold our
	// reading df in chunk, cool but doesnt handle typing
	// records, err := df.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var raw [][]string

	for {

		record, err := df.Read()
		if err == io.EOF {
			break
		}
		raw = append(raw, record)

	}
}
