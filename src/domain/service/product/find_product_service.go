package productservice

import (
	productusecase "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/port/input/product"
	productdata "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/port/output/product"
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
)

type findProductService struct {
	findProductData productdata.FindProductData
}

func NewFindProductService(findProductData productdata.FindProductData) productusecase.FindProductUseCase {
	return &findProductService{
		findProductData,
	}
}

func (fps *findProductService) ById(id string) (productdomain.Product, error) {
	return fps.findProductData.ById(id)
}
