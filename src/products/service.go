package products

import (
	"log"

	"github.com/jnka9755/go-detail-product-api/src/products/models"
)

type Service interface {
	GetProductsDetail() []models.Product
	GetProductsDetailById(id string) (models.Product, error)
}

type serv struct {
	log  *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) Service {
	return &serv{
		log:  log,
		repo: repo,
	}
}

func (s *serv) GetProductsDetail() []models.Product {

	s.log.Println("getProductsDetail Service")

	products := s.repo.GetProductsDetail()

	return products
}

func (s *serv) GetProductsDetailById(id string) (models.Product, error) {

	s.log.Println("getProductsDetailById Service")

	product, err := s.repo.GetProductsDetailById(id)
	return product, err
}
