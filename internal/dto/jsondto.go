// Package: dto
// File: jsondto.go
// Description : This file contains the json dto for the data received from the api
package dto

// VehicleType data structure
type JVehicleType struct {
	Name      string `json:"Name"`
	IsPrimary bool   `json:"IsPrimary"`
}

// Manufacturer data structure
type JManufacter struct {
	Id         uint           `json:"Mfr_ID"`
	Name       string         `json:"Mfr_Name"`
	CommonName string         `json:"Mfr_CommonName"`
	Country    string         `json:"Country"`
	Vehicles   []JVehicleType `json:"VehicleTypes"`
}

// Main data structure
type JsonData struct {
	Count         int           `json:"Count"`
	Message       string        `json:"Message"`
	Manufacturers []JManufacter `json:"Results"`
}
