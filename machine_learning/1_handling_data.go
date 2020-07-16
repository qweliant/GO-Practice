package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	// "strconv"
)

type CSVRecord struct {
	id                      string
	name                    string
	date                    string
	manner_of_death         string
	armed                   string
	age                     string
	gender                  string
	race                    string
	city                    string
	state                   string
	signs_of_mental_illness bool
	threat_level            string
	flee                    string
	body_camera             bool
	ParseError              error
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
	var csvData []CSVRecord
	// FieldsPerRecord is set negative in case each row has a variable number of fields
	df.FieldsPerRecord = 14

	// rawCSVData will hold our
	// reading df in chunk, cool but doesnt handle typing
	// records, err := df.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// iterate through each row to check types
	for {

		record, err := df.Read()
		if err == io.EOF {
			break
		}

		var csvRecord CSVRecord
		const check = 0
		for idx, value := range record {
			fmt.Printf("VALUE: %d", value)

			// // check types
			// if idx == 0 {
			// 	if value < check {
			// 		log.Printf("Unexpected type in column %d\n", idx)
			// 		csvRecord.ParseError = fmt.Errorf("!!Empty string value!!")
			// 		break
			// 	}

			// 	csvRecord.id = value
			// 	continue
			// }

			// if idx == 5 {
			// 	if value == check {
			// 		log.Printf("Unexpected type in column %d\n", idx)
			// 		csvRecord.ParseError = fmt.Errorf("!!Empty string value!!")
			// 		break
			// 	}

			// 	csvRecord.armed = value
			// 	continue
			// }

			if idx == 10 {
				if value == "" {
					log.Printf("Unexpected type in column %d\n", idx)
					csvRecord.ParseError = fmt.Errorf("!!Empty string value!!")
					break
				}
				if value == "True" {
					csvRecord.signs_of_mental_illness = true
					continue
				}
				if value == "False" {
					csvRecord.signs_of_mental_illness = false
					continue
				}

			}

			if idx == 13 {
				if value == "" {
					log.Printf("Unexpected type in column %d\n", idx)
					csvRecord.ParseError = fmt.Errorf("!!Empty string value!!")
					break
				}
				if value == "True" {
					csvRecord.body_camera = true
					continue
				}
				if value == "False" {
					csvRecord.body_camera = false
					continue
				}
			}

			if value == "" {
				log.Printf("Unexpected type in column %d\n", idx)
				csvRecord.ParseError = fmt.Errorf("!!Empty string value!!")
				break
			}

			var colValue string

			switch idx {
			case 0:
				csvRecord.name = colValue
			case 1:
				csvRecord.date = colValue
			case 2:
				csvRecord.manner_of_death = colValue
			case 3:
				csvRecord.armed = colValue
			case 4:
				csvRecord.gender = colValue
			case 5:
				csvRecord.race = colValue
			case 6:
				csvRecord.city = colValue
			case 7:
				csvRecord.state = colValue
			case 8:
				csvRecord.threat_level = colValue
			case 9:
				csvRecord.flee = colValue
			case 10:
				csvRecord.id = colValue
			case 11:
				csvRecord.age = colValue
			}
		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

		// # append record to dataset

	}
}
