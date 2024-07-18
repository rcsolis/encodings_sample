// Package: dto
// File: xmldto.go
// Description : This file contains the xml dto for the data received from the api
package dto

import "encoding/xml"

// VehicleType data structure
type XVehicleType struct {
	VehicleType xml.Name `xml:"VehicleType"`
	Name        string   `xml:"Name"`
	IsPrimary   bool     `xml:"IsPrimary"`
}

// Vehicles data structure
type XVehicles struct {
	XMLName  xml.Name       `xml:"VehicleTypes"`
	Vehicles []XVehicleType `xml:"VehicleType"`
}

// Manufacturer data structure
type XManufacter struct {
	XMLName    xml.Name  `xml:"Manufacturers"`
	Id         uint      `xml:"Mfr_ID"`
	Name       string    `xml:"Mfr_Name"`
	CommonName string    `xml:"Mfr_CommonName"`
	Country    string    `xml:"Country"`
	Vehicles   XVehicles `xml:"VehicleTypes"`
}

// Results data structure
type XResults struct {
	XMLName       xml.Name      `xml:"Results"`
	Manufacturers []XManufacter `xml:"Manufacturers"`
}

// Main data structure
type XmlData struct {
	XMLName xml.Name `xml:"Response"`
	Count   int      `xml:"Count"`
	Message string   `xml:"Message"`
	Results XResults `xml:"Results"`
}
