// Package: jsonparser
// File: jsonparser.go
// Description : This file contains the json parser for the data received from the api
package jsonparser

import (
	"encoding/json"
	"log"

	models "github.com/rcsolis/encodings_sample/internal/dto"
	"github.com/rcsolis/encodings_sample/internal/utils"
)

// Variable to store the parsed data
var jsonData models.JsonData

// Function to print the results
func printResults(data models.JsonData) {
	// Print results
	log.Println("--- JSON DATA PARSED ---")
	log.Printf("Results Count: %d", data.Count)
	log.Printf("Message: %s", data.Message)
	log.Println("Manufacturers:")
	for _, m := range data.Manufacturers {
		log.Printf(">ID: %d - Name: %s - CommonName: %s - Country: %s \n", m.Id, m.Name, m.CommonName, m.Country)
	}
	log.Println("-----------------------")
}

// Main function to parse the json data
func Parse(dataSource string) models.JsonData {
	var rawData []byte
	// Make request
	rawData, err := utils.MakeRequest(dataSource, "json")
	if err != nil {
		log.Fatal(err)
	}
	// Parse XML
	if err := json.Unmarshal(rawData, &jsonData); err != nil {
		log.Fatal(err)
	}
	// Print results
	printResults(jsonData)
	return jsonData
}
