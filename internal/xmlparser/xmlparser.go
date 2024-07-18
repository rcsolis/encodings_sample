// Package: xmlparser
// File: xmlparser.go
// Description : This file contains the xml parser for the data received from the api
package xmlparser

import (
	"encoding/xml"
	"log"

	models "github.com/rcsolis/encodings_sample/internal/dto"
	"github.com/rcsolis/encodings_sample/internal/utils"
)

// Variable to store the parsed data
var xmlData models.XmlData

// Function to print the results
func printResults(data models.XmlData) {
	// Print results
	log.Println("--- XML DATA PARSED ---")
	log.Printf("Results Count: %d", data.Count)
	log.Printf("Message: %s", data.Message)
	log.Println("Manufacturers:")
	for _, m := range data.Results.Manufacturers {
		log.Printf(">ID: %d - Name: %s - CommonName: %s - Country: %s \n", m.Id, m.Name, m.CommonName, m.Country)
	}
	log.Println("-----------------------")
}

// Function to parse the xml data
func Parse(dataSource string) {
	var rawData []byte
	// Make request
	rawData, err := utils.MakeRequest(dataSource, "xml")
	if err != nil {
		log.Fatal(err)
	}
	// Parse XML
	if err := xml.Unmarshal(rawData, &xmlData); err != nil {
		log.Fatal(err)
	}
	// Print results
	printResults(xmlData)
}
