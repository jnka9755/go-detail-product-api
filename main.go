package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnka9755/go-detail-product-api/src/products"
)

func main() {

	router := mux.NewRouter()
	logger := log.New(os.Stdout, "[server] ", log.LstdFlags|log.Lshortfile)

	productsRepo := products.NewRepository(logger)
	productsServ := products.NewService(logger, productsRepo)
	productController := products.MakeEndpoints(productsServ)

	router.HandleFunc("/products", productController.GetProductsDetail).Methods("GET")
	router.HandleFunc("/product/{id}", productController.GetProductsDetailById).Methods("GET")

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
