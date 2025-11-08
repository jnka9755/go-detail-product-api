package products

import (
	"log"
)

type Service interface {
	GetProductsDetail() []string
	GetProductsDetailById(id int) (string, error)
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

func (s *serv) GetProductsDetail() []string {

	s.log.Println("getProductsDetail Service")

	products := s.repo.GetProductsDetail()

	return products
}

func (s *serv) GetProductsDetailById(id int) (string, error) {

	s.log.Println("getProductsDetailById Service")

	product, err := s.repo.GetProductsDetailById(id)
	return product, err
}
