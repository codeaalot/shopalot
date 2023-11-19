package main

import (
	"log"
	"net/http"

	"shopalot-backend/handlers"
)

func main() {
	http.HandleFunc("/products", handlers.ProductsHandler)
	http.HandleFunc("/products/", handlers.GetProductByID) // Notice the trailing slash
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
