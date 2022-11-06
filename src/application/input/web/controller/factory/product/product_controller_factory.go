package productcontrollerfactory

import (
	"fmt"

	productcontroller "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/controller/product"
	productrepository "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/output/persistence/product"
	productservice "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/service/product"
)

func GetFindProductController() (productcontroller.IFindProductController, error) {
	repository := productrepository.NewFindProductRepository()
	service := productservice.NewFindProductService(repository)
	controller := productcontroller.NewFindProductController(service)

	if controller == nil {
		return nil, fmt.Errorf("error getting: find_product_controller")
	}

	return controller, nil
}
