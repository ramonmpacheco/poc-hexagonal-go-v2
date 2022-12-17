package producttemplates

import (
	"github.com/google/uuid"
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
)

func GetProductDomain(name string) *productdomain.Product {
	return &productdomain.Product{
		ID:   uuid.NewString(),
		Name: getString(name, "TV"),
	}
}

func getString(v string, fallback string) string {
	if v == "" {
		return fallback
	}
	return v
}
