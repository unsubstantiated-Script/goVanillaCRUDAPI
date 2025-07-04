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

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	switch r.Method {
	case http.MethodGet:
		//Handle get a single product by ID. Loop through items, find the item in the array (if it exits) return item as json
		for _, p := range Products {
			if p.ID == id {
				err := json.NewEncoder(w).Encode(p)
				if err != nil {
					return
				}
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	case http.MethodPut:
		var updatedProduct Product
		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			return
		}
		for i, p := range Products {
			if p.ID == id {
				Products[i] = updatedProduct
				err := json.NewEncoder(w).Encode(Products[i])
				if err != nil {
					return
				}
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	case http.MethodDelete:
		for i, p := range Products {
			if p.ID == id {
				Products = append(Products[:i], Products[i+1:]...)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
