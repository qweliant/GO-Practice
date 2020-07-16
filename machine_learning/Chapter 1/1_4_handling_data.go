package main

import (
	"fmt"
	// "log"
	"os"
	"github.com/kniren/gota/dataframe"
)

func main(){
	
	// read file & check for error, close file stream
	shootings, err := os.Open("data/shooting.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer shootings.Close()
	
	// set as dataframe
	df := dataframe.ReadCSV(shootings)

	fmt.Println(df)
}
