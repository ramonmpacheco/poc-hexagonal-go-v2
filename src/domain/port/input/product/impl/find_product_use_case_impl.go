package productusecaseimpl

import (
	"errors"
	"log"

	productusecase "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/port/input/product"
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type findProductUseCaseImpl struct {
	findProductData productdomain.FindProductDataOutputPort
}

func NewFindProductUseCaseImpl(findProductData productdomain.FindProductDataOutputPort) productusecase.FindProductUseCase {
	return &findProductUseCaseImpl{
		findProductData,
	}
}

func (fps *findProductUseCaseImpl) ById(id string) (*productdomain.Product, error) {
	log.Default().Println("findProductUseCaseImpl, init")

	if !primitive.IsValidObjectID(id) {
		log.Default().Println("findProductUseCaseImpl, error, invalid id")
		return nil, errors.New("invalid id")
	}

	product := productdomain.Product{ID: id}
	if err := product.FindById(fps.findProductData); err != nil {
		log.Default().Printf("findProductUseCaseImpl, error, %s\n", err.Error())
		return nil, err
	}

	log.Default().Println("findProductUseCaseImpl, success")
	return &product, nil
}
