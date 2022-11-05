package routes

import (
	"github.com/gin-gonic/gin"
	productcontroller "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/controllers/product"
	productrepository "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/output/persistence/product"
	productservice "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/service/product"
)

type Routes struct {
	Router *gin.Engine
}

func (r *Routes) LoadRoutes() (err error) {
	findProductRepository := productrepository.NewFindProductRepository()
	findProductService := productservice.NewFindProductService(findProductRepository)
	findProductController := productcontroller.NewFindProductController(findProductService)

	rest := r.Router.Group("/hex-go")
	rest.GET("/v1/products/:id", findProductController.ById)
	return
}
