package productrepository

import (
	productdata "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/port/output/product"
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
)

type findProductRepository struct {
}

func NewFindProductRepository() productdata.FindProductData {
	return &findProductRepository{}
}

func (fpr *findProductRepository) ById(id string) (productdomain.Product, error) {
	product := productdomain.Product{
		ID:   id,
		Name: "teste",
	}

	return product, nil
}
