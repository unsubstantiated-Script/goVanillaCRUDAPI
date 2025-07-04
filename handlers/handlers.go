package handlers

import (
	"encoding/json"
	"net/http"
	"vanillaCRUDAPI/storage"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//Handle Get All Products request
		json.NewEncoder(w).Encode(storage.Products)
	case http.MethodPost:
		var newProduct storage.Product

		//Read the value in from the response body to the newProduct
		json.NewDecoder(r.Body).Decode(&newProduct)

		//Loop through array to see if product is already there. If so, send error instead.
		for _, p := range storage.Products {
			if p.ID == newProduct.ID {
				http.Error(w, "Product already exists", http.StatusConflict)
				return
			}
		}

		storage.Products = append(storage.Products, newProduct)
		json.NewEncoder(w).Encode(newProduct)
		storage.WriteToFile()
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	switch r.Method {
	case http.MethodGet:
		//Handle get a single product by ID. Loop through items, find the item in the array (if it exits) return item as json
		for _, p := range storage.Products {
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
		var updatedProduct storage.Product
		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			return
		}
		for i, p := range storage.Products {
			if p.ID == id {
				storage.Products[i] = updatedProduct
				json.NewEncoder(w).Encode(storage.Products[i])
				storage.WriteToFile()
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	case http.MethodDelete:
		for i, p := range storage.Products {
			if p.ID == id {
				storage.Products = append(storage.Products[:i], storage.Products[i+1:]...)
				w.WriteHeader(http.StatusOK)
				storage.WriteToFile()
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
