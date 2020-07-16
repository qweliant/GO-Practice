package main

import (
	"encoding/csv"
	"io"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CSVRecord struct {
	id int32
	name string
	date string
	manner_of_death string
	armed string
	age int16
	gender string
	race string
	city string
	state string
	signs_of_mental_illness bool
	threat_level string
	flee string
	body_camera bool
	ParseError error
}


func main() {

	// open csv
	f, err := os.Open("data/shooting.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read in csv
	df := csv.NewReader(f)
	var csvData = []CSVRecord
	// FieldsPerRecord is set negative in case each row has a variable number of fields
	df.FieldsPerRecord = 14

	// rawCSVData will hold our
	// reading df in chunk, cool but doesnt handle typing
	// records, err := df.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	
	var raw [][]string

	// iterate through each row to check types
	for {

		record, err := df.Read()
		if err == io.EOF {
			break
		}

		var csvRecord CSVRecord
		const check =  0 
		for idx, value:= range record{

			// check types
			if idx == 0 {
				if value < check {
					log.Printf("Unexpected type in column %d\n", idx)
					csvRecord.ParseError = fmt.Errorf("!!Empty string value!!")
					break
				}
				
				csvRecord.id = value
				continue
			}

			if idx == 5 {
				if value == check {
					log.Printf("Unexpected type in column %d\n", idx)
					csvRecord.ParseError = fmt.Errorf("!!Empty string value!!")
					break
				}
				
				csvRecord.armed = value
				continue
			}

			if idx == 10 {
				if value == "" {
					log.Printf("Unexpected type in column %d\n", idx)
					csvRecord.ParseError = fmt.Errorf("!!Empty string value!!")
					break
				}
				if value == "True"{
					csvRecord.signs_of_mental_illness = true
					continue
				}
				if value == "True"{
					csvRecord.signs_of_mental_illness = true
					continue
				}
				
			}

			var colValue string
			switch idx {
			case 0:
				csvRecord.id
			}
		}
		if err != nil {
			log.Println(err)
			continue
		}

		// # append record to dataset
		raw = append(raw, record)

	}
}
