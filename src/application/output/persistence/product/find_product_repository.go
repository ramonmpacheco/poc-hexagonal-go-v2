package productrepository

import (
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
)

type findProductRepository struct {
}

func NewFindProductRepository() productdomain.FindProductDataOutputPort {
	return &findProductRepository{}
}

func (fpr *findProductRepository) ById(product *productdomain.Product) error {
	product.Name = "teste"
	return nil
}
