// Package: main
// File: main.go
// Description : This file contains the main function to run the application
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	json "github.com/rcsolis/encodings_sample/internal/jsonparser"
	"github.com/rcsolis/encodings_sample/internal/utils/database"
	xml "github.com/rcsolis/encodings_sample/internal/xmlparser"
)

// Function to load environment variables from .env file
// Returns the base url
func loadEnvFile() string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("BASE_URL")
}

// Main function
func main() {
	// Loading .env file
	base_url := loadEnvFile()
	// Parsing XML
	xml.Parse(base_url)
	// Parsing JSON
	data := json.Parse(base_url)

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close(db)
	// Insert into database
	for _, m := range data.Manufacturers {
		vechicles := []database.VehicleType{}
		for _, v := range m.Vehicles {
			vechicles = append(vechicles, database.VehicleType{Name: v.Name, IsPrimary: v.IsPrimary})
		}
		manufacturer := database.Manufacturer{
			Key:          m.Id,
			Name:         m.Name,
			CommonName:   &m.CommonName,
			Country:      m.Country,
			VehicleTypes: vechicles,
		}
		if err := database.Insert(manufacturer, db); err != nil {
			log.Fatal(err)
		}
	}

	// Search for a manufacturer
	var search_criteria string
	fmt.Println("Enter the search criteria: ")
	fmt.Scan(&search_criteria)
	manufacturers, err := database.SearchManufacturer(search_criteria, db)
	if err != nil {
		log.Fatalln("Error searching for manufacturer")
		log.Fatal(err)
	}
	// Print search results
	log.Println("Search results:")
	for _, m := range manufacturers {
		fmt.Println("Manufacturer: ", m.Name, "Country: ", m.Country)
		if len(m.VehicleTypes) == 0 {
			fmt.Println("> No vehicles found")
			continue
		}
		for _, v := range m.VehicleTypes {
			fmt.Println("Vehicle: ", v.Name)
		}
	}
}
