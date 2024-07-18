// Package: utils
// File: utils.go
// Description : This file contains the utility functions for the data received from the api
package utils

import (
	"io"
	"log"
	"net/http"
)

// MakeRequest function to make a request to the data source
func MakeRequest(dataSource string, format string) ([]byte, error) {
	// Load data
	res, err := http.Get(dataSource + format)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if res.StatusCode != 200 {
		log.Fatal("Status code error")
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
