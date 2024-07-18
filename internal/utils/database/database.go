// Package: database
// File: database.go
// Description : This file contains the database operations for the data received from the api
package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Data model for Manufacturer table
// Has a one-to-many relationship with VehicleType
type Manufacturer struct {
	gorm.Model
	Key          uint
	Name         string
	CommonName   *string
	Country      string
	VehicleTypes []VehicleType
}

// Data model for VehicleType table
// Has a belongs-to relationship with Manufacturer
type VehicleType struct {
	gorm.Model
	Name           string
	IsPrimary      bool `gorm:"default:false"`
	ManufacturerID uint
}

// Connect to the database
func Connect() (*gorm.DB, error) {
	// Connect to database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	if err := db.AutoMigrate(&VehicleType{}, &Manufacturer{}); err != nil {
		log.Fatalln("Failed to migrate the schema!")
		return nil, err
	}

	return db, nil
}

// Close the database connection
func Close(db *gorm.DB) {
	DB, err := db.DB()
	if err != nil {
		log.Fatalln("Failed to get the database connection!")
	}
	// Close the connection
	DB.Close()
}

// Insert data into the database
func Insert(manufacturer Manufacturer, conn *gorm.DB) error {
	if err := conn.Create(&manufacturer).Error; err != nil {
		log.Fatalln("Failed to insert data!")
		return err
	}
	return nil
}

// Search for a manufacturer
func SearchManufacturer(search_criteria string, conn *gorm.DB) ([]Manufacturer, error) {
	var manufacturers []Manufacturer
	err := conn.Preload("VehicleTypes").Where("common_name = ?", search_criteria).Find(&manufacturers).Error
	return manufacturers, err
}
