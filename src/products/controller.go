package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	handleerrors "github.com/jnka9755/go-detail-product-api/src/products/handle-errors"
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

		product, err := service.GetProductsDetailById(id)

		if err != nil {
			if errors.Is(err, handleerrors.ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(product)
	}
}
