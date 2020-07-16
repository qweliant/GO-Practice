package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


// json url
const citiBike = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// data structure to unmarshal JSON
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL int `json:"ttl"`
	Data struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}


func main() {
	
	// get res from url
	response, err := http.Get(citiBike)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// read res
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Declare a variable of type staionData
	var sd stationData

	// unmarshall from json to variable
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n", sd.Data.Stations[0])

	// Marshall data
	outputData, err := json.Marshal(sd)
	if err != nil {
		log.Fatal(err)
	}
	
	// save data
	if err := ioutil.WriteFile("data/citibike.json", outputData, 0644); err != nil {
		log.Fatal(err)
	}



}
