package main

import (
	"fmt"
	"log"
	"net/http"
	"vanillaCRUDAPI/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", handlers.ProductsHandler)
	mux.HandleFunc("/product/{id}", handlers.ProductHandler)

	log.Printf("Starting server on port %d", 8080)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Sprintf(`There is a problem with the server: %s`, err)
		return
	}
}
