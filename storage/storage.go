package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var Products []Product //In-memory data store

var FilePath = "storage/products.json"

func WriteToFile() {
	jsonData, err := json.Marshal(Products)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile(FilePath, jsonData, 0644) //Permission code for read/write

	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	fmt.Println("Successfully wrote to file", FilePath)
}
