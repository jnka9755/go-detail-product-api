package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnka9755/go-detail-product-api/src/products"
)

func main() {

	router := mux.NewRouter()

	userController := products.MakeEndpoints()

	router.HandleFunc("/products", userController.GetProductsDetail).Methods("GET")

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
