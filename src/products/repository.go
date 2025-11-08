package products

import (
	"log"
)

type Repository interface {
	GetProductsDetail() []string
	GetProductsDetailById(id int) (string, error)
}

type repo struct {
	log *log.Logger
}

func NewRepository(log *log.Logger) Repository {
	return &repo{
		log: log,
	}
}

func (r *repo) GetProductsDetail() []string {

	r.log.Println("getProductsDetail Repository")
	return []string{"product1", "product2"}
}

func (r *repo) GetProductsDetailById(id int) (string, error) {

	r.log.Println("getProductsDetailById Repository")
	return "product2", nil
}
