package productusecase

import (
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
)

type findProductUseCaseImpl struct {
	findProductData productdomain.FindProductDataOutputPort
}

func NewFindProductUseCaseImpl(findProductData productdomain.FindProductDataOutputPort) FindProductUseCase {
	return &findProductUseCaseImpl{
		findProductData,
	}
}

func (fps *findProductUseCaseImpl) ById(id string) (*productdomain.Product, error) {
	product := productdomain.Product{ID: id}
	if err := product.FindById(fps.findProductData); err != nil {
		return nil, err
	}

	return &product, nil
}
