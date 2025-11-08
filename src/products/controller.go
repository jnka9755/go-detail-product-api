package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		GetProductsDetail     Controller
		GetProductsDetailById Controller
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)

func MakeEndpoints(service Service) Endpoints {
	return Endpoints{
		GetProductsDetail:     makeGetProductsDetailEndpoint(service),
		GetProductsDetailById: makeGetProductsDetailByIdEndpoint(service),
	}
}

func makeGetProductsDetailEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getProductsDetail Controller")

		products := service.GetProductsDetail()
		json.NewEncoder(w).Encode(products)
	}
}

func makeGetProductsDetailByIdEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getProductsDetailById Controller")

		path := mux.Vars(r)
		id := path["id"]
		idInt, err := strconv.Atoi(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid id"})
			return
		}

		product, err := service.GetProductsDetailById(idInt)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Erorr getting product"})
			return
		}

		json.NewEncoder(w).Encode(product)
	}
}
