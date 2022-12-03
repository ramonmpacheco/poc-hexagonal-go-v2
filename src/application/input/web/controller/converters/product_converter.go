package converters

import (
	"log"

	"github.com/jinzhu/copier"
	"github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/controller/responses"
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
)

func ProductDomainToResponse(product *productdomain.Product) responses.ProductResponse {
	var response responses.ProductResponse
	err := copier.Copy(&response, &product)
	if err != nil {
		log.Fatal("error converting product domain to response")
	}

	return response
}
