package productcontrollerfactory

import (
	"fmt"

	"github.com/gin-gonic/gin"
	productcontroller "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/controller/product"
	productrepository "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/output/persistence/product"
	mongoclient "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/config/db"
	productusecaseimpl "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/port/input/product/impl"
)

func AddProductControllers(rest *gin.RouterGroup) error {
	repository := productrepository.NewFindProductRepository(mongoclient.GetCollection("products"))
	findProductUseCase := productusecaseimpl.NewFindProductUseCaseImpl(repository)
	findProductController := productcontroller.NewFindProductController(findProductUseCase)
	if findProductController == nil {
		return fmt.Errorf("error adding: find_product_controller")
	}
	rest.GET("/v1/products/:id", findProductController.ById)
	return nil
}
