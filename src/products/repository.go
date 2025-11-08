package products

import (
	"encoding/json"
	"log"
	"os"

	handleerrors "github.com/jnka9755/go-detail-product-api/src/products/handle-errors"
	"github.com/jnka9755/go-detail-product-api/src/products/models"
)

type Repository interface {
	GetProductsDetail() []models.Product
	GetProductsDetailById(id string) (models.Product, error)
}

type repo struct {
	log        *log.Logger
	categories map[string]models.Category
	sellers    map[string]models.Seller
	dataPath   string
}

func NewRepository(log *log.Logger, dataPath string) Repository {
	r := &repo{
		log:        log,
		categories: make(map[string]models.Category),
		sellers:    make(map[string]models.Seller),
		dataPath:   dataPath,
	}
	r.loadCategories()
	r.loadSellers()
	return r
}

func (r *repo) GetProductsDetail() []models.Product {
	r.log.Println("getProductsDetail Repository")

	file, err := os.Open(r.dataPath + "/products.json")
	if err != nil {
		r.log.Printf("Error opening products file: %v", err)
		return []models.Product{}
	}
	defer file.Close()

	var products []models.Product
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&products); err != nil {
		r.log.Printf("Error decoding products: %v", err)
		return []models.Product{}
	}

	// Enrich products with category and seller information
	for i := range products {
		if category, exists := r.categories[products[i].CategoryID]; exists {
			products[i].Category = category
		}
		if seller, exists := r.sellers[products[i].SellerID]; exists {
			products[i].Seller = seller
		}
	}

	return products
}

func (r *repo) GetProductsDetailById(id string) (models.Product, error) {
	r.log.Println("getProductsDetailById Repository")

	products := r.GetProductsDetail()

	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	return models.Product{}, handleerrors.ErrNotFound
}

func (r *repo) loadCategories() {
	file, err := os.Open(r.dataPath + "/categories.json")
	if err != nil {
		r.log.Printf("Error opening categories file: %v", err)
		return
	}
	defer file.Close()

	var categories []models.Category
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&categories); err != nil {
		r.log.Printf("Error decoding categories: %v", err)
		return
	}

	for _, category := range categories {
		r.categories[category.ID] = category
	}
}

func (r *repo) loadSellers() {
	file, err := os.Open(r.dataPath + "/sellers.json")
	if err != nil {
		r.log.Printf("Error opening sellers file: %v", err)
		return
	}
	defer file.Close()

	var sellers []models.Seller
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&sellers); err != nil {
		r.log.Printf("Error decoding sellers: %v", err)
		return
	}

	for _, seller := range sellers {
		r.sellers[seller.ID] = seller
	}
}
