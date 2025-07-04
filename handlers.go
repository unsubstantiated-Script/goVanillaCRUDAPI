package main

import (
	"encoding/json"
	"net/http"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//Handle Get All Products request
		json.NewEncoder(w).Encode(Products)
	case http.MethodPost:
		var newProduct Product

		//Read the value in from the response body to the newProduct
		json.NewDecoder(r.Body).Decode(&newProduct)

		//Loop through array to see if product is already there. If so, send error instead.
		for _, p := range Products {
			if p.ID == newProduct.ID {
				http.Error(w, "Product already exists", http.StatusConflict)
				return
			}
		}

		Products = append(Products, newProduct)
		json.NewEncoder(w).Encode(newProduct)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
