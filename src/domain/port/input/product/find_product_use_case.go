package productusecase

import (
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
)

type FindProductUseCase interface {
	ById(id string) (productdomain.Product, error)
}
