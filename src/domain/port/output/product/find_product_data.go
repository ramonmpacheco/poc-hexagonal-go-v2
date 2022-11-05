package productdata

import productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"

type FindProductData interface {
	ById(id string) (productdomain.Product, error)
}
